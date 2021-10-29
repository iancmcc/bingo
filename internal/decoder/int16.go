package decoder

import (
	"unsafe"

	"github.com/iancmcc/bingo/internal/bytes"
	"github.com/iancmcc/bingo/internal/types"
)

func DecodeInt16(b []byte, ptr uintptr) (int, error) {
	encoded := b[1:3]
	if b[0] == types.TypeByteInt16Inverse {
		encoded = make([]byte, 2)
		copy(encoded, b[1:3])
		bytes.InvertArraySmall(encoded)
	}
	val := int16((encoded[0] ^ 0x80) & 0xff)
	val = (val << 8) + int16(encoded[1]&0xff)
	**(**int16)(unsafe.Pointer(ptr)) = *(*int16)(unsafe.Pointer(&val))
	return types.SizeInt16, nil
}
