package redis

import "github.com/go-redis/redis/v8"

type Store struct {
	redis *redis.Client
}

func NewStore(redis *redis.Client) *Store {
	return &Store{
		redis: redis,
	}
}
