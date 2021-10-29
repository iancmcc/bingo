package decoder

import (
	"unsafe"

	"github.com/iancmcc/bingo/internal/bytes"
	"github.com/iancmcc/bingo/internal/types"
)

func DecodeInt8(b []byte, ptr uintptr) (int, error) {
	encoded := b[1]
	if b[0] == types.TypeByteInt8Inverse {
		encoded = bytes.InvertByte(encoded)
	}
	val := int8((encoded ^ 0x80) & 0xff)
	**(**int8)(unsafe.Pointer(&ptr)) = *(*int8)(unsafe.Pointer(&val))
	return types.SizeInt8, nil
}
