package codecs

import (
	"fmt"
	"time"
)

func EncodeValue(b []byte, v interface{}, inverse bool) int {
	if v == nil {
		return EncodeNil(b, inverse)
	}
	switch c := v.(type) {
	case int8:
		return EncodeInt8(b, c, inverse)
	case int16:
		return EncodeInt16(b, c, inverse)
	case int:
		return EncodeInt32(b, int32(c), inverse)
	case int32:
		return EncodeInt32(b, c, inverse)
	case int64:
		return EncodeInt64(b, c, inverse)
	case float32:
		return EncodeFloat32(b, c, inverse)
	case float64:
		return EncodeFloat64(b, c, inverse)
	case string:
		return EncodeString(b, c, inverse)
	case time.Time:
		return EncodeTime(b, c, inverse)
	default:
		panic(fmt.Sprintf("unknown type: %T", v))
	}
}
