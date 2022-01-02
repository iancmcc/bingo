package bingo

import (
	"errors"
	"fmt"

	"github.com/iancmcc/bingo/codecs"
)

const defaultSchema Schema = 0

// Pack returns a byte slice containing the key composed of the values provided.
func Pack(vals ...interface{}) []byte {
	return defaultSchema.Pack(vals...)
}

// PackInto packs vals into b.
func PackInto(b []byte, vals ...interface{}) (n int) {
	return defaultSchema.PackSlice(b, vals)
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
		if n >= len(b) {
			return errors.New(fmt.Sprintf("No data at index %d", idx))
		}
		n += codecs.SizeNext(b[n:])
	}
	return Unpack(b[n:], dest)
}
