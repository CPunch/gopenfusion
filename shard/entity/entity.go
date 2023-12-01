package entity

import "github.com/CPunch/gopenfusion/cnpeer"

type EntityKind int

const (
	ENTITY_KIND_PLAYER EntityKind = iota
	ENTITY_KIND_NPC
)

type Entity interface {
	GetKind() EntityKind

	GetChunkPos() ChunkPosition
	GetPosition() (x int, y int, z int)
	GetAngle() int

	SetChunkPos(chunk ChunkPosition)
	SetPosition(x, y, z int)
	SetAngle(angle int)

	DisappearFromViewOf(peer *cnpeer.CNPeer)
	EnterIntoViewOf(peer *cnpeer.CNPeer)
}
