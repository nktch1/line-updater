package store

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/nikitych1w/softpro-task/internal/config"
	"github.com/sirupsen/logrus"
	"strconv"
)

// Store ...
type Store struct {
	Client *redis.Client
}

// New создает подключение к redis
func New(cfg *config.Config) *Store {
	return &Store{
		Client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.Database.Host, cfg.Database.Port),
			Password: cfg.Database.Password,
			DB:       0,
		}),
	}
}

// GetLastValueByKey получает последнее значение по заданному ключу из базы данных
func (s *Store) GetLastValueByKey(key string) (float32, error) {
	var val float64
	var err error

	out := s.Client.LRange(key, -1, -1)
	if out.Err() != nil {
		return 0, out.Err()
	}

	if len(out.Val()) > 0 {
		if val, err = strconv.ParseFloat(out.Val()[0], 64); err != nil {
			return 0, err
		}
	} else {
		logrus.Errorf("there is no key like this '%s'! GRPC Request is wrong!", key)
		return 0, err
	}

	return float32(val), out.Err()
}

// Set производит запись в базу данных
func (s *Store) Set(key string, val interface{}) error {
	err := s.Client.RPush(key, val)
	if err.Err() != nil {
		return err.Err()
	}

	return nil
}

// Ping проверяет готовности базы данных
func (s *Store) Ping() error {
	_, err := s.Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
