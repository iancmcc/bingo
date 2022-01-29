package codecs

import (
	"unsafe"

	"github.com/goccy/go-reflect"
	"github.com/iancmcc/bingo/bytes"
)

const (
	typeByteUint8        byte = 0x19
	typeByteUint8Inverse      = typeByteUint8 ^ 0xff
	sizeUint8                 = int(unsafe.Sizeof(uint8(0))) + 1
)

func encodeUint8(b []byte, v uint8, inverse bool) (int, error) {
	if cap(b) < sizeUint8 {
		return 0, ErrByteSliceSize
	}
	b = b[:sizeUint8]
	b[0] = typeByteUint8
	b[1] = byte(uint8(v) ^ 1<<7)
	if inverse {
		bytes.InvertArraySmall(b)
	}
	return sizeUint8, nil
}

func decodeUint8(b []byte, v reflect.Value) (int, error) {
	encoded := b[1]
	if b[0] == typeByteUint8Inverse {
		encoded = bytes.InvertByte(encoded)
	}
	val := uint8((encoded ^ 0x80) & 0xff)
	ptr := v.Pointer()
	**(**uint8)(unsafe.Pointer(&ptr)) = *(*uint8)(unsafe.Pointer(&val))
	return sizeUint8, nil
}
