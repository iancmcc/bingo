package codecs

import "github.com/goccy/go-reflect"

// Decodes a value from b into the pointer v. Returns the number of bytes
// consumed
func DecodeValue(b []byte, v interface{}) (int, error) {
	rv := reflect.ValueNoEscapeOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		panic("Yo I need a pointer")
	}
	switch b[0] {
	case typeByteNil, typeByteNilInverse:
		rv.Set(reflect.ValueNoEscapeOf(nil))
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
		panic("Unknown type")
	}
}
