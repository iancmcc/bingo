package codecs_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/iancmcc/bingo/codecs"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTime(t *testing.T) {

	Convey("time value", t, func() {

		var a = RandomTime()
		expectedSize, err := EncodedSize(a)
		So(err, ShouldBeNil)

		for _, inverse := range []bool{true, false} {
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

				Convey("and decoded into a nil pointer", func() {
					var v time.Time
					n, err := DecodeValue(b, &v)
					So(n, ShouldEqual, expectedSize)
					So(err, ShouldBeNil)
					So(v, ShouldEqual, a)
				})

				Convey("and maintain lexicographical order", func() {
					c := make([]byte, expectedSize, expectedSize)
					var addend = time.Second
					if inverse {
						addend = -addend
					}
					EncodeValue(c, a.Add(addend), inverse)
					So(bytes.Compare(b, c), ShouldBeLessThan, 0)
				})

			})
		}

		Convey("returns an error when byte array is insufficient", func() {
			b := make([]byte, expectedSize-1)
			_, err := EncodeValue(b, a, false)
			So(err, ShouldEqual, ErrByteSliceSize)
		})

	})

}

func RandomTime() time.Time {
	return time.Time(time.Unix(int64(rand.Int31()), int64(rand.Int31())))
}
