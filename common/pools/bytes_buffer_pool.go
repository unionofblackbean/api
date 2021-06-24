package pools

import (
	"bytes"
	"sync"
)

var bytesBufferPool sync.Pool

func GetBytesBuffer() *bytes.Buffer {
	if buf, ok := bytesBufferPool.Get().(*bytes.Buffer); ok {
		buf.Reset()
		return buf
	} else {
		return new(bytes.Buffer)
	}
}

func PutBytesBuffer(buffer *bytes.Buffer) {
	bytesBufferPool.Put(buffer)
}
