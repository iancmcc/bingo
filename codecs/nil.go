package codecs

import "github.com/iancmcc/bingo/bytes"

const (
	typeByteNil        = 0x05
	typeByteNilInverse = typeByteNil ^ 0xff
	sizeNil            = 1
)

func EncodeNil(b []byte, inverse bool) int {
	b = b[:1]
	b[0] = typeByteNil
	if inverse {
		bytes.InvertArraySmall(b)
	}
	return sizeNil
}

func DecodeNil(b []byte) int {
	return sizeNil
}
