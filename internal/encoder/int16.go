package encoder

import (
	"encoding/binary"

	"github.com/iancmcc/bingo/internal/bytes"
	"github.com/iancmcc/bingo/internal/errors"
	"github.com/iancmcc/bingo/internal/types"
)

func EncodeInt16(b []byte, v interface{}, inverse bool) (int, error) {
	if cap(b) < types.SizeInt16 {
		return 0, errors.ErrArrayTooSmall
	}
	b = b[:types.SizeInt16]
	b[0] = types.TypeByteInt16
	binary.BigEndian.PutUint16(b[1:], uint16(v.(int16))^(1<<15))
	if inverse {
		bytes.InvertArraySmall(b)
	}
	return types.SizeInt16, nil
}
