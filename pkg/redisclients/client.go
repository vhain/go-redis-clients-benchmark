package redisclients

type GenericRedisClient interface {
	Get(string) (string, error)
	Set(string, string) error
	Name() string
	Teardown()
}
