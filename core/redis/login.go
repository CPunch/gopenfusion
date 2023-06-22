package redis

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/CPunch/gopenfusion/config"
)

type LoginMetadata struct {
	FEKey    []byte `json:",omitempty"`
	PlayerID int32  `json:",omitempty"`
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
