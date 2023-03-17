package config

var (
	SPAWN_X = 632032
	SPAWN_Y = 187177
	SPAWN_Z = -5500

	AEQUIP_COUNT = 9
	AINVEN_COUNT = 50
	ABANK_COUNT  = 119

	LOGIN_PORT = 23000
	SHARD_PORT = 23001
	SHARD_IP   = "127.0.0.1"
)

func GetMaxHP(level int) int {
	return (925 + 75*(level))
}
