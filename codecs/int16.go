package codecs

import (
	"encoding/binary"
	"unsafe"

	"github.com/goccy/go-reflect"
	"github.com/iancmcc/bingo/bytes"
)

const (
	typeByteInt16        byte = 0x2a
	typeByteInt16Inverse      = typeByteInt16 ^ 0xff
	sizeInt16                 = int(unsafe.Sizeof(int16(0))) + 1
)

func encodeInt16(b []byte, v int16, inverse bool) (int, error) {
	if cap(b) < sizeInt16 {
		return 0, ErrByteSliceSize
	}
	b = b[:sizeInt16]
	b[0] = typeByteInt16
	binary.BigEndian.PutUint16(b[1:], uint16(v)^(1<<15))
	if inverse {
		bytes.InvertArraySmall(b)
	}
	return sizeInt16, nil
}

func decodeInt16(b []byte, v reflect.Value) (int, error) {
	encoded := b[1:3]
	if b[0] == typeByteInt16Inverse {
		encoded = make([]byte, 2)
		copy(encoded, b[1:3])
		bytes.InvertArraySmall(encoded)
	}
	val := int16((encoded[0] ^ 0x80) & 0xff)
	val = (val << 8) + int16(encoded[1]&0xff)

	ptr := v.Pointer()
	**(**int16)(unsafe.Pointer(&ptr)) = *(*int16)(unsafe.Pointer(&val))
	return sizeInt16, nil
}
