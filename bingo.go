package bingo

import (
	"github.com/iancmcc/bingo/internal/codecs"
)

const defaultSchema Schema = 0

// Pack returns a byte slice containing the key composed of the values provided.
func Pack(vals ...interface{}) []byte {
	return defaultSchema.Pack(vals...)
}

// PackInto packs vals into b.
func PackInto(b []byte, vals ...interface{}) (n int) {
	return defaultSchema.packSlice(b, vals)
}

func Packer(b []byte) schemaPacker {
	return defaultSchema.Packer(b)
}

func Unpack(b []byte, dests ...interface{}) error {
	for _, dest := range dests {
		var (
			n   int
			err error
		)
		if dest == nil {
			// Skip this one
			n = codecs.SizeNext(b)
		} else {
			n, err = codecs.DecodeValue(b, dest)
			if err != nil {
				return err
			}
		}
		b = b[n:]
	}
	return nil
}

func UnpackIndex(b []byte, idx int, dest interface{}) error {
	var n int
	for i := 0; i < idx; i++ {
		n += codecs.SizeNext(b[n:])
	}
	return Unpack(b[n:], dest)
}

// WithDesc returns a Schema that will produce packed keys with the indicated
// values encoded to sort in descending order.
func WithDesc(cols ...bool) Schema {
	var s Schema
	for i, t := range cols {
		if t {
			s |= (1 << i)
		}
	}
	return s
}
