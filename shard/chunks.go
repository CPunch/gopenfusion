package shard

import "github.com/CPunch/gopenfusion/core/entity"

func (server *ShardServer) getChunk(pos entity.ChunkPosition) *entity.Chunk {
	chunk, ok := server.chunks[pos]
	if !ok {
		chunk = entity.NewChunk(pos)
		server.chunks[pos] = chunk
	}

	return chunk
}

func (server *ShardServer) getViewableChunks(plr *entity.Player) []*entity.Chunk {
	chunks := make([]*entity.Chunk, 0, 9)

	for _, pos := range plr.GetChunk().GetAdjacentPositions() {
		chunks = append(chunks, server.getChunk(pos))
	}

	return chunks
}
