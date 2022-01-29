package codecs_test

import (
	"fmt"
	"testing"

	. "github.com/iancmcc/bingo/codecs"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBool(t *testing.T) {

	for _, a := range []bool{false, true} {
		desc := "false"
		if a {
			desc = "true"
		}
		Convey(fmt.Sprintf("a %s bool", desc), t, func() {
			expectedSize, err := EncodedSize(a)
			So(err, ShouldBeNil)
			for _, inverse := range []bool{false, true} {
				invdesc := "natural"
				if inverse {
					invdesc = "inverse"
				}
				Convey(fmt.Sprintf("can be encoded in %s order into a sufficient byte array", invdesc), func() {
					b := make([]byte, expectedSize, expectedSize)
					n, err := EncodeValue(b, a, inverse)
					So(n, ShouldEqual, expectedSize)
					So(err, ShouldBeNil)
					nn, err := SizeNext(b)
					So(nn, ShouldEqual, expectedSize)
					So(err, ShouldBeNil)

					Convey("and decoded into a bool pointer", func() {
						var v bool
						n, err := DecodeValue(b, &v)
						So(n, ShouldEqual, expectedSize)
						So(err, ShouldBeNil)
						So(v, ShouldEqual, a)
					})
				})
			}
			Convey("throws an error when encoded into an insufficient array", func() {
				b := make([]byte, expectedSize-1)
				_, err := EncodeValue(b, a, false)
				So(err, ShouldEqual, ErrByteSliceSize)
			})
		})
	}
}
