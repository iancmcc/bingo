package keypack_test

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/iancmcc/keypack/internal"
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
		EncodeValue(out, i, false)
		EncodeValue(out, i, true)
	}
}

func BenchmarkTestInt16Encoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 3, 3)
	var i int16
	b.StartTimer()
	for i = 0; i < 255; i++ {
		EncodeValue(out, i, false)
		EncodeValue(out, i, true)
	}
}

func BenchmarkTestInt32Encoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 5, 5)
	var i int32
	b.StartTimer()
	for i = 0; i < 1024; i++ {
		EncodeValue(out, i, false)
		EncodeValue(out, i, true)
	}
}

func BenchmarkTestInt64Encoder(b *testing.B) {
	out := make([]byte, 9, 9)
	var i int64
	for i = 0; i < 2048; i++ {
		EncodeValue(out, i, false)
		EncodeValue(out, i, true)
	}
}

func BenchmarkTestFloat32Encoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 9, 9)
	var i float32
	b.StartTimer()
	for i = 0; i < 2048; i++ {
		EncodeValue(out, i+0.5, false)
		EncodeValue(out, i+0.5, true)
	}
}

func BenchmarkTestFloat64Encoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 9, 9)
	var i float64
	b.StartTimer()
	for i = 0; i < 2048; i++ {
		EncodeValue(out, i+0.5, false)
		EncodeValue(out, i+0.5, true)
	}
}

func BenchmarkTestStringEncoder(b *testing.B) {
	b.StopTimer()
	s := randomString(80)
	out := make([]byte, 100, 100)
	var i int
	for i = 0; i < 1022; i++ {
		b.StartTimer()
		EncodeValue(out, s, false)
		EncodeValue(out, s, true)
		b.StopTimer()
	}
}

func BenchmarkNilEncoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 1, 1)
	var i int
	b.StartTimer()
	for i = 0; i < 127; i++ {
		EncodeValue(out, nil, false)
		EncodeValue(out, nil, true)
	}
}

func BenchmarkInt8Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 2, 2)
	EncodeValue(out, int8(71), true)
	var v int8
	b.StartTimer()
	DecodeValue(out, &v)
	b.StopTimer()
	if v != int8(71) {
		panic("int8 bad decode")
	}
	EncodeValue(out, int8(69), false)
	b.StartTimer()
	DecodeValue(out, &v)
	b.StopTimer()
	if v != int8(69) {
		panic("int8 bad decode")
	}
}

func BenchmarkInt16Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 4, 4)
	EncodeValue(out, int16(254), true)
	var v int16
	b.StartTimer()
	DecodeValue(out, &v)
	b.StopTimer()
	if v != int16(254) {
		panic("int16 bad decode")
	}
	EncodeValue(out, int16(269), false)
	b.StartTimer()
	DecodeValue(out, &v)
	b.StopTimer()
	if v != int16(269) {
		panic("int16 bad decode")
	}
}

func BenchmarkInt32Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 5, 5)
	EncodeValue(out, int32(254), true)
	var v int32
	b.StartTimer()
	DecodeValue(out, &v)
	b.StopTimer()
	if v != int32(254) {
		panic("int32 bad decode")
	}
	EncodeValue(out, int32(269), false)
	b.StartTimer()
	DecodeValue(out, &v)
	b.StopTimer()
	if v != int32(269) {
		panic("int32 bad decode")
	}
}

func BenchmarkInt64Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 9, 9)
	EncodeValue(out, int64(182832901), true)
	var v int64
	b.StartTimer()
	DecodeValue(out, &v)
	b.StopTimer()
	if v != int64(182832901) {
		panic("int64 bad decode")
	}
	EncodeValue(out, int64(-12398201928), false)
	b.StartTimer()
	DecodeValue(out, &v)
	b.StopTimer()
	if v != int64(-12398201928) {
		panic("int64 bad decode")
	}
}

func BenchmarkFloat32Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 5, 5)
	EncodeValue(out, float32(1828.32901), true)
	var v float32
	b.StartTimer()
	DecodeValue(out, &v)
	b.StopTimer()
	if v != float32(1828.32901) {
		panic("float32 bad decode")
	}
	EncodeValue(out, float32(-1239.8201928), false)
	b.StartTimer()
	DecodeValue(out, &v)
	b.StopTimer()
	if v != float32(-1239.8201928) {
		panic("float32 bad decode")
	}
}

func BenchmarkFloat64Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 9, 9)
	EncodeValue(out, float64(1828.64901), true)
	var v float64
	b.StartTimer()
	DecodeValue(out, &v)
	b.StopTimer()
	if v != float64(1828.64901) {
		panic("float64 bad decode")
	}
	EncodeValue(out, float64(-1239.8201928), false)
	b.StartTimer()
	DecodeValue(out, &v)
	b.StopTimer()
	if v != float64(-1239.8201928) {
		panic("float64 bad decode")
	}
}

func BenchmarkStringDecoder(b *testing.B) {
	s := "now is the time for all good men"
	b.StopTimer()
	out := make([]byte, 256, 256)
	EncodeValue(out, s, true)
	var v string
	b.StartTimer()
	i, _ := DecodeValue(out, &v)
	b.StopTimer()
	if v != s {
		panic(fmt.Sprintf("string bad decode: %s != %s", s, v))
	}
	if i != len(s)+2 {
		panic("bad string length")
	}
	EncodeValue(out, s, false)
	b.StartTimer()
	i, _ = DecodeValue(out, &v)
	b.StopTimer()
	if v != s {
		panic("string bad decode")
	}
	if i != len(s)+2 {
		panic("bad string length")
	}
}
