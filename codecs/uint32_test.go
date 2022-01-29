package codecs_test

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"testing"

	. "github.com/iancmcc/bingo/codecs"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUint32(t *testing.T) {

	Convey("a uint32", t, func() {

		var a uint32 = uint32(rand.Int63n(math.MaxUint32 - 1))
		expectedSize, err := EncodedSize(a)
		So(err, ShouldBeNil)
		for _, inverse := range []bool{false, true} {
			invdesc := "natural"
			if inverse {
				invdesc = "inverse"
			}
			So(err, ShouldBeNil)
			Convey(fmt.Sprintf("can be encoded in %s order into a sufficient byte array", invdesc), func() {

				b := make([]byte, expectedSize, expectedSize)
				n, err := EncodeValue(b, a, inverse)
				So(n, ShouldEqual, expectedSize)
				So(err, ShouldBeNil)
				nn, err := SizeNext(b)
				So(nn, ShouldEqual, expectedSize)
				So(err, ShouldBeNil)

				Convey("and decoded into an uint32 pointer", func() {
					var v uint32
					n, err := DecodeValue(b, &v)

					So(n, ShouldEqual, expectedSize)
					So(err, ShouldBeNil)
					So(v, ShouldEqual, a)
				})

				Convey("and maintain lexicographical order", func() {
					c := make([]byte, expectedSize, expectedSize)
					var addend uint32 = 1
					if inverse {
						EncodeValue(c, a-addend, inverse)
					} else {
						EncodeValue(c, a+addend, inverse)
					}
					So(bytes.Compare(b, c), ShouldBeLessThan, 0)
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
