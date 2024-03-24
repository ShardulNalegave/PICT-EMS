package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ShardulNalegave/PICT-EMS/database/models"
	"github.com/ShardulNalegave/PICT-EMS/sessions"
	"github.com/ShardulNalegave/PICT-EMS/tsdb"
	"github.com/ShardulNalegave/PICT-EMS/utils"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func initSessionRoutes(r *chi.Mux) {
	r.Post("/sessions", createOrEndSessionHandler)
	r.Post("/sessions/end-day", endDay)
}

func endDay(w http.ResponseWriter, r *http.Request) {
	t := r.Context().Value(utils.TSDBKey).(*tsdb.TSDB)
	sm := r.Context().Value(utils.SessionManagerKey).(*sessions.SessionManager)
	loc := r.Context().Value(utils.LocationKey).(string)

	keys, err := sm.GetSessionKeys(loc)
	if err != nil {
		http.Error(w, "Couldn't get session keys", http.StatusInternalServerError)
		return
	}

	for _, key := range keys {
		_ = endSession(t, sm, key)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Done"))
}

func createOrEndSessionHandler(w http.ResponseWriter, r *http.Request) {
	t := r.Context().Value(utils.TSDBKey).(*tsdb.TSDB)
	db := r.Context().Value(utils.DatabaseKey).(*gorm.DB)
	sm := r.Context().Value(utils.SessionManagerKey).(*sessions.SessionManager)
	loc := r.Context().Value(utils.LocationKey).(string)

	var body struct {
		RegistrationID string `json:"registration_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var person models.Person
	if err := db.First(&person, "registration_id = ?", body.RegistrationID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "No one with given RegistrationID exists", http.StatusNotFound)
			return
		} else {
			http.Error(w, "Couldn't create or end session", http.StatusInternalServerError)
			return
		}
	}

	session_id := fmt.Sprintf("%s-%s", body.RegistrationID, loc)

	s, err := sm.GetSession(session_id)
	if err == nil && s != nil {
		// Session already exists!
		err := endSession(t, sm, session_id)
		if err != nil {
			http.Error(w, "Couldn't end session", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Done"))
		return
	}

	// Session doesn't already exist
	s = &sessions.Session{
		SessionID:      session_id,
		RegistrationID: body.RegistrationID,
		EntryTime:      time.Now(),
		Location:       loc,
	}

	err = sm.CreateSession(*s)
	if err != nil {
		http.Error(w, "Couldn't create or end session", http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(s)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func endSession(
	t *tsdb.TSDB,
	sm *sessions.SessionManager,
	session_id string,
) error {
	s, err := sm.GetSession(session_id)
	if err != nil {
		// Session doesn't already exist
		return err
	}

	// Session already exists!
	err = t.WriteSessionData(s)
	if err != nil {
		return err
	}

	err = sm.DeleteSession(session_id)
	if err != nil {
		return err
	}

	return nil
}
