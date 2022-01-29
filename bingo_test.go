package bingo_test

import (
	"bytes"
	"errors"
	"testing"

	. "github.com/iancmcc/bingo"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBingo(t *testing.T) {

	Convey("Bingo", t, func() {

		Convey("should pack values while preserving order", func() {
			a, err := Pack(1, "hi", int64(67))
			So(err, ShouldBeNil)
			b, err := Pack(1, "hi 1", int64(67))
			So(err, ShouldBeNil)
			So(bytes.Compare(a, b), ShouldEqual, -1)
		})

		Convey("should pack values panicking on invalid type", func() {
			So(func() { MustPack(1, "hi", int64(67)) }, ShouldNotPanic)
			So(func() { MustPack(1, struct{}{}, int64(67)) }, ShouldPanic)

			So(func() { WithDesc().MustPack(1, "hi", int64(67)) }, ShouldNotPanic)
			So(func() { WithDesc().MustPack(1, struct{}{}, int64(67)) }, ShouldPanic)
		})

		Convey("should pack values into a given array while preserving order", func() {
			buf := make([]byte, 32)
			bufd := make([]byte, 32)
			PackTo(buf, 1, "hi", int64(67))
			PackTo(bufd, 1, "hi 1", int64(67))
			So(bytes.Compare(buf, bufd), ShouldEqual, -1)
		})

		Convey("should pack values to a given array panicking on invalid type", func() {
			buf := make([]byte, 32)
			So(func() { MustPackTo(buf, 1, "hi", int64(67)) }, ShouldNotPanic)
			So(func() { MustPackTo(buf, 1, struct{}{}, int64(67)) }, ShouldPanic)

			So(func() { WithDesc().MustPackTo(buf, 1, "hi", int64(67)) }, ShouldNotPanic)
			So(func() { WithDesc().MustPackTo(buf, 1, struct{}{}, int64(67)) }, ShouldPanic)
		})

		Convey("should pack mixed-order values while preserving order", func() {
			s := WithDesc(false, true, false)
			a, err := s.Pack(1, "hi", int64(67))
			So(err, ShouldBeNil)
			b, err := s.Pack(1, "hi 1", int64(67))
			So(err, ShouldBeNil)
			So(bytes.Compare(a, b), ShouldEqual, 1)
		})

		Convey("should pack mixed-order values into a given byte array while preserving order", func() {
			buf := make([]byte, 32)
			bufd := make([]byte, 32)
			s := WithDesc(false, true, false)
			s.PackTo(buf, 1, "hi", int64(67))
			s.PackTo(bufd, 1, "hi 1", int64(67))
			So(bytes.Compare(buf, bufd), ShouldEqual, 1)
		})

		Convey("should unpack values", func() {
			a := "this is a test"
			b := int8(69)
			c := float32(1.61803398875)
			packed, err := Pack(a, b, c)
			So(err, ShouldBeNil)

			var (
				adest string
				bdest int8
				cdest float32
			)
			Unpack(packed, &adest, &bdest, &cdest)
			So(a, ShouldEqual, adest)
			So(b, ShouldEqual, bdest)
			So(c, ShouldEqual, cdest)
		})

		Convey("should unpack only those asked for", func() {
			a := "this is a test"
			b := int8(69)
			c := float32(1.61803398875)
			packed, err := Pack(a, b, c)
			So(err, ShouldBeNil)

			var (
				adest string
				cdest float32
			)
			Unpack(packed, &adest, nil, &cdest)
			So(a, ShouldEqual, adest)
			So(c, ShouldEqual, cdest)
		})

		Convey("should unpack a value at a specific index", func() {
			a := "this is a test"
			b := int8(69)
			c := float32(1.61803398875)
			packed, err := Pack(a, b, c)
			So(err, ShouldBeNil)

			var (
				adest string
				cdest float32
			)

			UnpackIndex(packed, 2, &cdest)
			So(c, ShouldEqual, cdest)

			UnpackIndex(packed, 0, &adest)
			So(a, ShouldEqual, adest)
		})

		Convey("should error when a nonexistent index is requested", func() {
			a := "this is a test"
			b := int8(69)
			c := float32(1.61803398875)
			packed, err := Pack(a, b, c)
			So(err, ShouldBeNil)

			var cdest float32
			So(UnpackIndex(packed, 7, &cdest), ShouldResemble, errors.New("No data at index 7"))
		})

		Convey("should error when unpacking a random string", func() {
			b := []byte("abcde")
			var adest string
			So(Unpack(b, &adest), ShouldResemble, errors.New("unknown type"))
		})

	})

}
