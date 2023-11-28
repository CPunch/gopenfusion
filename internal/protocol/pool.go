package protocol

import (
	"bytes"
	"sync"
)

var allocator = &sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

// grabs a *bytes.Buffer from the pool
func GetBuffer() *bytes.Buffer {
	return allocator.Get().(*bytes.Buffer)
}

// returns a *bytes.Buffer to the pool
func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	allocator.Put(buf)
}
