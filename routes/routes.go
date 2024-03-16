package routes

import "github.com/go-chi/chi/v5"

func GetRouter() *chi.Mux {
	r := chi.NewRouter()
	initPeopleRoutes(r)
	return r
}
