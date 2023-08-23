package redis

/*
	used for state management between shard and login servers
*/

import (
	"context"

	"github.com/CPunch/gopenfusion/config"

	"github.com/redis/go-redis/v9"
)

type RedisHandler struct {
	client *redis.Client
	ctx    context.Context
}

const (
	SHARD_SET = "shards"
)

func OpenRedis(addr string) (*RedisHandler, error) {
	client := redis.NewClient(&redis.Options{
		Addr:                addr,
		CredentialsProvider: config.GetRedisCredentials(),
	})

	_, err := client.Ping(context.Background()).Result()
	return &RedisHandler{client: client, ctx: context.Background()}, err
}

func (r *RedisHandler) Close() error {
	return r.client.Close()
}
