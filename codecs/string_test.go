package codecs_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"

	. "github.com/iancmcc/bingo/codecs"
	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {

	Convey("a string value", t, func() {

		var a = RandomString(rand.Intn(1024))
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

				Convey("and decoded into a string pointer", func() {
					var v string
					n, err := DecodeValue(b, &v)
					So(n, ShouldEqual, expectedSize)
					So(err, ShouldBeNil)
					So(a, ShouldEqual, a)
				})

				Convey("and maintain lexicographical order", func() {
					c := make([]byte, expectedSize+1, expectedSize+1)
					var addend string = "a"
					EncodeValue(c, a+addend, inverse)
					if inverse {
						So(bytes.Compare(c, b), ShouldBeLessThan, 0)
					} else {
						So(bytes.Compare(b, c), ShouldBeLessThan, 0)
					}
				})

			})
			Convey(fmt.Sprintf("returns an error when the %s-order string contains a null byte", invdesc), func() {
				b := make([]byte, expectedSize+1, expectedSize+1)
				c := []byte(a)
				c[rand.Intn(len(c))] = 0
				d := string(c)
				_, err := EncodeValue(b, d, inverse)
				So(err, ShouldEqual, ErrNullByte)
			})
		}

		Convey("returns an error when byte array is insufficient", func() {
			b := make([]byte, expectedSize-1)
			_, err := EncodeValue(b, a, false)
			So(err, ShouldEqual, ErrByteSliceSize)
		})
	})

}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
