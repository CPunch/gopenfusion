package entity

import (
	"log"
	"sync"
)

type Chunk struct {
	Position ChunkPosition
	Entities map[Entity]struct{}
	lock     sync.Mutex
}

func NewChunk(position ChunkPosition) *Chunk {
	return &Chunk{
		Position: position,
		Entities: make(map[Entity]struct{}),
	}
}

func (c *Chunk) AddEntity(entity Entity) {
	c.Entities[entity] = struct{}{}
}

func (c *Chunk) RemoveEntity(entity Entity) {
	delete(c.Entities, entity)
}

// send packet to all peers in this chunk and kill each peer if error
func (c *Chunk) SendPacket(typeID uint32, pkt ...interface{}) {
	c.SendPacketExclude(nil, typeID, pkt...)
}

func (c *Chunk) SendPacketExclude(exclude Entity, typeID uint32, pkt ...interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	for entity := range c.Entities {
		// only send to players, and exclude the player that sent the packet
		if entity.GetKind() != ENTITY_KIND_PLAYER || entity == exclude {
			continue
		}

		plr, ok := entity.(*Player)
		if !ok {
			log.Panic("Chunk.SendPacket: entity kind was player, but is not a *Player")
		}
		peer := plr.Peer

		if err := peer.Send(typeID, pkt...); err != nil {
			log.Printf("Error sending packet to peer %p: %v", peer, err)
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

// https://stackoverflow.com/a/45428032 lol
func ChunkSliceDifference(a, b []*Chunk) []*Chunk {
	m := make(map[*Chunk]struct{})
	for _, item := range b {
		m[item] = struct{}{}
	}

	var diff []*Chunk
	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}

	return diff
}
