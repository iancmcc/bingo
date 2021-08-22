package codecs

import "bytes"

func EncodedSize(v interface{}) int {
	if v == nil {
		return 1
	}
	switch t := v.(type) {
	case int8:
		return sizeInt8
	case int16:
		return sizeInt16
	case int, int32:
		return sizeInt32
	case int64:
		return sizeInt64
	case float32:
		return sizeFloat32
	case float64:
		return sizeFloat64
	case string:
		return len(t) + 2
	default:
		panic("unknown type")
	}
}

func SizeNext(b []byte) int {
	switch b[0] {
	case typeByteNil, typeByteNilInverse:
		return sizeNil
	case typeByteInt8, typeByteInt8Inverse:
		return sizeInt8
	case typeByteInt16, typeByteInt16Inverse:
		return sizeInt16
	case typeByteInt32, typeByteInt32Inverse:
		return sizeInt32
	case typeByteInt64, typeByteInt64Inverse:
		return sizeInt64
	case typeByteFloat32, typeByteFloat32Inverse:
		return sizeFloat32
	case typeByteFloat64, typeByteFloat64Inverse:
		return sizeFloat64
	case typeByteString:
		return bytes.IndexByte(b, terminatorByte) + 1
	case typeByteStringInverse:
		return bytes.IndexByte(b, terminatorByteInverse) + 1
	default:
		panic("Unknown type")
	}
}
