package internal

import (
	"encoding/binary"
	"reflect"
	"unsafe"
)

const (
	typeByteInt32        byte = 0x2b
	typeByteInt32Inverse      = typeByteInt32 ^ 0xff
	sizeInt32                 = int(unsafe.Sizeof(int32(0))) + 1
)

var (
	intType = reflect.TypeOf(int(0))
)

func EncodeInt32(b []byte, v int32, inverse bool) {
	if len(b) < sizeInt32 {
		panic("slice is too small to hold an int32")
	}
	b[0] = typeByteInt32
	binary.BigEndian.PutUint32(b[1:], uint32(v)^(1<<31))
	if inverse {
		invertArray(b)
	}
}

func DecodeInt32(b []byte, v reflect.Value) (int, error) {
	encoded := b[1:5]
	if b[0] == typeByteInt32Inverse {
		encoded = make([]byte, 4)
		copy(encoded, b[1:5])
		invertArray(encoded)
	}
	val := int32((encoded[0] ^ 0x80) & 0xff)
	for i := 1; i < 4; i++ {
		val = (val << 8) + int32(encoded[i]&0xff)
	}
	e := v.Elem()
	if e.Type() == intType {
		e.Set(reflect.ValueOf(int(val)))
	} else {
		e.Set(reflect.ValueOf(val))
	}
	return sizeInt32, nil
}
