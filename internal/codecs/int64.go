package codecs

import (
	"encoding/binary"
	"reflect"
	"unsafe"
)

const (
	typeByteInt64        byte = 0x2c
	typeByteInt64Inverse      = typeByteInt64 ^ 0xff
	sizeInt64                 = int(unsafe.Sizeof(int64(0))) + 1
)

func EncodeInt64(b []byte, v int64, inverse bool) int {
	if cap(b) < sizeInt64 {
		panic("slice is too small to hold an int8")
	}
	b = b[:sizeInt64]
	b[0] = typeByteInt64
	binary.BigEndian.PutUint64(b[1:], uint64(v)^(1<<63))
	if inverse {
		invertArray(b)
	}
	return sizeInt64
}

func DecodeInt64(b []byte, v reflect.Value) (int, error) {
	encoded := b[1:9]
	if b[0] == typeByteInt64Inverse {
		encoded = make([]byte, 8)
		copy(encoded, b[1:9])
		invertArray(encoded)
	}
	val := int64((encoded[0] ^ 0x80) & 0xff)
	for i := 1; i < 8; i++ {
		val = (val << 8) + int64(encoded[i]&0xff)
	}
	v.Elem().Set(reflect.ValueOf(val))
	return sizeInt64, nil
}
