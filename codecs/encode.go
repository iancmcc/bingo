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
	case bool:
		return encodeBool(b, c, inverse)
	case uint8:
		return encodeUint8(b, c, inverse)
	case uint16:
		return encodeUint16(b, c, inverse)
	case uint:
		if intsize == 4 {
			return encodeUint32(b, uint32(c), inverse)
		}
		return encodeUint64(b, uint64(c), inverse)
	case uint32:
		return encodeUint32(b, c, inverse)
	case uint64:
		return encodeUint64(b, c, inverse)
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
