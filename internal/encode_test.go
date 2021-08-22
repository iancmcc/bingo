package internal_test

import (
	"bytes"
	"fmt"

	. "github.com/iancmcc/keypack/internal"
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
	}
	alltypes = []string{"int8", "int16", "int", "int32", "int64", "float32", "float64", "string"}
)

var _ = Describe("Codec", func() {
	for _, typeName := range alltypes {
		var v [][]interface{}
		if typeName == "string" {
			v = testVals["string"]
		} else {
			v = testVals["numeric"]
		}
		var entries = make([]TableEntry, 0, len(v))
		for _, args := range v {
			entries = append(entries, Entry(args[0], cast(typeName, args[1]), cast(typeName, args[2])))
		}
		Describe(fmt.Sprintf("for type %s", typeName), func() {
			DescribeTable("correctly orders values which are", func(x, y interface{}) {
				a, b := make([]byte, SizeOf(x), SizeOf(x)), make([]byte, SizeOf(y), SizeOf(y))
				EncodeValue(a, x, false)
				EncodeValue(b, y, false)
				Ω(bytes.Compare(a, b)).Should(BeNumerically("==", -1))
			},
				entries...)
		})
	}

	/*
		Describe("for type int8", func() {
			DescribeTable("correctly orders values which are",
				func(x int, y int) {
					a, b := make([]byte, 2, 2), make([]byte, 2, 2)
					EncodeInt8(a, int8(x), false)
					EncodeInt8(b, int8(y), false)
					Ω(bytes.Compare(a, b)).Should(BeNumerically("==", -1))
				},
				Entry("positive", 1, 2),
				Entry("negative", -2, -1),
				Entry("non-negative", 0, 1),
				Entry("non-positive", -1, 0),
				Entry("mixed", -1, 1),
			)
			DescribeTable("correctly inversely orders values which are",
				func(x int, y int) {
					a, b := make([]byte, 2, 2), make([]byte, 2, 2)
					EncodeInt8(a, int8(x), true)
					EncodeInt8(b, int8(y), true)
					Ω(bytes.Compare(a, b)).Should(BeNumerically("==", 1))
				},
				Entry("positive", 1, 2),
				Entry("negative", -2, -1),
				Entry("non-negative", 0, 1),
				Entry("non-positive", -1, 0),
				Entry("mixed", -1, 1),
			)
		})
	*/
})

/*
func BenchmarkTestInt8Encoder(b *testing.B) {
	out := make([]byte, 2, 2)
	var i int8
	for i = 0; i < 127; i++ {
		EncodeInt8(out, i, false)
	}
}
*/

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
	default:
		panic("Unsupported type")
	}
}
