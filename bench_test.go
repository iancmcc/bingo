package bingo_test

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/iancmcc/bingo"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

//go:noinline
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
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Packer(out).Pack(int8(55))
		WithDesc(true).Packer(out).Pack(int8(55))
	}
}

func BenchmarkTestInt16Encoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 3, 3)
	var i int16
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		PackInto(out, int16(123))
		WithDesc(true).PackInto(out, i)
	}
}

func BenchmarkTestInt32Encoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 5, 5)
	var i int32
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		PackInto(out, int32(83928329))
		WithDesc(true).PackInto(out, i)
	}
}

func BenchmarkTestInt64Encoder(b *testing.B) {
	out := make([]byte, 9, 9)
	var i int64
	for j := 0; j < b.N; j++ {
		PackInto(out, int64(898293012839))
		WithDesc(true).PackInto(out, i)
	}
}

func BenchmarkTestFloat32Encoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 9, 9)
	var i float32
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		PackInto(out, float32(123.5647372))
		WithDesc(true).PackInto(out, i)
	}
}

func BenchmarkTestFloat64Encoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 9, 9)
	var i float64
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		PackInto(out, float64(3.14159265358979323))
		WithDesc(true).PackInto(out, i)
	}
}

func BenchmarkTestStringEncoder(b *testing.B) {
	b.StopTimer()
	s := randomString(80)
	out := make([]byte, 85, 85)
	args := []interface{}{s}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		WithDesc(false).PackSlice(out, args)
		WithDesc(true).PackSlice(out, args)
	}
}

func BenchmarkInt8Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 2, 2)
	outdesc := make([]byte, 2, 2)
	var v int8
	WithDesc(false).PackInto(out, int8(71))
	WithDesc(true).PackInto(outdesc, int8(71))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Unpack(out, &v)
		Unpack(outdesc, &v)
	}
}

func BenchmarkInt16Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 4, 4)
	outdesc := make([]byte, 4, 4)
	var v int16
	WithDesc(false).PackInto(out, int16(279))
	WithDesc(true).PackInto(outdesc, int16(279))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Unpack(out, &v)
		Unpack(outdesc, &v)
	}
}

func BenchmarkIntDecoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 5, 5)
	outdesc := make([]byte, 5, 5)
	var v int
	WithDesc(true).PackInto(out, int(254))
	WithDesc(false).PackInto(outdesc, int(7485))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Unpack(out, &v)
		Unpack(outdesc, &v)
	}
}

func BenchmarkInt32Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 5, 5)
	outdesc := make([]byte, 5, 5)
	var v int32
	WithDesc(true).PackInto(out, int32(254))
	WithDesc(false).PackInto(outdesc, int32(7485))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Unpack(out, &v)
		Unpack(outdesc, &v)
	}
}

func BenchmarkInt64Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 9, 9)
	outdesc := make([]byte, 9, 9)
	var v int64
	WithDesc(false).PackInto(out, int64(182832901))
	WithDesc(true).PackInto(outdesc, int64(182832901))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Unpack(out, &v)
		Unpack(outdesc, &v)
	}
}

func BenchmarkFloat32Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 5, 5)
	outdesc := make([]byte, 5, 5)
	var v float32
	WithDesc(false).PackInto(out, float32(1828.32901))
	WithDesc(true).PackInto(outdesc, float32(1828.32901))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Unpack(out, &v)
		Unpack(outdesc, &v)
	}
}

func BenchmarkFloat64Decoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 9, 9)
	outdesc := make([]byte, 9, 9)
	var v float64
	WithDesc(false).PackInto(out, float64(1828.64901))
	WithDesc(true).PackInto(outdesc, float64(1828.64901))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Unpack(out, &v)
		Unpack(outdesc, &v)
	}
}

func BenchmarkStringDecoder(b *testing.B) {
	s := "now is the time for all good men"
	b.StopTimer()
	out := make([]byte, 256, 256)
	outdesc := make([]byte, 256, 256)
	var v string
	WithDesc(true).PackInto(outdesc, s)
	WithDesc(false).PackInto(out, s)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Unpack(out, &v)
		Unpack(outdesc, &v)
	}
}

func BenchmarkTimeEncoder(b *testing.B) {
	b.StopTimer()
	out := make([]byte, 16, 16)
	t := time.Now().Add(-10 * time.Hour).Add(-32 * time.Minute)
	args := []interface{}{t}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		WithDesc(true).PackSlice(out, args)
		WithDesc(false).PackSlice(out, args)
	}
}

func BenchmarkTimeDecoder(b *testing.B) {
	b.StopTimer()
	t := time.Now().Add(-10 * time.Hour).Add(-32 * time.Minute)
	out := make([]byte, 16, 16)
	outdesc := make([]byte, 16, 16)
	WithDesc(true).PackInto(outdesc, t)
	WithDesc(false).PackInto(out, t)
	var v time.Time
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Unpack(out, &v)
		Unpack(outdesc, &v)
	}
}

func BenchmarkPack(b *testing.B) {
	var (
		a int
		c int64
		d string
		e time.Time
	)
	for i := 0; i < b.N; i++ {
		v := Pack(1, int64(4), "this is a string", time.Now())
		Unpack(v, &a, &c, &d, &e)
	}

}
