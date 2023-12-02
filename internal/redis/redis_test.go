package redis_test

import (
	"os"
	"testing"

	"github.com/CPunch/gopenfusion/internal/redis"
	"github.com/alicebob/miniredis/v2"
	"github.com/matryer/is"
)

var (
	rh *redis.RedisHandler
)

func TestMain(m *testing.M) {
	r, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer r.Close()

	rh, err = redis.OpenRedis(r.Addr())
	if err != nil {
		panic(err)
	}
	defer rh.Close()

	os.Exit(m.Run())
}

func TestRedisLogin(t *testing.T) {
	is := is.New(t)

	// test data
	serialKey := int64(1234)
	data := redis.LoginMetadata{
		FEKey:     []byte("test"),
		PlayerID:  1,
		AccountID: 2,
	}

	// queue login
	is.NoErr(rh.QueueLogin(serialKey, data))

	// get login
	loginData, err := rh.GetLogin(serialKey)
	is.NoErr(err)

	// compare
	is.Equal(loginData, data) // received data should be the same as sent data

	// delete login
	is.NoErr(rh.RemoveLogin(serialKey))

	// get login
	_, err = rh.GetLogin(serialKey)
	is.True(err != nil) // should fail to get removed login
}

func TestRedisShard(t *testing.T) {
	is := is.New(t)

	// test data
	shard := redis.ShardMetadata{
		IP:   "0.0.0.0",
		Port: 1234,
	}

	// register shard
	is.NoErr(rh.RegisterShard(shard))

	// get shards
	shards := rh.GetShards()
	is.True(len(shards) == 1)  // should only be 1 shard
	is.Equal(shards[0], shard) // received data should be the same as sent data
}
