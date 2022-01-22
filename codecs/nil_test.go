package codecs_test

import (
	"fmt"
	"testing"

	. "github.com/iancmcc/bingo/codecs"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNil(t *testing.T) {

	Convey("nil value", t, func() {

		var a interface{}
		expectedSize, err := EncodedSize(a)
		So(err, ShouldBeNil)

		for _, inverse := range []bool{true, false} {
			b := make([]byte, expectedSize, expectedSize)

			invdesc := "natural"
			if inverse {
				invdesc = "inverse"
			}
			So(err, ShouldBeNil)
			Convey(fmt.Sprintf("can be encoded in %s order into a sufficient byte array", invdesc), func() {

				n, err := EncodeValue(b, a, inverse)
				So(n, ShouldEqual, expectedSize)
				So(err, ShouldBeNil)
				nn, err := SizeNext(b)
				So(nn, ShouldEqual, expectedSize)
				So(err, ShouldBeNil)

				Convey("and decoded into a nil pointer", func() {
					var v interface{}
					n, err := DecodeValue(b, &v)
					So(n, ShouldEqual, expectedSize)
					So(err, ShouldBeNil)
					So(a, ShouldBeNil)
				})

			})
		}

		Convey("returns an error when byte array is insufficient", func() {
			var v interface{}
			b := make([]byte, expectedSize-1)
			_, err := EncodeValue(b, v, false)
			So(err, ShouldEqual, ErrByteSliceSize)
		})
	})

}
