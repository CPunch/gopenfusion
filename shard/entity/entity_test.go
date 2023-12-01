package entity_test

import (
	"testing"

	"github.com/CPunch/gopenfusion/shard/entity"
	"github.com/matryer/is"
)

func TestChunkSliceDifference(t *testing.T) {
	is := is.New(t)
	chunks := []*entity.Chunk{
		entity.NewChunk(entity.MakeChunkPosition(0, 0)),
		entity.NewChunk(entity.MakeChunkPosition(0, 1)),
		entity.NewChunk(entity.MakeChunkPosition(1, 0)),
		entity.NewChunk(entity.MakeChunkPosition(1, 1)),
	}

	c1 := []*entity.Chunk{
		chunks[0],
		chunks[1],
		chunks[2],
		chunks[3],
	}

	c2 := []*entity.Chunk{
		chunks[0],
		chunks[1],
		chunks[2],
	}

	diff := entity.ChunkSliceDifference(c1, c2)

	is.True(len(diff) == 1)       // should be 1 chunk in difference
	is.True(diff[0] == chunks[3]) // should be chunks[3] in difference
}
