package codecs

import (
	"github.com/goccy/go-reflect"
)

var ()

// DecodeValue decodes the first value in b into the location defined by pointer
// v.
func DecodeValue(b []byte, v interface{}) (int, error) {
	rv := reflect.ValueNoEscapeOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return 0, ErrNotAPointer
	}
	switch b[0] {
	case typeByteNil, typeByteNilInverse:
		return 1, nil
	case typeByteInt8, typeByteInt8Inverse:
		return DecodeInt8(b, rv)
	case typeByteInt16, typeByteInt16Inverse:
		return DecodeInt16(b, rv)
	case typeByteInt32, typeByteInt32Inverse:
		return DecodeInt32(b, rv)
	case typeByteInt64, typeByteInt64Inverse:
		return DecodeInt64(b, rv)
	case typeByteFloat32, typeByteFloat32Inverse:
		return DecodeFloat32(b, rv)
	case typeByteFloat64, typeByteFloat64Inverse:
		return DecodeFloat64(b, rv)
	case typeByteTime, typeByteTimeInverse:
		return DecodeTime(b, rv)
	case typeByteString, typeByteStringInverse:
		return DecodeString(b, rv)
	default:
		return 0, ErrUnknownType
	}
}
