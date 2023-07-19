package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	Repository
	client *redis.Client
}

func NewRedisRepository(addr, password string) *RedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return &RedisRepository{
		client: client,
	}
}

func (r *RedisRepository) CreateLink(shortLink, url string) error {
	ctx := context.Background()
	err := r.client.Set(ctx, shortLink, url, 0).Err()
	return err
}

func (r *RedisRepository) GetLink(shortLink string) (string, bool, error) {
	ctx := context.Background()
	exists, err := r.LinkExists(shortLink)
	if err != nil {
		return "", false, err
	}

	if !exists {
		return "", false, nil
	}

	link, err := r.client.Get(ctx, shortLink).Result()
	if err != nil {
		return "", false, err
	}

	return link, true, nil
}

func (r *RedisRepository) LinkExists(shortLink string) (bool, error) {
	ctx := context.Background()
	exists, err := r.client.Exists(ctx, shortLink).Result()
	if err != nil {
		return false, err
	}

	return exists != 0, nil
}
