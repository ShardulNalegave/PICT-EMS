package sessions

import (
	"context"
	"encoding/json"
	"fmt"

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

func (sm *SessionManager) GetSession(id string) (*Session, error) {
	val, err := sm.client.Get(context.Background(), id).Bytes()
	if err != nil {
		return nil, err
	}

	var s Session
	if err := json.Unmarshal(val, &s); err != nil {
		return nil, err
	}

	return &s, nil
}

func (sm *SessionManager) DeleteSession(id string) error {
	return sm.client.Del(context.Background(), id).Err()
}

func (sm *SessionManager) CreateSession(s Session) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return sm.client.Set(context.Background(), s.SessionID, data, 0).Err()
}

func (sm *SessionManager) GetSessionKeys(loc string) ([]string, error) {
	keys, err := sm.client.Keys(context.Background(), fmt.Sprintf("*-%s", loc)).Result()
	if err != nil {
		return nil, err
	}
	return keys, nil
}
