package redisclients

import (
	radix "github.com/mediocregopher/radix/v3"
)

type RadixClient struct {
	conn radix.Conn
}

func (c *RadixClient) Name() string {
	return "https://github.com/mediocregopher/radix/v3"
}

func (c *RadixClient) Get(key string) (string, error) {
	var str string
	return str, c.conn.Do(radix.Cmd(&str, "GET", key))
}

func (c *RadixClient) Set(key, value string) error {
	return c.conn.Do(radix.Cmd(nil, "SET", key, value))
}

func (c *RadixClient) Teardown() {
	c.conn.Close()
}

var _ GenericRedisClient = &RadixClient{}

// NewRadixClient creates RadixClient
func NewRadixClient() GenericRedisClient {
	conn, err := radix.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	return &RadixClient{
		conn: conn,
	}
}
