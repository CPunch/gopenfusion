package config

import (
	"os"
	"time"
)

/*
	Available environment variables:
		REDIS_ADDR
		REDIS_USER
		REDIS_PASS
		DB_ADDR
		DB_USER
		DB_PASS
		ANNOUNCE_IP
*/

const (
	AEQUIP_COUNT = 9
	AINVEN_COUNT = 50
	ABANK_COUNT  = 119

	NANO_COUNT = 37
)

var (
	SPAWN_X = 632032
	SPAWN_Y = 187177
	SPAWN_Z = -5500

	LOGIN_PORT = 23000
	SHARD_PORT = 23001

	LOGIN_TIMEOUT = time.Second * 30
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func GetMaxHP(level int) int {
	return (925 + 75*(level))
}

func GetRedisAddr() string {
	return getEnv("REDIS_ADDR", "localhost:6379")
}

func GetRedisCredentials() func() (string, string) {
	return func() (string, string) {
		return getEnv("REDIS_USER", ""), getEnv("REDIS_PASS", "")
	}
}

func GetDBName() string {
	return getEnv("DB_NAME", "postgres")
}

func GetDBAddr() string {
	return getEnv("DB_ADDR", "localhost:5432")
}

func GetDBUser() string {
	return getEnv("DB_USER", "postgres")
}

func GetDBPass() string {
	return getEnv("DB_PASS", "")
}

// needed for shard
func GetAnnounceIP() string {
	return getEnv("ANNOUNCE_IP", "127.0.0.1")
}
