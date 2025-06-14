package pool

import (
	"bytes"
	"sync"
)

// BufferPool manages a pool of reusable byte buffers
type BufferPool struct {
	pool sync.Pool
}

// NewBufferPool creates a new buffer pool
func NewBufferPool() *BufferPool {
	return &BufferPool{
		pool: sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
}

// Get retrieves a buffer from the pool
func (p *BufferPool) Get() *bytes.Buffer {
	return p.pool.Get().(*bytes.Buffer)
}

// Put returns a buffer to the pool after resetting it
func (p *BufferPool) Put(buf *bytes.Buffer) {
	buf.Reset()
	p.pool.Put(buf)
}

// StringPool manages a pool of string builders
type StringPool struct {
	pool sync.Pool
}

// NewStringPool creates a new string pool
func NewStringPool() *StringPool {
	return &StringPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 0, 1024)
			},
		},
	}
}

// Get retrieves a byte slice from the pool
func (p *StringPool) Get() []byte {
	return p.pool.Get().([]byte)[:0]
}

// Put returns a byte slice to the pool
func (p *StringPool) Put(b []byte) {
	if cap(b) > 64*1024 {
		// Don't pool very large buffers
		return
	}
	p.pool.Put(b)
}