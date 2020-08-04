package store

import "github.com/go-redis/redis"

type Store struct {
	User, Pass, Url string
	Client          *redis.Client
}

func NewStore(url, pass string) *Store {
	return &Store{
		Client: redis.NewClient(&redis.Options{
			Addr:     url,
			Password: pass,
			DB:       0,
		}),
	}
}
