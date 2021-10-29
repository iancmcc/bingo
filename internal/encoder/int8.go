package encoder

import (
	"github.com/iancmcc/bingo/internal/bytes"
	"github.com/iancmcc/bingo/internal/errors"
	"github.com/iancmcc/bingo/internal/types"
)

func EncodeInt8(b []byte, v interface{}, inverse bool) (int, error) {
	if cap(b) < types.SizeInt8 {
		return 0, errors.ErrArrayTooSmall
	}
	b = b[:types.SizeInt8]
	b[0] = types.TypeByteInt8
	b[1] = byte(uint8(v.(int8)) ^ 1<<7)
	if inverse {
		bytes.InvertArraySmall(b)
	}
	return types.SizeInt8, nil
}
