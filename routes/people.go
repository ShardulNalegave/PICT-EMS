package routes

import (
	"net/http"

	"github.com/ShardulNalegave/PICT-EMS/utils"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func initPeopleRoutes(r *chi.Mux) {
	r.Get("/people/students", getStudents)
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	_ = r.Context().Value(utils.DatabaseKey).(*gorm.DB)
}
