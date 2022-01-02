package bingo

import (
	"encoding/binary"
	"math"
	"unsafe"

	"github.com/goccy/go-reflect"
)

const (
	typeByteFloat64        byte = 0x31
	typeByteFloat64Inverse      = typeByteFloat64 ^ 0xff
	sizeFloat64                 = int(unsafe.Sizeof(float64(0))) + 1
)

func EncodeFloat64(b []byte, v float64, inverse bool) int {
	if cap(b) < sizeFloat64 {
		panic("slice is too small to hold a float64")
	}
	b = b[:sizeFloat64]
	int64Val := int64(math.Float64bits(v))
	int64Val ^= (int64Val >> 63) | (-1 << 63)
	b[0] = typeByteFloat64
	binary.BigEndian.PutUint64(b[1:], uint64(int64Val))
	if inverse {
		InvertArraySmall(b)
	}
	return sizeFloat64
}

func DecodeFloat64(b []byte, v reflect.Value) (int, error) {
	encoded := b[1:9]
	if b[0] == typeByteFloat64Inverse {
		encoded = make([]byte, 8)
		copy(encoded, b[1:9])
		InvertArraySmall(encoded)
	}
	val := int64(binary.BigEndian.Uint64(encoded))
	val ^= (^val >> 63) | (-1 << 63)
	fv := math.Float64frombits(uint64(val))
	ptr := v.Pointer()
	**(**float64)(unsafe.Pointer(&ptr)) = *(*float64)(unsafe.Pointer(&fv))
	return sizeFloat64, nil
}
