package redisclients

import (
	goresp3 "github.com/stfnmllr/go-resp3/client"
)

type GoResp3Client struct {
	conn goresp3.Conn
}

func (c *GoResp3Client) Name() string {
	return "https://github.com/stfnmllr/go-resp3/client"
}

func (c *GoResp3Client) Get(key string) (string, error) {
	return c.conn.Get(key).ToString()
}

func (c *GoResp3Client) Set(key, value string) error {
	return c.conn.Set(key, value).Err()
}

func (c *GoResp3Client) Teardown() {
	c.conn.Close()
}

var _ GenericRedisClient = &GoResp3Client{}

// NewGoResp3Client creates GoResp3Client
func NewGoResp3Client() GenericRedisClient {
	conn, err := goresp3.Dial("")
	if err != nil {
		panic(err)
	}

	return &GoResp3Client{
		conn: conn,
	}
}
