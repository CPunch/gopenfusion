package redis

import "encoding/json"

type ShardMetadata struct {
	IP   string
	Port int
}

const (
	SHARD_SET = "shards"
)

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
