package keypack

import (
	"github.com/iancmcc/keypack/internal/codecs"
)

var (
	Encode      = codecs.Encode
	EncodeValue = codecs.EncodeValue

	DecodeValue = codecs.DecodeValue
)
