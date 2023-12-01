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

func MakeChunkPosition(x, y int) ChunkPosition {
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
	c.lock.Lock()
	defer c.lock.Unlock()

	c.entities[entity] = struct{}{}
}

func (c *Chunk) RemoveEntity(entity Entity) {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.entities, entity)
}

// send packet to all peers in this chunk and kill each peer if error
func (c *Chunk) SendPacket(typeID uint32, pkt ...interface{}) {
	c.SendPacketExclude(nil, typeID, pkt...)
}

// calls f for each entity in this chunk, if f returns true, stop iterating
// f can safely add/remove entities from the chunk
func (c *Chunk) ForEachEntity(f func(entity Entity) bool) {
	for entity := range c.entities {
		if f(entity) {
			break
		}
	}
}

func (c *Chunk) SendPacketExclude(exclude Entity, typeID uint32, pkt ...interface{}) {
	c.ForEachEntity(func(entity Entity) bool {
		// only send to players, and exclude the player that sent the packet
		if entity.GetKind() != ENTITY_KIND_PLAYER || entity == exclude {
			return false
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

		return false
	})
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
