package bingo_test

import (
	"bytes"
	"testing"

	. "github.com/iancmcc/bingo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Packing", func() {

	It("should pack values while preserving order", func() {
		a, err := Pack(1, "hi", int64(67))
		Ω(err).ShouldNot(HaveOccurred())
		b, err := Pack(1, "hi 1", int64(67))
		Ω(err).ShouldNot(HaveOccurred())
		Ω(bytes.Compare(a, b)).Should(Equal(-1))
	})

	It("should pack values into a given array while preserving order", func() {
		buf := make([]byte, 32)
		bufd := make([]byte, 32)
		PackTo(buf, 1, "hi", int64(67))
		PackTo(bufd, 1, "hi 1", int64(67))
		Ω(bytes.Compare(buf, bufd)).Should(Equal(-1))
	})

	It("should pack mixed-order values while preserving order", func() {
		s := WithDesc(false, true, false)
		a, err := s.Pack(1, "hi", int64(67))
		Ω(err).ShouldNot(HaveOccurred())
		b, err := s.Pack(1, "hi 1", int64(67))
		Ω(err).ShouldNot(HaveOccurred())
		Ω(bytes.Compare(a, b)).Should(Equal(1))
	})

	It("should pack mixed-order values into a given byte array while preserving order", func() {
		buf := make([]byte, 32)
		bufd := make([]byte, 32)
		s := WithDesc(false, true, false)
		s.PackTo(buf, 1, "hi", int64(67))
		s.PackTo(bufd, 1, "hi 1", int64(67))
		Ω(bytes.Compare(buf, bufd)).Should(Equal(1))
	})
})
var _ = Describe("Unpacking", func() {

	It("should unpack values", func() {
		a := "this is a test"
		b := int8(69)
		c := float32(1.61803398875)
		packed, err := Pack(a, b, c)
		Ω(err).ShouldNot(HaveOccurred())

		var (
			adest string
			bdest int8
			cdest float32
		)
		Unpack(packed, &adest, &bdest, &cdest)
		Ω(a).Should(Equal(adest))
		Ω(b).Should(Equal(bdest))
		Ω(c).Should(Equal(cdest))
	})

	It("should unpack only those asked for", func() {
		a := "this is a test"
		b := int8(69)
		c := float32(1.61803398875)
		packed, err := Pack(a, b, c)
		Ω(err).ShouldNot(HaveOccurred())

		var (
			adest string
			cdest float32
		)
		Unpack(packed, &adest, nil, &cdest)
		Ω(a).Should(Equal(adest))
		Ω(c).Should(Equal(cdest))
	})

	It("should unpack a value at a specific index", func() {
		a := "this is a test"
		b := int8(69)
		c := float32(1.61803398875)
		packed, err := Pack(a, b, c)
		Ω(err).ShouldNot(HaveOccurred())

		var (
			adest string
			cdest float32
		)

		UnpackIndex(packed, 2, &cdest)
		Ω(c).Should(Equal(cdest))

		UnpackIndex(packed, 0, &adest)
		Ω(a).Should(Equal(adest))
	})

	It("should error when a nonexistent index is requested", func() {
		a := "this is a test"
		b := int8(69)
		c := float32(1.61803398875)
		packed, err := Pack(a, b, c)
		Ω(err).ShouldNot(HaveOccurred())

		var cdest float32
		Ω(UnpackIndex(packed, 7, &cdest)).Should(HaveOccurred())
	})

	It("should error when unpacking a random string", func() {
		b := []byte("abcde")
		var adest string
		Ω(Unpack(b, &adest)).Should(HaveOccurred())
	})
})

func TestBingo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bingo Suite")
}
