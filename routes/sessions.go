package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ShardulNalegave/PICT-EMS/database/models"
	"github.com/ShardulNalegave/PICT-EMS/sessions"
	"github.com/ShardulNalegave/PICT-EMS/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func initSessionRoutes(r *chi.Mux) {
	r.Get("/sessions/{reg_id}", getSession)
	r.Delete("/sessions/{reg_id}", deleteSession)
	r.Post("/sessions", createSession)
}

func getSession(w http.ResponseWriter, r *http.Request) {
	sm := r.Context().Value(utils.SessionManagerKey).(*sessions.SessionManager)

	s, err := sm.GetSession(chi.URLParam(r, "reg_id"))
	if err != nil {
		if err == redis.Nil {
			http.Error(w, "Session doesn't exist", http.StatusNotFound)
			return
		} else {
			http.Error(w, "Couldn't get session record", http.StatusInternalServerError)
			return
		}
	}

	data, _ := json.Marshal(s)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func createSession(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(utils.DatabaseKey).(*gorm.DB)
	sm := r.Context().Value(utils.SessionManagerKey).(*sessions.SessionManager)

	var body struct {
		RegistrationID string `json:"registration_id"`
		Kind           string `json:"kind"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if body.Kind == "student" {
		var student models.Student
		if err := db.First(&student, "registration_id = ?", body.RegistrationID).Error; err != nil {
			http.Error(w, "Student with given RegistrationID doesn't exist", http.StatusNotFound)
			return
		}
	} else if body.Kind == "staff" {
		var staffMember models.StaffMember
		if err := db.First(&staffMember, "registration_id = ?", body.RegistrationID).Error; err != nil {
			http.Error(w, "Staff-Member with given RegistrationID doesn't exist", http.StatusNotFound)
			return
		}
	} else {
		http.Error(w, "'kind' not provided in request body", http.StatusBadRequest)
		return
	}

	s := sessions.Session{
		RegistrationID: body.RegistrationID,
		EntryTime:      time.Now(),
	}
	err := sm.CreateSession(s)
	if err != nil {
		http.Error(w, "Couldn't create session record", http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(s)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func deleteSession(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(utils.DatabaseKey).(*gorm.DB)
	sm := r.Context().Value(utils.SessionManagerKey).(*sessions.SessionManager)

	s, err := sm.GetSession(chi.URLParam(r, "reg_id"))
	if err != nil {
		if err == redis.Nil {
			http.Error(w, "Session doesn't exist", http.StatusNotFound)
			return
		} else {
			http.Error(w, "Couldn't get session record", http.StatusInternalServerError)
			return
		}
	}

	err = db.Save(&models.Session{
		SessionID:      uuid.NewString(),
		RegistrationID: s.RegistrationID,
		EntryTime:      s.EntryTime,
		ExitTime:       time.Now(),
	}).Error
	if err != nil {
		http.Error(w, "Couldn't delete session record", http.StatusInternalServerError)
		return
	}

	err = sm.DeleteSession(chi.URLParam(r, "reg_id"))
	if err != nil {
		http.Error(w, "Couldn't delete session record", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Done"))
}
