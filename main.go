package main

//go:generate go run github.com/a-h/templ/cmd/templ@latest generate

import (
	"net/http"
	"os"

	"github.com/ShardulNalegave/PICT-EMS/database"
	"github.com/ShardulNalegave/PICT-EMS/routes"
	"github.com/ShardulNalegave/PICT-EMS/sessions"
	"github.com/ShardulNalegave/PICT-EMS/tsdb"
	"github.com/ShardulNalegave/PICT-EMS/utils"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	ADDR = ":8080"
)

func main() {
	if godotenv.Load() != nil {
		panic("Couldn't load .env")
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	if len(os.Args) < 2 {
		panic("Not enough arguments provided")
	}

	loc := os.Args[1]
	t := tsdb.ConnectToTSDB()
	db := database.ConnectToDatabase()
	sm := sessions.NewSessionManager()

	r := chi.NewRouter()
	r.Use(tsdb.TSDBMiddleware(t))
	r.Use(database.DatabaseMiddleware(db))
	r.Use(sessions.SessionManagerMiddleware(sm))
	r.Use(utils.LocationMiddleware(loc))
	r.Mount("/", routes.GetRouter())

	log.Info().
		Str("Addr", ADDR).
		Msg("Listening...")
	err := http.ListenAndServe(ADDR, r)
	log.Fatal().Err(err).Msg("Server shut down")
}
