package bingo_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	. "github.com/iancmcc/bingo"
)

func BenchmarkCodecs(b *testing.B) {
	fns := []struct {
		name string
		v    interface{}
	}{
		{
			name: "int8",
			v:    int8(rand.Intn(math.MaxInt8)),
		}, {
			name: "int16",
			v:    int16(rand.Intn(math.MaxInt16)),
		}, {
			name: "int32",
			v:    int32(rand.Intn(math.MaxInt32)),
		}, {
			name: "int64",
			v:    int64(rand.Intn(math.MaxInt64)),
		}, {
			name: "float32",
			v:    float32(rand.Float32() * math.MaxInt16),
		}, {
			name: "float64",
			v:    float64(rand.Float64() * math.MaxInt32),
		}, {
			name: "string",
			v:    randomString(256),
		}, {
			name: "time",
			v:    time.Now(),
		},
	}
	for _, fn := range fns {
		b.Run(fn.name, func(b *testing.B) {
			buf := make([]byte, 1024, 1024)
			invbuf := make([]byte, 1024, 1024)
			b.Run("encode", func(b *testing.B) {
				b.Run("natural", func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for j := 0; j < b.N; j++ {
						PackInto(buf, fn.v)
					}
				})
				b.Run("inverse", func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for j := 0; j < b.N; j++ {
						WithDesc(true).PackInto(invbuf, fn.v)
					}
				})
			})
			b.Run("decode", func(b *testing.B) {
				b.Run("natural", func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for j := 0; j < b.N; j++ {
						Unpack(buf, &(fn.v))
					}
				})
				b.Run("inverse", func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for j := 0; j < b.N; j++ {
						Unpack(invbuf, &(fn.v))
					}
				})
			})
		})
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

//go:noinline
func randomString(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
