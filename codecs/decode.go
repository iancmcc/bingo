package codecs

import (
	"github.com/goccy/go-reflect"
)

var ()

// DecodeValue decodes the first value in slice b into the location defined by
// pointer v.
func DecodeValue(b []byte, v interface{}) (int, error) {
	rv := reflect.ValueNoEscapeOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return 0, ErrInvalidTarget
	}
	switch b[0] {
	case typeByteNil, typeByteNilInverse:
		return 1, nil
	case typeByteInt8, typeByteInt8Inverse:
		return decodeInt8(b, rv)
	case typeByteInt16, typeByteInt16Inverse:
		return decodeInt16(b, rv)
	case typeByteInt32, typeByteInt32Inverse:
		return decodeInt32(b, rv)
	case typeByteInt64, typeByteInt64Inverse:
		return decodeInt64(b, rv)
	case typeByteFloat32, typeByteFloat32Inverse:
		return decodeFloat32(b, rv)
	case typeByteFloat64, typeByteFloat64Inverse:
		return decodeFloat64(b, rv)
	case typeByteTime, typeByteTimeInverse:
		return decodeTime(b, rv)
	case typeByteString, typeByteStringInverse:
		return decodeString(b, rv)
	default:
		return 0, ErrUnknownType
	}
}
