package entity_test

import (
	"testing"

	"github.com/CPunch/gopenfusion/internal/entity"
)

func TestChunkSliceDifference(t *testing.T) {
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
	if len(diff) != 1 {
		t.Logf("%+v", diff)
		t.Error("expected 1 chunk in difference")
	}

	if diff[0] != chunks[3] {
		t.Logf("%+v", diff)
		t.Error("wrong difference")
	}
}
