package main

import (
	"testing"

	"github.com/vhain/go-redis-clients-benchmark/pkg/redisclients"
)

func BenchmarkRedisClients(b *testing.B) {
	clients := []redisclients.GenericRedisClient{
		redisclients.NewGoRedisClient(),
		redisclients.NewGoResp3Client(),
		redisclients.NewPascaldekloeRedisClient(),
		redisclients.NewRedigoClient(),
		redisclients.NewRadixClient(),
	}

	for _, client := range clients {
		client := client
		b.Run(client.Name(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := client.Set("test", "test")
				if err != nil {
					b.Error(err)
				}

				_, err = client.Get("test")
				if err != nil {
					b.Error(err)
				}
			}
		})
	}

	for _, client := range clients {
		client.Teardown()
	}
}
