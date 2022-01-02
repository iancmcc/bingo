package bingo

type Schema uint64

type schemaPacker struct {
	s Schema
	b []byte
	i int
}

func (s Schema) Pack(vals ...interface{}) []byte {
	var size int
	for _, v := range vals {
		size += EncodedSize(v)
	}
	buf := make([]byte, size, size)
	s.packSlice(buf, vals)
	return buf
}

func (s Schema) PackInto(b []byte, vals ...interface{}) (n int) {
	return s.packSlice(b, vals)
}

func (s Schema) PackSlice(b []byte, vals []interface{}) (n int) {
	return s.packSlice(b, vals)
}

func (s Schema) packSlice(b []byte, vals []interface{}) (n int) {
	for i, v := range vals {
		desc := s&(1<<i) > 0
		n += EncodeValue(b[n:], v, desc)
	}
	return
}

func (s Schema) Packer(b []byte) schemaPacker {
	return schemaPacker{s, b, 0}
}

func (s schemaPacker) Pack(v interface{}) schemaPacker {
	desc := s.s&(1<<s.i) > 0
	n := EncodeValue(s.b, v, desc)
	s.b = s.b[n:]
	s.i += 1
	return s
}

func (s schemaPacker) Done() {
	s.b = nil
}
