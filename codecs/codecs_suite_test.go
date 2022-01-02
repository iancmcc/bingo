package codecs_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCodecs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Codecs Suite")
}
