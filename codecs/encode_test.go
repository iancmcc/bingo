package codecs_test

import (
	"testing"

	. "github.com/iancmcc/bingo/codecs"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEncode(t *testing.T) {
	Convey("encoding a value", t, func() {
		Convey("should return an error if a type is unknown", func() {
			m := map[string]int{}

			_, err := EncodeValue([]byte{}, m, false)
			So(err, ShouldEqual, ErrUnknownType)
		})
	})
}
