package pool

import (
	"bytes"
	"sync"
)

var allocator = &sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

func Get() *bytes.Buffer {
	return allocator.Get().(*bytes.Buffer)
}

func Put(buf *bytes.Buffer) {
	buf.Reset()
	allocator.Put(buf)
}
