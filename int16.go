package bingo

import (
	"encoding/binary"
	"unsafe"

	"github.com/goccy/go-reflect"
)

const (
	typeByteInt16        byte = 0x2a
	typeByteInt16Inverse      = typeByteInt16 ^ 0xff
	sizeInt16                 = int(unsafe.Sizeof(int16(0))) + 1
)

func EncodeInt16(b []byte, v int16, inverse bool) int {
	if cap(b) < sizeInt16 {
		panic("slice is too small to hold an int16")
	}
	b = b[:sizeInt16]
	b[0] = typeByteInt16
	binary.BigEndian.PutUint16(b[1:], uint16(v)^(1<<15))
	if inverse {
		InvertArraySmall(b)
	}
	return sizeInt16
}

func DecodeInt16(b []byte, v reflect.Value) (int, error) {
	encoded := b[1:3]
	if b[0] == typeByteInt16Inverse {
		encoded = make([]byte, 2)
		copy(encoded, b[1:3])
		InvertArraySmall(encoded)
	}
	val := int16((encoded[0] ^ 0x80) & 0xff)
	val = (val << 8) + int16(encoded[1]&0xff)

	ptr := v.Pointer()
	**(**int16)(unsafe.Pointer(&ptr)) = *(*int16)(unsafe.Pointer(&val))
	return sizeInt16, nil
}
