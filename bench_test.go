package keypack_test

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/iancmcc/keypack"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomString(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func BenchmarkTestInt8Encoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 2, 2)
	var i int8
	b.StartTimer()
	for i = 0; i < 127; i++ {
		PackInto(out, i)
		WithDesc(true).PackInto(out, i)
	}
}

func BenchmarkTestInt16Encoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 3, 3)
	var i int16
	b.StartTimer()
	for i = 0; i < 255; i++ {
		PackInto(out, i)
		WithDesc(true).PackInto(out, i)
	}
}

func BenchmarkTestInt32Encoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 5, 5)
	var i int32
	b.StartTimer()
	for i = 0; i < 1024; i++ {
		PackInto(out, i)
		WithDesc(true).PackInto(out, i)
	}
}

func BenchmarkTestInt64Encoder(b *testing.B) {
	out := make([]byte, 9, 9)
	var i int64
	for i = 0; i < 2048; i++ {
		PackInto(out, i)
		WithDesc(true).PackInto(out, i)
	}
}

func BenchmarkTestFloat32Encoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 9, 9)
	var i float32
	b.StartTimer()
	for i = 0; i < 2048; i++ {
		PackInto(out, i)
		WithDesc(true).PackInto(out, i)
	}
}

func BenchmarkTestFloat64Encoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 9, 9)
	var i float64
	b.StartTimer()
	for i = 0; i < 2048; i++ {
		PackInto(out, i)
		WithDesc(true).PackInto(out, i)
	}
}

func BenchmarkTestStringEncoder(b *testing.B) {
	b.StopTimer()
	s := randomString(80)
	out := make([]byte, 100, 100)
	var i int
	for i = 0; i < 1022; i++ {
		b.StartTimer()
		PackInto(out, i)
		WithDesc(true).PackInto(out, s)
		b.StopTimer()
	}
}

func BenchmarkInt8Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 2, 2)
	WithDesc(true).PackInto(out, int8(71))
	var v int8
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != int8(71) {
		panic("int8 bad decode")
	}
	PackInto(out, int8(69))
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != int8(69) {
		panic("int8 bad decode")
	}
}

func BenchmarkInt16Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 4, 4)
	WithDesc(true).PackInto(out, int16(254))
	var v int16
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != int16(254) {
		panic("int16 bad decode")
	}
	PackInto(out, int16(269))
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != int16(269) {
		panic("int16 bad decode")
	}
}

func BenchmarkInt32Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 5, 5)
	WithDesc(true).PackInto(out, int32(254))
	var v int32
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != int32(254) {
		panic("int32 bad decode")
	}
	PackInto(out, int32(269))
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != int32(269) {
		panic("int32 bad decode")
	}
}

func BenchmarkInt64Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 9, 9)
	WithDesc(true).PackInto(out, int64(182832901))
	var v int64
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != int64(182832901) {
		panic("int64 bad decode")
	}
	PackInto(out, int64(-12398201928))
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != int64(-12398201928) {
		panic("int64 bad decode")
	}
}

func BenchmarkFloat32Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 5, 5)
	WithDesc(true).PackInto(out, float32(1828.32901))
	var v float32
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != float32(1828.32901) {
		panic("float32 bad decode")
	}
	PackInto(out, float32(-1239.8201928))
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != float32(-1239.8201928) {
		panic("float32 bad decode")
	}
}

func BenchmarkFloat64Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 9, 9)
	WithDesc(true).PackInto(out, float64(1828.64901))
	var v float64
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != float64(1828.64901) {
		panic("float64 bad decode")
	}
	PackInto(out, float64(-1239.8201928))
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != float64(-1239.8201928) {
		panic("float64 bad decode")
	}
}

func BenchmarkStringDecoder(b *testing.B) {
	s := "now is the time for all good men"
	b.StopTimer()
	out := make([]byte, 256, 256)
	WithDesc(true).PackInto(out, s)
	var v string
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != s {
		panic(fmt.Sprintf("string bad decode: %s != %s", s, v))
	}
	PackInto(out, s)
	b.StartTimer()
	Unpack(out, &v)
	b.StopTimer()
	if v != s {
		panic("string bad decode")
	}
}
