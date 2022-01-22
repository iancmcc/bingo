package codecs_test

import (
	"math"
	"math/rand"
	"testing"

	. "github.com/iancmcc/bingo/codecs"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDecode(t *testing.T) {

	Convey("decoding a value", t, func() {

		var a int32 = int32(rand.Int63n(math.MaxInt32 - 1))
		b := make([]byte, 16)
		EncodeValue(b, a, false)

		Convey("should fail if a non-pointer is passed as a target", func() {
			var v int32
			_, err := DecodeValue(b, v)
			So(err, ShouldEqual, ErrInvalidTarget)
		})

		Convey("should fail if an unknown type is passed", func() {
			b[0] = 0xFF
			var v int32
			_, err := DecodeValue(b, &v)
			So(err, ShouldEqual, ErrUnknownType)

		})

	})

}
