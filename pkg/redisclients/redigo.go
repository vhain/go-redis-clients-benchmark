package redisclients

import (
	"reflect"

	redigo "github.com/gomodule/redigo/redis"
)

type RedigoClient struct {
	conn redigo.Conn
}

func (c *RedigoClient) Name() string {
	return reflect.TypeOf(c.conn).Elem().PkgPath()
}

func (c *RedigoClient) Get(key string) (string, error) {
	val, err := c.conn.Do("GET", key)
	return string([]byte(val.([]uint8))), err
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
