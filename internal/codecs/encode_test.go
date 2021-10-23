package codecs_test

import (
	"bytes"
	"fmt"
	"time"

	. "github.com/iancmcc/bingo/internal/codecs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var (
	testVals = map[string][][]interface{}{
		"numeric": {
			{"positive", 1, 2},
			{"negative", -2, -1},
			{"non-negative", 0, 1},
			{"non-positive", -1, 0},
			{"mixed", -1, 1},
		},
		"string": {
			{"the same length", "aa", "ab"},
			{"different lengths", "aaa", "abcde"},
			{"empty", "", "a"},
		},
		"time": {
			{"the same time zone", time.Now().Add(-time.Hour), time.Now()},
			{"different time zones", time.Now().Add(-time.Hour).UTC(), time.Now()},
		},
	}
	alltypes = []string{"int8", "int16", "int", "int32", "int64", "float32", "float64", "string", "time"}
)

var _ = Describe("Codec", func() {
	for _, typeName := range alltypes {
		var v [][]interface{}
		if typeName == "string" {
			v = testVals["string"]
		} else if typeName == "time" {
			v = testVals["time"]
		} else {
			v = testVals["numeric"]
		}
		var entries = make([]TableEntry, 0, len(v))
		for _, args := range v {
			entries = append(entries, Entry(args[0], cast(typeName, args[1]), cast(typeName, args[2])))
		}
		Describe(fmt.Sprintf("for type %s", typeName), func() {
			DescribeTable("correctly orders values which are", func(x, y interface{}) {
				a, b := make([]byte, EncodedSize(x), EncodedSize(x)), make([]byte, EncodedSize(y), EncodedSize(y))
				EncodeValue(a, x, false)
				EncodeValue(b, y, false)
				Ω(bytes.Compare(a, b)).Should(BeNumerically("==", -1))
			},
				entries...)
		})
	}

})

func cast(s string, v interface{}) interface{} {
	switch s {
	case "int8":
		return int8(v.(int))
	case "int16":
		return int16(v.(int))
	case "int":
		return v.(int)
	case "int32":
		return int32(v.(int))
	case "int64":
		return int64(v.(int))
	case "float32":
		return float32(v.(int))
	case "float64":
		return float64(v.(int))
	case "string":
		return v.(string)
	case "time":
		return v.(time.Time)
	default:
		panic("Unsupported type")
	}
}
