package codecs

import (
	"encoding/binary"
	"unsafe"

	"github.com/goccy/go-reflect"
	"github.com/iancmcc/bingo/bytes"
)

const (
	typeByteUint32        byte = 0x1b
	typeByteUint32Inverse      = typeByteUint32 ^ 0xff
	sizeUint32                 = int(unsafe.Sizeof(uint32(0))) + 1
)

func encodeUint32(b []byte, v uint32, inverse bool) (int, error) {
	if cap(b) < sizeUint32 {
		return 0, ErrByteSliceSize
	}
	b = b[:sizeUint32]
	b[0] = typeByteUint32
	binary.BigEndian.PutUint32(b[1:], uint32(v)^(1<<31))
	if inverse {
		bytes.InvertArraySmall(b)
	}
	return sizeUint32, nil
}

func decodeUint32(b []byte, v reflect.Value) (int, error) {
	encoded := b[1:5]
	if b[0] == typeByteUint32Inverse {
		encoded = make([]byte, 4)
		copy(encoded, b[1:5])
		bytes.InvertArraySmall(encoded)
	}
	val := uint32((encoded[0] ^ 0x80) & 0xff)
	for i := 1; i < 4; i++ {
		val = (val << 8) + uint32(encoded[i]&0xff)
	}
	uptr := v.Pointer()
	**(**uint32)(unsafe.Pointer(&uptr)) = *(*uint32)(unsafe.Pointer(&val))
	return sizeUint32, nil
}
