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

func TestInt16(t *testing.T) {

	for _, negative := range []bool{false, true} {
		negdesc := "positive"
		if negative {
			negdesc = "negative"
		}
		Convey(fmt.Sprintf("a %s int16", negdesc), t, func() {

			var a int16 = int16(rand.Int63n(math.MaxInt16 - 1))
			if negative {
				a = -a
			}
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

					Convey("and decoded into an int16 pointer", func() {
						var v int16
						n, err := DecodeValue(b, &v)

						So(n, ShouldEqual, expectedSize)
						So(err, ShouldBeNil)
						So(v, ShouldEqual, a)
					})
					Convey("and maintain lexicographical order", func() {
						c := make([]byte, expectedSize, expectedSize)
						var addend int16 = 1
						if inverse {
							addend *= -1
						}
						EncodeValue(c, a+addend, inverse)
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

}
