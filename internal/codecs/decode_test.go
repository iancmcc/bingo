package codecs_test

import (
	. "github.com/iancmcc/keypack/internal/codecs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Decode", func() {

	It("should decode int8", func() {
		b := make([]byte, 10, 10)
		EncodeValue(b, int8(71), false)

		var v int8
		n, err := DecodeValue(b, &v)
		Ω(n).Should(BeNumerically("==", 2))
		Ω(err).ShouldNot(HaveOccurred())
		Ω(v).Should(Equal(int8(71)))
	})

})
