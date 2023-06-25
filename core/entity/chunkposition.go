package entity

import "github.com/CPunch/gopenfusion/config"

type ChunkPosition struct {
	X int
	Y int
}

func MakeChunkPosition(x, y int) ChunkPosition {
	return ChunkPosition{
		X: x / (config.VIEW_DISTANCE / 3),
		Y: y / (config.VIEW_DISTANCE / 3),
	}
}
