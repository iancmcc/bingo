package types

import (
	"time"

	"github.com/iancmcc/bingo/internal/errors"
)

const (
	TerminatorByte  byte = 0x00
	TypeByteBool         = 0x04
	TypeByteNil          = 0x05
	TypeByteUint8        = 0x19
	TypeByteInt8         = 0x29
	TypeByteUint16       = 0x1a
	TypeByteInt16        = 0x2a
	TypeByteUint32       = 0x1b
	TypeByteInt32        = 0x2b
	TypeByteUint64       = 0x1c
	TypeByteInt64        = 0x2c
	TypeByteFloat32      = 0x30
	TypeByteFloat64      = 0x31
	TypeByteString       = 0x34
	TypeByteTime         = 0x35

	TerminatorByteInverse  = TerminatorByte ^ 0xff
	TypeByteBoolInverse    = TypeByteBool ^ 0xff
	TypeByteNilInverse     = TypeByteNil ^ 0xff
	TypeByteFloat32Inverse = TypeByteFloat32 ^ 0xff
	TypeByteFloat64Inverse = TypeByteFloat64 ^ 0xff
	TypeByteUint8Inverse   = TypeByteUint8 ^ 0xff
	TypeByteInt8Inverse    = TypeByteInt8 ^ 0xff
	TypeByteUint16Inverse  = TypeByteUint16 ^ 0xff
	TypeByteInt16Inverse   = TypeByteInt16 ^ 0xff
	TypeByteUint32Inverse  = TypeByteUint32 ^ 0xff
	TypeByteInt32Inverse   = TypeByteInt32 ^ 0xff
	TypeByteUint64Inverse  = TypeByteUint64 ^ 0xff
	TypeByteInt64Inverse   = TypeByteInt64 ^ 0xff
	TypeByteStringInverse  = TypeByteString ^ 0xff
	TypeByteTimeInverse    = TypeByteTime ^ 0xff
)

func TypeByte(v interface{}) (tb byte, err error) {
	if v == nil {
		tb = TypeByteNil
		return
	}
	switch v.(type) {
	case bool:
		tb = TypeByteBool
	case uint8:
		tb = TypeByteUint8
	case int8:
		tb = TypeByteInt8
	case uint16:
		tb = TypeByteUint16
	case int16:
		tb = TypeByteInt16
	case uint32:
		tb = TypeByteUint32
	case int32:
		tb = TypeByteInt32
	case uint64:
		tb = TypeByteUint64
	case int64:
		tb = TypeByteInt64
	case float32:
		tb = TypeByteFloat32
	case float64:
		tb = TypeByteFloat64
	case time.Time:
		tb = TypeByteTime
	case string:
		tb = TypeByteString
	default:
		err = errors.ErrUnsupportedType
	}
	return
}
