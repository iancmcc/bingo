package codecs

import (
	"encoding/binary"
	"unsafe"

	"github.com/goccy/go-reflect"
	"github.com/iancmcc/bingo/bytes"
)

const (
	typeByteUint16        byte = 0x1a
	typeByteUint16Inverse      = typeByteUint16 ^ 0xff
	sizeUint16                 = int(unsafe.Sizeof(uint16(0))) + 1
)

func encodeUint16(b []byte, v uint16, inverse bool) (int, error) {
	if cap(b) < sizeUint16 {
		return 0, ErrByteSliceSize
	}
	b = b[:sizeUint16]
	b[0] = typeByteUint16
	binary.BigEndian.PutUint16(b[1:], uint16(v)^(1<<15))
	if inverse {
		bytes.InvertArraySmall(b)
	}
	return sizeUint16, nil
}

func decodeUint16(b []byte, v reflect.Value) (int, error) {
	encoded := b[1:3]
	if b[0] == typeByteUint16Inverse {
		encoded = make([]byte, 2)
		copy(encoded, b[1:3])
		bytes.InvertArraySmall(encoded)
	}
	val := uint16((encoded[0] ^ 0x80) & 0xff)
	val = (val << 8) + uint16(encoded[1]&0xff)

	ptr := v.Pointer()
	**(**uint16)(unsafe.Pointer(&ptr)) = *(*uint16)(unsafe.Pointer(&val))
	return sizeUint16, nil
}
