package internal

import (
	"encoding/binary"
	"math"
	"reflect"
	"unsafe"
)

const (
	typeByteFloat32        byte = 0x30
	typeByteFloat32Inverse      = typeByteFloat32 ^ 0xff
	sizeFloat32                 = int(unsafe.Sizeof(float32(0))) + 1
)

func EncodeFloat32(b []byte, v float32, inverse bool) {
	if len(b) < sizeFloat32 {
		panic("slice is too small to hold a float32")
	}
	b[0] = typeByteFloat32
	int32Val := int32(math.Float32bits(v))
	int32Val ^= (int32Val >> 31) | (-1 << 31)
	binary.BigEndian.PutUint32(b[1:], uint32(int32Val))
	if inverse {
		invertArray(b)
	}
}

func DecodeFloat32(b []byte, v reflect.Value) (int, error) {
	encoded := b[1:5]
	if b[0] == typeByteFloat32Inverse {
		encoded = make([]byte, 4)
		copy(encoded, b[1:5])
		invertArray(encoded)
	}
	val := int32(binary.BigEndian.Uint32(encoded))
	val ^= (^val >> 31) | (-1 << 31)
	fv := math.Float32frombits(uint32(val))
	v.Elem().Set(reflect.ValueOf(fv))
	return sizeFloat32, nil
}
