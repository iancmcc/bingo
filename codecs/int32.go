package codecs

import (
	"encoding/binary"
	"unsafe"

	"github.com/goccy/go-reflect"
	"github.com/iancmcc/bingo/bytes"
)

const (
	typeByteInt32        byte = 0x2b
	typeByteInt32Inverse      = typeByteInt32 ^ 0xff
	sizeInt32                 = int(unsafe.Sizeof(int32(0))) + 1
)

func encodeInt32(b []byte, v int32, inverse bool) (int, error) {
	if cap(b) < sizeInt32 {
		return 0, ErrByteSliceSize
	}
	b = b[:sizeInt32]
	b[0] = typeByteInt32
	binary.BigEndian.PutUint32(b[1:], uint32(v)^(1<<31))
	if inverse {
		bytes.InvertArraySmall(b)
	}
	return sizeInt32, nil
}

func decodeInt32(b []byte, v reflect.Value) (int, error) {
	encoded := b[1:5]
	if b[0] == typeByteInt32Inverse {
		encoded = make([]byte, 4)
		copy(encoded, b[1:5])
		bytes.InvertArraySmall(encoded)
	}
	val := int32((encoded[0] ^ 0x80) & 0xff)
	for i := 1; i < 4; i++ {
		val = (val << 8) + int32(encoded[i]&0xff)
	}
	uptr := v.Pointer()
	**(**int32)(unsafe.Pointer(&uptr)) = *(*int32)(unsafe.Pointer(&val))
	return sizeInt32, nil
}
