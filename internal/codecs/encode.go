package codecs

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
		size += EncodedSize(v)
	}
	buf := make([]byte, size, size)
	for _, v := range vals {
		next := i + EncodedSize(v)
		Encode(buf[i:next], v)
		i = next
	}
	return buf
}
