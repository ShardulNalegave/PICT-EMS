package main

import (
	"net/http"
	"os"

	"github.com/ShardulNalegave/PICT-EMS/database"
	"github.com/ShardulNalegave/PICT-EMS/sessions"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	ADDR = ":8080"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	db := database.ConnectToDatabase()
	sm := sessions.NewSessionManager()

	r := chi.NewRouter()
	r.Use(database.DatabaseMiddleware(db))
	r.Use(sessions.SessionManagerMiddleware(sm))

	log.Info().
		Str("Addr", ADDR).
		Msg("Listening...")
	err := http.ListenAndServe(ADDR, r)
	log.Fatal().Err(err).Msg("Server shut down")
}
