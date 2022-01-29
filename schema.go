package bingo

import (
	"io"

	"github.com/iancmcc/bingo/codecs"
)

type (
	// Schema captures whether fields of a key should be encoded in inverse or
	// natural order.
	Schema uint64
)

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

// Pack encodes the values passed, returning the resulting byte slice.
func (s Schema) Pack(vals ...interface{}) ([]byte, error) {
	return s.pack(vals)
}

// MustPack encodes the values passed, returning the resulting byte slice and
// panicking on error.
func (s Schema) MustPack(vals ...interface{}) []byte {
	result, err := s.pack(vals)
	if err != nil {
		panic(err)
	}
	return result
}

// PackTo encodes the values passed into the provided byte slice, returning
// the number of bytes written.
func (s Schema) PackTo(b []byte, vals ...interface{}) (n int, err error) {
	return s.packTo(b, vals)
}

// MustPackTo encodes the values passed into the provided byte slice, returning
// the number of bytes written. Panics on error.
func (s Schema) MustPackTo(b []byte, vals ...interface{}) int {
	result, err := s.packTo(b, vals)
	if err != nil {
		panic(err)
	}
	return result
}

// WritePackedTo encodes the values passed and writes the result to the
// io.Writer specified. Note: this requires 1 heap alloc for the intermediate
// byte array.
func (s Schema) WritePackedTo(w io.Writer, vals ...interface{}) (n int, err error) {
	return s.writePackedTo(w, vals)
}

func (s Schema) pack(vals []interface{}) ([]byte, error) {
	size, err := PackedSize(vals)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, size, size)
	_, err = s.packTo(buf, vals)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (s Schema) mustPack(vals []interface{}) []byte {
	result, err := s.pack(vals)
	if err != nil {
		panic(err)
	}
	return result
}

func (s Schema) packTo(b []byte, vals []interface{}) (n int, err error) {
	for i, v := range vals {
		desc := s&(1<<i) > 0
		m, err := codecs.EncodeValue(b[n:], v, desc)
		if err != nil {
			return n, err
		}
		n += m
	}
	return
}

func (s Schema) mustPackTo(b []byte, vals []interface{}) int {
	result, err := s.packTo(b, vals)
	if err != nil {
		panic(err)
	}
	return result
}

func (s Schema) writePackedTo(w io.Writer, vals []interface{}) (n int, err error) {
	b, err := s.pack(vals)
	if err != nil {
		return 0, err
	}
	return w.Write(b)
}
