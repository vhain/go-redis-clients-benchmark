package redisclients

import (
	pascaldekloe_redis "github.com/pascaldekloe/redis"
)

type PascaldekloeRedisClient struct {
	client *pascaldekloe_redis.Client
}

func (c *PascaldekloeRedisClient) Name() string {
	return "https://github.com/pascaldekloe/redis"
}

func (c *PascaldekloeRedisClient) Get(key string) (string, error) {
	value, _, err := c.client.GETString(key)
	return value, err
}

func (c *PascaldekloeRedisClient) Set(key, value string) error {
	return c.client.SETString(key, value)
}

func (c *PascaldekloeRedisClient) Teardown() {
	c.client.Close()
}

var _ GenericRedisClient = &PascaldekloeRedisClient{}

// NewPascaldekloeRedisClient creates PascaldekloeRedisClient
func NewPascaldekloeRedisClient() GenericRedisClient {
	return &PascaldekloeRedisClient{
		client: pascaldekloe_redis.NewClient(":6379", 0, 0),
	}
}
