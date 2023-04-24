package redisclients

import (
	"context"
	"reflect"

	radix "github.com/mediocregopher/radix/v4"
)

type RadixClient struct {
	conn radix.Conn
}

func (c *RadixClient) Name() string {
	return reflect.TypeOf(c.conn).Elem().PkgPath()
}

func (c *RadixClient) Get(key string) (string, error) {
	var str string
	return str, c.conn.Do(context.Background(), radix.Cmd(&str, "GET", key))
}

func (c *RadixClient) Set(key, value string) error {
	return c.conn.Do(context.Background(), radix.Cmd(nil, "SET", key, value))
}

func (c *RadixClient) Teardown() {
	c.conn.Close()
}

var _ GenericRedisClient = &RadixClient{}

// NewRadixClient creates RadixClient
func NewRadixClient() GenericRedisClient {
	conn, err := radix.Dial(context.Background(), "tcp", ":6379")
	if err != nil {
		panic(err)
	}

	return &RadixClient{
		conn: conn,
	}
}
