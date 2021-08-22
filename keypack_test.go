package keypack_test

import (
	"bytes"
	"testing"

	. "github.com/iancmcc/keypack"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Packing", func() {
	It("should pack values while preserving order", func() {
		a := Pack(1, "hi", int64(67))
		b := Pack(1, "hi 1", int64(67))
		Ω(bytes.Compare(a, b)).Should(Equal(-1))
	})
	It("should pack mixed-order values while preserving order", func() {
		s := WithDesc(false, true, false)
		a := s.Pack(1, "hi", int64(67))
		b := s.Pack(1, "hi 1", int64(67))
		Ω(bytes.Compare(a, b)).Should(Equal(1))
	})

	It("should unpack values", func() {
		a := "this is a test"
		b := int8(69)
		c := float32(1.61803398875)
		packed := Pack(a, b, c)

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
		packed := Pack(a, b, c)

		var (
			adest string
			cdest float32
		)
		Unpack(packed, &adest, nil, &cdest)
		Ω(a).Should(Equal(adest))
		Ω(c).Should(Equal(cdest))
	})
})

func TestKeypack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keypack Suite")
}
