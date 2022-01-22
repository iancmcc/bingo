package codecs

import (
	"strings"
	"unsafe"

	"github.com/goccy/go-reflect"
	"github.com/iancmcc/bingo/bytes"
)

const (
	typeByteString        byte = 0x34
	terminatorByte        byte = 0x00
	typeByteStringInverse      = typeByteString ^ 0xff
	terminatorByteInverse      = terminatorByte ^ 0xff
)

func encodeString(b []byte, v string, inverse bool) (int, error) {
	size := len(v) + 2
	if cap(b) < size {
		return 0, ErrByteSliceSize
	}
	// Apparently an upfront check + copy is faster than checking during range
	if strings.IndexByte(v, terminatorByte) > -1 {
		return 0, ErrNullByte
	}
	b = b[:size]
	b[0] = typeByteString
	copy(b[1:len(v)+1], v)
	b[len(v)+1] = terminatorByte
	if inverse {
		bytes.InvertArray(b)
	}
	return size, nil
}

func decodeString(b []byte, v reflect.Value) (int, error) {
	var (
		encoded []byte
	)
	idx, err := SizeNext(b)
	if err != nil {
		return 0, err
	}
	encoded = b[1 : idx-1]

	if b[0] == typeByteStringInverse {
		bytes.InvertArray(encoded)
	}

	ptr := v.Pointer()
	**(**string)(unsafe.Pointer(&ptr)) = *(*string)(unsafe.Pointer(&encoded))

	return idx, nil
}
