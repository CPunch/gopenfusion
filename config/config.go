package config

var (
	SPAWN_X = 632032
	SPAWN_Y = 187177
	SPAWN_Z = -5500

	AEQUIP_COUNT = 9
	AINVEN_COUNT = 50
	ABANK_COUNT  = 119
)

func GetMaxHP(level int) int {
	return (925 + 75*(level))
}
