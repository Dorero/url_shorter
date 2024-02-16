package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type IUrlRepository interface {
	Find(id string) (string, error)
	Create(path string) (string, error)
}

type UrlRepository struct {
	Cache *redis.Client
}

func (u *UrlRepository) Create(path string) (string, error) {
	val, err := u.Cache.Incr(context.Background(), "counter").Result()
	url := fmt.Sprintf("/url/%s", strconv.FormatInt(val, 10))

	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	err = u.Cache.Set(context.Background(), url, path, time.Minute*7*24).Err()
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return url, nil
}

func (u *UrlRepository) Find(id string) (string, error) {
	url, err := u.Cache.Get(context.Background(), id).Result()
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return url, nil
}
