package redisclients

import (
	"context"
	"reflect"

	goredis "github.com/redis/go-redis/v9"
)

type GoRedisClient struct {
	client *goredis.Client
}

func (c *GoRedisClient) Name() string {
	return reflect.TypeOf(c.client).Elem().PkgPath()
}

func (c *GoRedisClient) Get(key string) (string, error) {
	return c.client.Get(context.Background(), key).Result()
}

func (c *GoRedisClient) Set(key, value string) error {
	return c.client.Set(context.Background(), key, value, 0).Err()
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
