package redis

/*
	used for state management between shard and login servers
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/CPunch/gopenfusion/config"

	"github.com/redis/go-redis/v9"
)

type RedisHandler struct {
	client *redis.Client
	ctx    context.Context
}

type LoginMetadata struct {
	FEKey    []byte `json:",omitempty"`
	PlayerID int32  `json:",omitempty"`
}

type ShardMetadata struct {
	IP   string
	Port int
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

// we store login queues into redis with the name "loginMetadata_<serialKey>"
// set to expire after config.LOGIN_TIMEOUT duration. this way we can easily
// have a shared pool of active serial keys & player login data which any
// shard can pull from

func makeLoginMetadataKey(serialKey int64) string {
	return fmt.Sprintf("loginMetadata_%s", strconv.Itoa(int(serialKey)))
}

func (r *RedisHandler) QueueLogin(serialKey int64, data LoginMetadata) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// add to table
	return r.client.Set(r.ctx, makeLoginMetadataKey(serialKey), value, config.LOGIN_TIMEOUT).Err()
}

func (r *RedisHandler) GetLogin(serialKey int64) (LoginMetadata, error) {
	value, err := r.client.Get(r.ctx, makeLoginMetadataKey(serialKey)).Result()
	if err != nil {
		return LoginMetadata{}, err
	}

	var data LoginMetadata
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return LoginMetadata{}, err
	}

	return data, nil
}

func (r *RedisHandler) RemoveLogin(serialKey int64) error {
	return r.client.Del(r.ctx, makeLoginMetadataKey(serialKey)).Err()
}

func (r *RedisHandler) RegisterShard(shard ShardMetadata) error {
	value, err := json.Marshal(shard)
	if err != nil {
		return err
	}

	return r.client.SAdd(r.ctx, SHARD_SET, value).Err()
}

func (r *RedisHandler) GetShards() []ShardMetadata {
	shardData := r.client.SMembers(r.ctx, SHARD_SET).Val()

	// unmarshal all shards
	shards := make([]ShardMetadata, 0, len(shardData))
	for _, data := range shardData {
		var shard ShardMetadata
		if err := json.Unmarshal([]byte(data), &shard); err != nil {
			continue
		}

		shards = append(shards, shard)
	}

	return shards
}
