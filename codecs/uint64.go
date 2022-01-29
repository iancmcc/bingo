package codecs

import (
	"encoding/binary"
	"unsafe"

	"github.com/goccy/go-reflect"
	"github.com/iancmcc/bingo/bytes"
)

const (
	typeByteUint64        byte = 0x1c
	typeByteUint64Inverse      = typeByteUint64 ^ 0xff
	sizeUint64                 = int(unsafe.Sizeof(uint64(0))) + 1
)

func encodeUint64(b []byte, v uint64, inverse bool) (int, error) {
	if cap(b) < sizeUint64 {
		return 0, ErrByteSliceSize
	}
	b = b[:sizeUint64]
	b[0] = typeByteUint64
	binary.BigEndian.PutUint64(b[1:], uint64(v)^(1<<63))
	if inverse {
		bytes.InvertArraySmall(b)
	}
	return sizeUint64, nil
}

func decodeUint64(b []byte, v reflect.Value) (int, error) {
	encoded := b[1:9]
	if b[0] == typeByteUint64Inverse {
		encoded = make([]byte, 8)
		copy(encoded, b[1:9])
		bytes.InvertArraySmall(encoded)
	}
	val := uint64((encoded[0] ^ 0x80) & 0xff)
	for i := 1; i < 8; i++ {
		val = (val << 8) + uint64(encoded[i]&0xff)
	}
	ptr := v.Pointer()
	**(**uint64)(unsafe.Pointer(&ptr)) = *(*uint64)(unsafe.Pointer(&val))
	return sizeUint64, nil
}
