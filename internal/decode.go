package internal

import "reflect"

// Decodes a value from b into the pointer v. Returns the number of bytes
// consumed
func DecodeValue(b []byte, v interface{}) (int, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		panic("Yo I need a pointer")
	}
	switch b[0] {
	case typeByteNil, typeByteNilInverse:
		rv.Set(reflect.ValueOf(nil))
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
	case typeByteString, typeByteStringInverse:
		return DecodeString(b, rv)
	default:
		panic("Unknown type")
	}
}

func ValueSize(b []byte) (int, error) {
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
	case typeByteString, typeByteStringInverse:
		return 5, nil
	default:
		panic("Unknown type")
	}
}
