package shard

import (
	"github.com/CPunch/gopenfusion/core/entity"
)

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

func (server *ShardServer) sendPacketToViewableChunks(plr *entity.Player, typeID uint32, pkt ...interface{}) error {
	for _, chunk := range server.getViewableChunks(plr.Chunk) {
		chunk.SendPacket(typeID, pkt...)
	}

	return nil
}

func (server *ShardServer) removeEntityFromChunks(chunks []*entity.Chunk, this entity.Entity) {
	for _, chunk := range chunks {
		for e, _ := range chunk.Entities {
			if e.GetKind() == entity.ENTITY_KIND_PLAYER {
				otherPlr := e.(*entity.Player)
				this.DisappearFromViewOf(otherPlr.Peer)
			}

			if this.GetKind() == entity.ENTITY_KIND_PLAYER {
				thisPlr := this.(*entity.Player)
				e.DisappearFromViewOf(thisPlr.Peer)
			}
		}
	}
}

func (server *ShardServer) addEntityToChunks(chunks []*entity.Chunk, this entity.Entity) {
	for _, chunk := range chunks {
		for e, _ := range chunk.Entities {
			if e.GetKind() == entity.ENTITY_KIND_PLAYER {
				otherPlr := e.(*entity.Player)
				this.EnterIntoViewOf(otherPlr.Peer)
			}

			if this.GetKind() == entity.ENTITY_KIND_PLAYER {
				thisPlr := this.(*entity.Player)
				e.EnterIntoViewOf(thisPlr.Peer)
			}
		}
	}
}

func (server *ShardServer) updateEntityChunk(e entity.Entity, from entity.ChunkPosition, to entity.ChunkPosition) {
	oldViewables := server.getViewableChunks(from)
	newViewables := server.getViewableChunks(to)

	// compute differences
	toExit := entity.ChunkSliceDifference(oldViewables, newViewables)
	toEnter := entity.ChunkSliceDifference(newViewables, oldViewables)

	// update chunks
	server.removeEntityFromChunks(toExit, e)
	server.addEntityToChunks(toEnter, e)
	server.getChunk(from).RemoveEntity(e)
	server.getChunk(to).AddEntity(e)
	e.SetChunk(to)
}
