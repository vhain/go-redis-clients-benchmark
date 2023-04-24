package redisclients

import (
	"reflect"

	pascaldekloe_redis "github.com/pascaldekloe/redis"
)

type PascaldekloeRedisClient struct {
	client *pascaldekloe_redis.Client
}

func (c *PascaldekloeRedisClient) Name() string {
	return reflect.TypeOf(c.client).Elem().PkgPath()
}

func (c *PascaldekloeRedisClient) Get(key string) (string, error) {
	val, err := c.client.GET(key)
	return string(val), err
}

func (c *PascaldekloeRedisClient) Set(key, value string) error {
	return c.client.SET(key, []byte(value))
}

func (c *PascaldekloeRedisClient) Teardown() {
	c.client.Close()
}

var _ GenericRedisClient = &RadixClient{}

// NewPascaldekloeRedisClient creates PascaldekloeRedisClient
func NewPascaldekloeRedisClient() GenericRedisClient {
	client := pascaldekloe_redis.NewClient(":6379", 0, 0)

	return &PascaldekloeRedisClient{
		client: client,
	}
}
