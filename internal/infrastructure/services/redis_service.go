package services

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type TokenBlacklist struct {
	rdb *redis.Client
	ctx context.Context
}

func NewTokenBlacklistRepository(rdb *redis.Client) *TokenBlacklist {
	return &TokenBlacklist{
		rdb: rdb,
		ctx: context.Background(),
	}
}

func (r *TokenBlacklist) Add(token string, exp time.Time) error {
	ttl := time.Until(exp)
	if ttl <= 0 {
		ttl = time.Hour
	}
	return r.rdb.Set(r.ctx, "blacklist:"+token, true, ttl).Err()
}

func (r *TokenBlacklist) IsBlacklisted(token string) (bool, error) {
	val, err := r.rdb.Exists(r.ctx, "blacklist:"+token).Result()
	if err != nil {
		return false, err
	}
	return val == 1, nil
}

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return client
}
