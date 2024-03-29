package shard

import (
	"github.com/CPunch/gopenfusion/shard/entity"
)

func (server *ShardServer) addEntity(e entity.Entity) {
	pos := e.GetChunkPos()
	server.addEntityToChunks(e, server.getViewableChunks(pos))
	server.getChunk(pos).AddEntity(e)
}

func (server *ShardServer) removeEntity(e entity.Entity) {
	// TODO: chunk cleanup
	pos := e.GetChunkPos()
	server.removeEntityFromChunks(e, server.getViewableChunks(pos))
	server.getChunk(pos).RemoveEntity(e)
}

func (server *ShardServer) getChunk(pos entity.ChunkPosition) *entity.Chunk {
	chunk, ok := server.chunks[pos]
	if !ok {
		chunk = entity.NewChunk(pos)
		server.chunks[pos] = chunk
	}

	return chunk
}

func (server *ShardServer) getViewableChunks(pos entity.ChunkPosition) []*entity.Chunk {
	chunks := make([]*entity.Chunk, 0, 9)
	for _, pos := range server.getChunk(pos).GetAdjacentPositions() {
		chunks = append(chunks, server.getChunk(pos))
	}

	return chunks
}

// sends a packet to all peers in the given chunks, excluding the given peer
func (server *ShardServer) sendOthersPacket(plr *entity.Player, typeID uint32, pkt ...interface{}) error {
	chunks := server.getViewableChunks(plr.Chunk)
	for _, chunk := range chunks {
		chunk.SendPacketExclude(plr, typeID, pkt...)
	}

	return nil
}

// sends a packet to all peers in the given chunks
func (server *ShardServer) sendAllPacket(plr *entity.Player, typeID uint32, pkt ...interface{}) error {
	chunks := server.getViewableChunks(plr.Chunk)
	for _, chunk := range chunks {
		chunk.SendPacket(typeID, pkt...)
	}

	return nil
}

func (server *ShardServer) removeEntityFromChunks(this entity.Entity, chunks []*entity.Chunk) {
	for _, chunk := range chunks {
		chunk.ForEachEntity(func(e entity.Entity) bool {
			if e == this {
				return false
			}

			// notify other players we're leaving
			if e.GetKind() == entity.ENTITY_KIND_PLAYER {
				otherPlr := e.(*entity.Player)
				this.DisappearFromViewOf(otherPlr.Peer)
			}

			// notify us they're leaving
			if this.GetKind() == entity.ENTITY_KIND_PLAYER {
				thisPlr := this.(*entity.Player)
				e.DisappearFromViewOf(thisPlr.Peer)
			}

			return false
		})
	}
}

func (server *ShardServer) addEntityToChunks(this entity.Entity, chunks []*entity.Chunk) {
	for _, chunk := range chunks {
		chunk.ForEachEntity(func(e entity.Entity) bool {
			if e == this {
				return false
			}

			// notify other players we're entering
			if e.GetKind() == entity.ENTITY_KIND_PLAYER {
				otherPlr := e.(*entity.Player)
				this.EnterIntoViewOf(otherPlr.Peer)
			}

			// notify us they're entering
			if this.GetKind() == entity.ENTITY_KIND_PLAYER {
				thisPlr := this.(*entity.Player)
				e.EnterIntoViewOf(thisPlr.Peer)
			}

			return false
		})
	}
}

func (server *ShardServer) updateEntityChunk(e entity.Entity, from entity.ChunkPosition, to entity.ChunkPosition) {
	// no change needed
	if from == to {
		return
	}

	oldViewables := server.getViewableChunks(from)
	newViewables := server.getViewableChunks(to)

	// compute differences
	toExit := entity.ChunkSliceDifference(oldViewables, newViewables)
	toEnter := entity.ChunkSliceDifference(newViewables, oldViewables)

	// update chunks
	server.removeEntityFromChunks(e, toExit)
	server.addEntityToChunks(e, toEnter)
	server.getChunk(from).RemoveEntity(e)
	server.getChunk(to).AddEntity(e)
	e.SetChunkPos(to)
}
