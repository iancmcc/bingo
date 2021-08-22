package internal

const (
	typeByteNil        = 0x05
	typeByteNilInverse = typeByteNil ^ 0xff
	sizeNil            = 1
)

func EncodeNil(b []byte, inverse bool) {
	b[0] = typeByteNil
	if inverse {
		invertArray(b)
	}
}

func DecodeNil(b []byte) int {
	return sizeNil
}
