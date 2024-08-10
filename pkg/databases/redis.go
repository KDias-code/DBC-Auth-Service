package databases

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func RedisConnection(redisAddress, redisPassword string, redisDb int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       redisDb,
	})

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
