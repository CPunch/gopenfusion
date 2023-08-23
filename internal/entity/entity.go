package entity

import "github.com/CPunch/gopenfusion/internal/protocol"

type EntityKind int

const (
	ENTITY_KIND_PLAYER EntityKind = iota
	ENTITY_KIND_NPC
)

type Entity interface {
	GetKind() EntityKind

	GetChunk() ChunkPosition
	GetPosition() (x int, y int, z int)
	GetAngle() int

	SetChunk(chunk ChunkPosition)
	SetPosition(x, y, z int)
	SetAngle(angle int)

	DisappearFromViewOf(peer *protocol.CNPeer)
	EnterIntoViewOf(peer *protocol.CNPeer)
}
