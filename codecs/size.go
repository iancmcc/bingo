package codecs

import (
	"bytes"
	"time"
)

// EncodedSize returns the number of bytes required to encode the value v.
func EncodedSize(v interface{}) (int, error) {
	if v == nil {
		return 1, nil
	}
	switch t := v.(type) {
	case int8:
		return sizeInt8, nil
	case int16:
		return sizeInt16, nil
	case int:
		if intsize == 4 {
			return sizeInt32, nil
		}
		return sizeInt64, nil
	case int32:
		return sizeInt32, nil
	case int64:
		return sizeInt64, nil
	case float32:
		return sizeFloat32, nil
	case float64:
		return sizeFloat64, nil
	case time.Time:
		return sizeTime, nil
	case string:
		return len(t) + 2, nil
	default:
		return 0, ErrUnknownType
	}
}

// SizeNext returns the number of bytes encompassing the next encoded value in
// the byte slice.
func SizeNext(b []byte) (int, error) {
	switch b[0] {
	case typeByteNil, typeByteNilInverse:
		return sizeNil, nil
	case typeByteInt8, typeByteInt8Inverse:
		return sizeInt8, nil
	case typeByteInt16, typeByteInt16Inverse:
		return sizeInt16, nil
	case typeByteInt32, typeByteInt32Inverse:
		return sizeInt32, nil
	case typeByteInt64, typeByteInt64Inverse:
		return sizeInt64, nil
	case typeByteFloat32, typeByteFloat32Inverse:
		return sizeFloat32, nil
	case typeByteFloat64, typeByteFloat64Inverse:
		return sizeFloat64, nil
	case typeByteTime, typeByteTimeInverse:
		return sizeTime, nil
	case typeByteString:
		return bytes.IndexByte(b, terminatorByte) + 1, nil
	case typeByteStringInverse:
		return bytes.IndexByte(b, terminatorByteInverse) + 1, nil
	default:
		return 0, ErrUnknownType
	}
}
