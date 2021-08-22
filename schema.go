package keypack

import "github.com/iancmcc/keypack/internal/codecs"

type Schema uint64

func (s Schema) Pack(vals ...interface{}) []byte {
	var size int
	for _, v := range vals {
		size += codecs.EncodedSize(v)
	}
	buf := make([]byte, size, size)
	s.PackInto(buf, vals...)
	return buf
}

func (s Schema) PackInto(b []byte, vals ...interface{}) (n int) {
	for i, v := range vals {
		desc := s&(1<<i) > 0
		n += codecs.EncodeValue(b[n:], v, desc)
	}
	return

}
