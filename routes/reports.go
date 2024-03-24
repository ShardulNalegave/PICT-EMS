package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ShardulNalegave/PICT-EMS/excel"
	"github.com/ShardulNalegave/PICT-EMS/tsdb"
	"github.com/ShardulNalegave/PICT-EMS/utils"
	"github.com/go-chi/chi/v5"
)

func initReportRoutes(r *chi.Mux) {
	r.Post("/report", getReport)
}

func getReport(w http.ResponseWriter, r *http.Request) {
	t := r.Context().Value(utils.TSDBKey).(*tsdb.TSDB)

	var body struct {
		StartTime string `json:"start_time"`
		StopTime  string `json:"stop_time"`
		Location  string `json:"location"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	start_time, err := time.Parse(time.RFC3339, body.StartTime)
	if err != nil {
		http.Error(w, "Couldn't parse start time", http.StatusBadRequest)
		return
	}

	stop_time, err := time.Parse(time.RFC3339, body.StopTime)
	if err != nil {
		http.Error(w, "Couldn't parse stop time", http.StatusBadRequest)
		return
	}

	s := t.GetSessions(start_time, stop_time, body.Location)
	fname := excel.CreateReportFile(s)
	http.ServeFile(w, r, fname)
}
