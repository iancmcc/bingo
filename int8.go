package bingo

import (
	"unsafe"

	"github.com/goccy/go-reflect"
)

const (
	typeByteInt8        byte = 0x29
	typeByteInt8Inverse      = typeByteInt8 ^ 0xff
	sizeInt8                 = int(unsafe.Sizeof(int8(0))) + 1
)

func EncodeInt8(b []byte, v int8, inverse bool) int {
	if cap(b) < sizeInt8 {
		panic("slice is too small to hold an int8")
	}
	b = b[:sizeInt8]
	b[0] = typeByteInt8
	b[1] = byte(uint8(v) ^ 1<<7)
	if inverse {
		InvertArraySmall(b)
	}
	return sizeInt8
}

func DecodeInt8(b []byte, v reflect.Value) (int, error) {
	encoded := b[1]
	if b[0] == typeByteInt8Inverse {
		encoded = InvertByte(encoded)
	}
	val := int8((encoded ^ 0x80) & 0xff)
	ptr := v.Pointer()
	**(**int8)(unsafe.Pointer(&ptr)) = *(*int8)(unsafe.Pointer(&val))
	return sizeInt8, nil
}
