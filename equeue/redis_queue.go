package equeue

import (
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
	sync.Mutex
}

func NewRedisClient() (*RedisClient, error) {

	config := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	client := redis.NewClient(config)

	if _, err := client.Ping(client.Context()).Result(); err != nil {
		return nil, err
	}
	return &RedisClient{client: client}, nil
}

func (r *RedisClient) Queue(queueName string, value string) error {
	r.Lock()
	defer r.Unlock()
	if _, err := r.client.Append(r.client.Context(), queueName, value).Result(); err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) Dequeue(queueName string) (string, error) {
	r.Lock()
	defer r.Unlock()
	item, err := r.client.BLPop(r.client.Context(), time.Duration(60), queueName).Result()
	if err != nil {
		return "", err
	}
	return item[0], nil
}

func (r *RedisClient) Front(queueName string) (string, error) {
	r.Lock()
	defer r.Unlock()
	item, err := r.client.LIndex(r.client.Context(), queueName, 0).Result()
	if err != nil {
		return "", err
	}
	return item, nil
}

func (r *RedisClient) IsFront(queueName string, value string) (bool, error) {
	frontValue, err := r.Front(queueName)
	if err != nil {
		return false, err
	}
	if frontValue == value {
		return true, nil
	} else {
		return false, nil
	}
}
