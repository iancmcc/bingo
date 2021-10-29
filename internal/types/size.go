package types

import (
	"bytes"
	"time"
	"unsafe"

	"github.com/iancmcc/bingo/internal/errors"
)

const (
	SizeBool    = 1
	SizeNil     = 1
	SizeUint8   = int(unsafe.Sizeof(uint8(0))) + 1
	SizeUint16  = int(unsafe.Sizeof(uint16(0))) + 1
	SizeUint32  = int(unsafe.Sizeof(uint32(0))) + 1
	SizeUint64  = int(unsafe.Sizeof(uint64(0))) + 1
	SizeInt8    = int(unsafe.Sizeof(int8(0))) + 1
	SizeInt16   = int(unsafe.Sizeof(int16(0))) + 1
	SizeInt32   = int(unsafe.Sizeof(int32(0))) + 1
	SizeInt64   = int(unsafe.Sizeof(int64(0))) + 1
	SizeFloat32 = int(unsafe.Sizeof(float32(0))) + 1
	SizeFloat64 = int(unsafe.Sizeof(float64(0))) + 1
	SizeTime    = 16
)

func Size(v interface{}) (n int, err error) {
	if v == nil {
		return SizeNil, nil
	}
	switch t := v.(type) {
	case bool:
		n = SizeBool
	case uint8:
		n = SizeUint8
	case int8:
		n = SizeInt8
	case uint16:
		n = SizeUint16
	case int16:
		n = SizeInt16
	case uint32:
		n = SizeUint32
	case int32:
		n = SizeInt32
	case uint64:
		n = SizeUint64
	case int64:
		n = SizeInt64
	case float32:
		n = SizeFloat32
	case float64:
		n = SizeFloat64
	case time.Time:
		n = SizeTime
	case string:
		n = len(t) + 2
	default:
		err = errors.ErrUnsupportedType
	}
	return
}

func nextSize(b []byte) (n int, err error) {
	switch b[0] {
	case TypeByteNil, TypeByteNilInverse:
		n = SizeNil
	case TypeByteBool, TypeByteBoolInverse:
		n = SizeBool
	case TypeByteUint8, TypeByteUint8Inverse:
		n = SizeUint8
	case TypeByteUint16, TypeByteUint16Inverse:
		n = SizeUint16
	case TypeByteUint32, TypeByteUint32Inverse:
		n = SizeUint32
	case TypeByteUint64, TypeByteUint64Inverse:
		n = SizeUint64
	case TypeByteInt8, TypeByteInt8Inverse:
		n = SizeInt8
	case TypeByteInt16, TypeByteInt16Inverse:
		n = SizeInt16
	case TypeByteInt32, TypeByteInt32Inverse:
		n = SizeInt32
	case TypeByteInt64, TypeByteInt64Inverse:
		n = SizeInt64
	case TypeByteFloat32, TypeByteFloat32Inverse:
		n = SizeFloat32
	case TypeByteFloat64, TypeByteFloat64Inverse:
		n = SizeFloat64
	case TypeByteTime, TypeByteTimeInverse:
		n = SizeTime
	case TypeByteString:
		n = bytes.IndexByte(b, TerminatorByte) + 1
	case TypeByteStringInverse:
		n = bytes.IndexByte(b, TerminatorByteInverse) + 1
	default:
		err = errors.ErrUnknownTypeByte
	}
	return
}
