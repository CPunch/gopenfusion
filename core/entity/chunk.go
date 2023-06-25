package entity

import (
	"log"
	"sync"

	"github.com/CPunch/gopenfusion/config"
)

type ChunkPosition struct {
	X int
	Y int
}

func makeChunkPosition(x, y int) ChunkPosition {
	return ChunkPosition{
		X: x / (config.VIEW_DISTANCE / 3),
		Y: y / (config.VIEW_DISTANCE / 3),
	}
}

type Chunk struct {
	Position ChunkPosition
	entities map[Entity]struct{}
	lock     sync.Mutex
}

func NewChunk(position ChunkPosition) *Chunk {
	return &Chunk{
		Position: position,
		entities: make(map[Entity]struct{}),
	}
}

func (c *Chunk) AddEntity(entity Entity) {
	entity.SetChunk(c)
	c.entities[entity] = struct{}{}
}

func (c *Chunk) RemoveEntity(entity Entity) {
	entity.SetChunk(nil)
	delete(c.entities, entity)
}

// send packet to all peers in this chunk and kill each peer if error
func (c *Chunk) SendPacket(typeID uint32, pkt ...interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	for entity := range c.entities {
		if entity.GetKind() != ENTITY_KIND_PLAYER {
			continue
		}

		plr, ok := entity.(*Player)
		if !ok {
			log.Panic("Chunk.SendPacket: entity kind was player, but is not a *Player")
		}
		peer := plr.Peer

		if err := peer.Send(typeID, pkt...); err != nil {
			peer.Kill()
		}
	}
}

func (c *Chunk) GetAdjacentPositions() []ChunkPosition {
	return []ChunkPosition{
		{c.Position.X - 1, c.Position.Y - 1},
		{c.Position.X - 1, c.Position.Y},
		{c.Position.X - 1, c.Position.Y + 1},
		{c.Position.X, c.Position.Y - 1},
		{c.Position.X, c.Position.Y},
		{c.Position.X, c.Position.Y + 1},
		{c.Position.X + 1, c.Position.Y - 1},
		{c.Position.X + 1, c.Position.Y},
		{c.Position.X + 1, c.Position.Y + 1},
	}
}
