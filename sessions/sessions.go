package sessions

import (
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var REDIS_OPTS = &redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       1,
}

type SessionManager struct {
	client *redis.Client
}

func NewSessionManager() *SessionManager {
	rdb := redis.NewClient(REDIS_OPTS)

	log.Info().Msg("Connected To Redis")
	return &SessionManager{
		client: rdb,
	}
}
