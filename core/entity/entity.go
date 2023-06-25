package entity

type EntityKind int

const (
	ENTITY_KIND_PLAYER EntityKind = iota
	ENTITY_KIND_NPC
)

type Entity interface {
	GetKind() EntityKind

	GetChunk() *Chunk
	GetPosition() (x int, y int, z int)
	GetAngle() int

	SetChunk(chunk *Chunk)
	SetPosition(x, y, z int)
	SetAngle(angle int)
}
