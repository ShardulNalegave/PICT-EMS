package routes

import (
	"github.com/ShardulNalegave/PICT-EMS/routes/templates"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func initTemplatesRouter(r *chi.Mux) {
	r.Mount("/", templ.Handler(templates.HomeView()))
	r.Mount("/generate-report", templ.Handler(templates.ReportsView()))
}
