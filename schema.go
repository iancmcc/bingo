package bingo

import "github.com/iancmcc/bingo/codecs"

type (
	// Schema captures whether fields of a key should be encoded in inverse or
	// natural order.
	Schema uint64

	schemaPacker struct {
		s Schema
		b []byte
		i int
	}
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
	var size int
	for _, v := range vals {
		n, err := codecs.EncodedSize(v)
		if err != nil {
			return nil, err
		}
		size += n
	}
	buf := make([]byte, size, size)
	s.packSlice(buf, vals)
	return buf, nil
}

// PackInto encodes the values passed into the provided byte slice, returning
// the number of bytes written.
func (s Schema) PackInto(b []byte, vals ...interface{}) (n int, err error) {
	return s.packSlice(b, vals)
}

// PackSlice encodes the values passed into the provided byte slice, returning
// the number of bytes written, for cases where use of the unpack operator would
// result in an extra allocation.
func (s Schema) PackSlice(b []byte, vals []interface{}) (n int, err error) {
	return s.packSlice(b, vals)
}

func (s Schema) packSlice(b []byte, vals []interface{}) (n int, err error) {
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

// NewPacker returns a Packer for the byte slice provided that uses this schema.
func (s Schema) NewPacker(b []byte) Packer {
	return schemaPacker{s, b, 0}
}

// Pack encodes the value provided into byte slice represented by the Packer.
func (s schemaPacker) Pack(v interface{}) Packer {
	desc := s.s&(1<<s.i) > 0
	n, _ := codecs.EncodeValue(s.b, v, desc)
	s.b = s.b[n:]
	s.i += 1
	return s
}

// Done releases the Packer's reference to the byte slice.
func (s schemaPacker) Done() {
	s.b = nil
}
