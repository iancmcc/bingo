package codecs_test

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/goccy/go-reflect"

	. "github.com/iancmcc/bingo/codecs"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInt8(t *testing.T) {

	for _, negative := range []bool{false, true} {
		negdesc := "positive"
		if negative {
			negdesc = "negative"
		}
		Convey(fmt.Sprintf("a %s int8", negdesc), t, func() {

			var a int8 = int8(rand.Intn(math.MaxInt8 - 1))
			if negative {
				a = -a
			}
			for _, inverse := range []bool{false, true} {
				invdesc := "natural"
				if inverse {
					invdesc = "inverse"
				}
				Convey(fmt.Sprintf("can be encoded in %s order into a sufficient byte array", invdesc), func() {

					b := make([]byte, 2, 2)
					EncodeInt8(b, a, inverse)

					Convey("and decoded into an int8 pointer", func() {
						var v int8
						n, err := DecodeInt8(b, reflect.ValueOf(&v))

						So(n, ShouldEqual, 2)
						So(err, ShouldBeNil)
						So(v, ShouldEqual, a)
					})
					Convey("and maintain lexicographical order", func() {
						c := make([]byte, 2, 2)
						var addend int8 = 1
						if inverse {
							addend *= -1
						}
						EncodeInt8(c, a+addend, inverse)
						So(bytes.Compare(b, c), ShouldBeLessThan, 0)
					})
				})
			}
			Convey("throws an error when encoded into an insufficient array", func() {
				b := make([]byte, 1, 1)
				_, err := EncodeInt8(b, a, false)
				So(err, ShouldEqual, ErrByteArraySize)
			})
		})
	}

}
