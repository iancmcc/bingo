package internal

func SizeOf(v interface{}) int {
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
	case float32:
		return sizeFloat32
	case int64:
		return sizeInt64
	case float64:
		return sizeFloat64
	case string:
		return len(t) + 2
	default:
		panic("unknown type")
	}
}

func EncodeValue(b []byte, v interface{}, inverse bool) {
	if v == nil {
		EncodeNil(b, inverse)
		return
	}
	switch c := v.(type) {
	case int8:
		EncodeInt8(b, c, inverse)
	case int16:
		EncodeInt16(b, c, inverse)
	case int:
		EncodeInt32(b, int32(c), inverse)
	case int32:
		EncodeInt32(b, c, inverse)
	case int64:
		EncodeInt64(b, c, inverse)
	case float32:
		EncodeFloat32(b, c, inverse)
	case float64:
		EncodeFloat64(b, c, inverse)
	case string:
		EncodeString(b, c, inverse)
	default:
		panic("unknown type")
	}
}

func Encode(vals ...interface{}) []byte {
	var i, size int
	for _, v := range vals {
		size += SizeOf(v)
	}
	buf := make([]byte, size, size)
	for _, v := range vals {
		next := i + SizeOf(v)
		Encode(buf[i:next], v)
		i = next
	}
	return buf
}
