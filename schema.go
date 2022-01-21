package bingo

import "github.com/iancmcc/bingo/codecs"

type Schema uint64

type schemaPacker struct {
	s Schema
	b []byte
	i int
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

func (s Schema) PackInto(b []byte, vals ...interface{}) (n int, err error) {
	return s.packSlice(b, vals)
}

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

func (s Schema) Packer(b []byte) schemaPacker {
	return schemaPacker{s, b, 0}
}

func (s schemaPacker) Pack(v interface{}) schemaPacker {
	desc := s.s&(1<<s.i) > 0
	n, _ := codecs.EncodeValue(s.b, v, desc)
	s.b = s.b[n:]
	s.i += 1
	return s
}

func (s schemaPacker) Done() {
	s.b = nil
}
