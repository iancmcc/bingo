package bingo

import (
	"errors"
	"fmt"

	"github.com/iancmcc/bingo/codecs"
)

const defaultSchema Schema = 0

// Pack encodes the values passed, returning the resulting byte slice.
func Pack(vals ...interface{}) ([]byte, error) {
	return defaultSchema.Pack(vals...)
}

// PackInto encodes the values passed into the provided byte slice, returning
// the number of bytes written.
func PackInto(b []byte, vals ...interface{}) (n int, err error) {
	return defaultSchema.PackSlice(b, vals)
}

// PackSlice encodes the values passed into the provided byte slice, returning
// the number of bytes written, for cases where use of the unpack operator would
// result in an extra allocation.
func PackSlice(b []byte, vals []interface{}) (n int, err error) {
	return defaultSchema.PackSlice(b, vals)
}

// Packer returns an object that can be used to pack values into b.
func NewPacker(b []byte) Packer {
	return defaultSchema.NewPacker(b)
}

// Unpack unpacks b into the targets provided.
func Unpack(b []byte, dests ...interface{}) error {
	for _, dest := range dests {
		var (
			n   int
			err error
		)
		if dest == nil {
			// Skip this one
			n, err = codecs.SizeNext(b)
		} else {
			n, err = codecs.DecodeValue(b, dest)
		}
		if err != nil {
			return err
		}
		b = b[n:]
	}
	return nil
}

// UnpackIndex unpacks the idx'th value from b into the target provided.
func UnpackIndex(b []byte, idx int, dest interface{}) error {
	var n int
	for i := 0; i < idx; i++ {
		if n >= len(b) {
			return errors.New(fmt.Sprintf("No data at index %d", idx))
		}
		nn, err := codecs.SizeNext(b[n:])
		if err != nil {
			return err
		}
		n += nn
	}
	return Unpack(b[n:], dest)
}
