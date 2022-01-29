package bingo

import (
	"fmt"
	"io"

	"github.com/iancmcc/bingo/codecs"
)

const defaultSchema Schema = 0

// Pack encodes the values passed, returning the resulting byte slice.
// This requires 1 heap alloc for the return value.
func Pack(vals ...interface{}) ([]byte, error) {
	return defaultSchema.pack(vals)
}

// MustPack encodes the values passed, returning the resulting byte slice.
// Panics on error.
func MustPack(vals ...interface{}) []byte {
	return defaultSchema.mustPack(vals)
}

// PackTo encodes the values passed into the provided byte slice, returning the
// number of bytes written.
func PackTo(b []byte, vals ...interface{}) (n int, err error) {
	return defaultSchema.packTo(b, vals)
}

// MustPackTo encodes the values passed into the provided byte slice, returning the
// number of bytes written. Panics on error.
func MustPackTo(b []byte, vals ...interface{}) int {
	return defaultSchema.mustPackTo(b, vals)
}

// WritePackedTo encodes the values passed and writes the result to the
// io.Writer specified. Note: this requires 1 heap alloc for the intermediate
// byte array.
func WritePackedTo(w io.Writer, vals ...interface{}) (n int, err error) {
	return defaultSchema.writePackedTo(w, vals)
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
			return fmt.Errorf("No data at index %d", idx)
		}
		nn, err := codecs.SizeNext(b[n:])
		if err != nil {
			return err
		}
		n += nn
	}
	return Unpack(b[n:], dest)
}

// PackedSize returns the number of bytes required to pack the values passed
func PackedSize(vals []interface{}) (int, error) {
	var size int
	for _, v := range vals {
		n, err := codecs.EncodedSize(v)
		if err != nil {
			return 0, err
		}
		size += n
	}
	return size, nil
}
