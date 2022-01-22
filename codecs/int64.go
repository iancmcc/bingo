package codecs

import (
	"encoding/binary"
	"unsafe"

	"github.com/goccy/go-reflect"
	"github.com/iancmcc/bingo/bytes"
)

const (
	typeByteInt64        byte = 0x2c
	typeByteInt64Inverse      = typeByteInt64 ^ 0xff
	sizeInt64                 = int(unsafe.Sizeof(int64(0))) + 1
)

func encodeInt64(b []byte, v int64, inverse bool) (int, error) {
	if cap(b) < sizeInt64 {
		return 0, ErrByteSliceSize
	}
	b = b[:sizeInt64]
	b[0] = typeByteInt64
	binary.BigEndian.PutUint64(b[1:], uint64(v)^(1<<63))
	if inverse {
		bytes.InvertArraySmall(b)
	}
	return sizeInt64, nil
}

func decodeInt64(b []byte, v reflect.Value) (int, error) {
	encoded := b[1:9]
	if b[0] == typeByteInt64Inverse {
		encoded = make([]byte, 8)
		copy(encoded, b[1:9])
		bytes.InvertArraySmall(encoded)
	}
	val := int64((encoded[0] ^ 0x80) & 0xff)
	for i := 1; i < 8; i++ {
		val = (val << 8) + int64(encoded[i]&0xff)
	}
	ptr := v.Pointer()
	**(**int64)(unsafe.Pointer(&ptr)) = *(*int64)(unsafe.Pointer(&val))
	return sizeInt64, nil
}
