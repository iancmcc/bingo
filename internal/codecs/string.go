package codecs

import (
	"reflect"
	"strings"
	"unsafe"
)

const (
	typeByteString        byte = 0x34
	terminatorByte        byte = 0x00
	typeByteStringInverse      = typeByteString ^ 0xff
	terminatorByteInverse      = terminatorByte ^ 0xff
)

func EncodeString(b []byte, v string, inverse bool) {
	if len(b) < len(v)+2 {
		panic("slice is too small to hold encoded string")
	}
	// Apparently an upfront check + copy is faster than checking during range
	if strings.IndexByte(v, terminatorByte) > -1 {
		panic("can't encode a string that contains a null byte")
	}
	b[0] = typeByteString
	copy(b[1:len(v)+1], stringToByte(v))
	b[len(v)+1] = terminatorByte
	if inverse {
		invertArray(b)
	}
}

func DecodeString(b []byte, v reflect.Value) (int, error) {
	var (
		encoded []byte
		idx     = SizeNext(b)
	)
	encoded = b[1:idx]

	if b[0] == typeByteStringInverse {
		invertArray(encoded)
	}
	v.Elem().Set(reflect.ValueOf(string(encoded)))
	return idx + 1, nil
}

func stringToByte(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(sh))
}
