package codecs

import (
	"unsafe"

	"github.com/goccy/go-reflect"
	"github.com/iancmcc/bingo/bytes"
)

const (
	typeByteBool        byte = 0x04
	typeByteBoolInverse      = typeByteBool ^ 0xff
	sizeBool                 = 2
)

func encodeBool(b []byte, v bool, inverse bool) (int, error) {
	if cap(b) < sizeBool {
		return 0, ErrByteSliceSize
	}
	b = b[:sizeBool]
	b[0] = typeByteBool
	if v {
		b[1] = 0xff
	}
	if inverse {
		bytes.InvertArraySmall(b)
	}
	return sizeBool, nil
}

func decodeBool(b []byte, v reflect.Value) (int, error) {
	val := b[1] == 0xff
	if b[0] == typeByteBoolInverse {
		val = !val
	}
	ptr := v.Pointer()
	**(**bool)(unsafe.Pointer(&ptr)) = *(*bool)(unsafe.Pointer(&val))
	return sizeBool, nil
}
