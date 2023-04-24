package redisclients

import redigo "github.com/gomodule/redigo/redis"

type RedigoClient struct {
	conn redigo.Conn
}

func (c *RedigoClient) Name() string {
	return "https://github.com/gomodule/redigo/redis"
}

func (c *RedigoClient) Get(key string) (string, error) {
	return redigo.String(c.conn.Do("GET", key))
}

func (c *RedigoClient) Set(key, value string) error {
	_, err := c.conn.Do("SET", key, value)
	return err
}

func (c *RedigoClient) Teardown() {
	c.conn.Close()
}

var _ GenericRedisClient = &RedigoClient{}

// NewRedigoClient creates RedigoClient
func NewRedigoClient() GenericRedisClient {
	conn, err := redigo.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	return &RedigoClient{
		conn: conn,
	}
}
