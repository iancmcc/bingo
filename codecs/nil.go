package codecs

const (
	typeByteNil        = 0x05
	typeByteNilInverse = typeByteNil ^ 0xff
	sizeNil            = 1
)

func encodeNil(b []byte, inverse bool) (int, error) {
	if cap(b) < sizeNil {
		return 0, ErrByteSliceSize
	}
	b = b[:1]
	if inverse {
		b[0] = typeByteNilInverse
	} else {
		b[0] = typeByteNil
	}
	return sizeNil, nil
}
