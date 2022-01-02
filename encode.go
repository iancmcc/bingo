package bingo

import (
	"fmt"
	"time"

	"github.com/goccy/go-reflect"
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

func BEncodeValue(b []byte, v interface{}, inverse bool) int {
	if v == nil {
		return EncodeNil(b, inverse)
	}
	typ := reflect.TypeOf(v)
	switch typ.Kind() {
	case reflect.Int8:
		return EncodeInt8(b, v.(int8), inverse)
	case reflect.Int16:
		return EncodeInt16(b, v.(int16), inverse)
	case reflect.Int:
		return EncodeInt32(b, int32(v.(int)), inverse)
	case reflect.Int32:
		return EncodeInt32(b, v.(int32), inverse)
	case reflect.Int64:
		return EncodeInt64(b, v.(int64), inverse)
	case reflect.Float32:
		return EncodeFloat32(b, v.(float32), inverse)
	case reflect.Float64:
		return EncodeFloat64(b, v.(float64), inverse)
	case reflect.String:
		return EncodeString(b, v.(string), inverse)
	default:
		if _, ok := v.(time.Time); ok {
			return EncodeTime(b, v.(time.Time), inverse)
		} else {
			panic("jfkdls")
		}
	}
}
