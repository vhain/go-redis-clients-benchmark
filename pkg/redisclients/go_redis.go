package redisclients

import (
	goredis "github.com/go-redis/redis/v7"
)

type GoRedisClient struct {
	client *goredis.Client
}

func (c *GoRedisClient) Name() string {
	return "https://github.com/go-redis/redis/v7"
}

func (c *GoRedisClient) Get(key string) (string, error) {
	return c.client.Get(key).Result()
}

func (c *GoRedisClient) Set(key, value string) error {
	return c.client.Set(key, value, 0).Err()
}

func (c *GoRedisClient) Teardown() {
	c.client.Close()
}

var _ GenericRedisClient = &GoRedisClient{}

// NewGoRedisClient creates GoRedisClient
func NewGoRedisClient() GenericRedisClient {
	return &GoRedisClient{
		client: goredis.NewClient(&goredis.Options{}),
	}
}
