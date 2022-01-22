package codecs

import (
	"time"
	"unsafe"
)

const intsize = unsafe.Sizeof(int(0))

// EncodeValue encodes value v to the beginning of byte slice b, optionally in
// inverse order, and returns the number of bytes written.
func EncodeValue(b []byte, v interface{}, inverse bool) (int, error) {
	if v == nil {
		return encodeNil(b, inverse)
	}
	switch c := v.(type) {
	case int8:
		return encodeInt8(b, c, inverse)
	case int16:
		return encodeInt16(b, c, inverse)
	case int:
		if intsize == 4 {
			return encodeInt32(b, int32(c), inverse)
		}
		return encodeInt64(b, int64(c), inverse)
	case int32:
		return encodeInt32(b, c, inverse)
	case int64:
		return encodeInt64(b, c, inverse)
	case float32:
		return encodeFloat32(b, c, inverse)
	case float64:
		return encodeFloat64(b, c, inverse)
	case string:
		return encodeString(b, c, inverse)
	case time.Time:
		return encodeTime(b, c, inverse)
	default:
		return 0, ErrUnknownType
	}
}
