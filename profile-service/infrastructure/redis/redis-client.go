package redis

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	sessionTime = 1 * time.Hour

	PrefixKeyProfiles        = "profiles"
	PrefixKeyEmployments     = "employments"
	PrefixKeyEducations      = "educations"
	PrefixKeySkills          = "skills"
	PrefixKeyWorkExperiences = "work_experiences"
)

type Redis struct {
	Redis  *redis.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

func NewReddisClient(ctx context.Context) (*Redis, error) {
	// Connect to Redis
	ctxRedis, cancelRedis := context.WithCancel(ctx)
	host := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	var opts *redis.Options = &redis.Options{
		Addr:        host,
		Password:    "", // No password for local development
		DB:          0,
		DialTimeout: 10 * time.Second,
	}
	client := redis.NewClient(opts)

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed ping server redis. err = %v", err)
	}

	var redis *Redis = &Redis{
		Redis:  client,
		Ctx:    ctxRedis,
		Cancel: cancelRedis,
	}
	return redis, nil
}

func (r *Redis) Close() {
	r.Redis.Close()
	r.Cancel()
}
