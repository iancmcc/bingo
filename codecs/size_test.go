package codecs_test

import (
	"testing"

	. "github.com/iancmcc/bingo/codecs"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSize(t *testing.T) {
	Convey("getting the encoded size of a value", t, func() {
		Convey("should return an error if a type is unknown", func() {
			m := map[string]int{}
			_, err := EncodedSize(m)
			So(err, ShouldEqual, ErrUnknownType)
		})
	})

	Convey("returning the size of the next field in a byte array", t, func() {
		Convey("should return an error if a type is unknown", func() {
			b := []byte{0xff, 0, 0, 0, 0, 0}
			_, err := SizeNext(b)
			So(err, ShouldEqual, ErrUnknownType)
		})
	})

}
