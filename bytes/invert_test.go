package bytes_test

import (
	"bytes"
	"math/rand"
	"strconv"
	"testing"

	. "github.com/iancmcc/bingo/bytes"
)

type inverter func([]byte)

func testInverter(name string, f inverter, t *testing.T) {
	a := make([]byte, 2)
	b := make([]byte, 2)
	rand.Read(a)
	for i := range a {
		b[i] = a[i] ^ 0xff
	}
	f(a)
	if !bytes.Equal(a, b) {
		t.Fatalf("%s failed to invert properly", name)
	}
}

func TestInvertByte(t *testing.T) {
	a := rand.Intn(256)
	b := InvertByte(byte(a))
	if b^0xff != byte(a) {
		t.Fatalf("InvertByte failed to invert properly")
	}
}

func TestInvertSmall(t *testing.T) {
	testInverter("InvertArraySmall", InvertArraySmall, t)
}

func TestInvertLarge(t *testing.T) {
	testInverter("InvertArrayLarge", InvertArrayLarge, t)
}

func TestInvertOptimized(t *testing.T) {
	testInverter("InvertArray", InvertArray, t)
}

var sizes = []int{
	2,
	3,
	4,
	8,
	16,
	32,
	64,
	128,
	512,
	1024,
	4096,
	16384,
}

func TestInvert(t *testing.T) {
	fns := []struct {
		name string
		fn   inverter
	}{
		{
			name: "small",
			fn:   InvertArraySmall,
		},
		{
			name: "large",
			fn:   InvertArrayLarge,
		},
		{
			name: "default",
			fn:   InvertArray,
		},
	}

	for _, size := range sizes {
		p := make([]byte, size)
		rand.Read(p)
		cmp := make([]byte, size)
		for i := range p {
			cmp[i] = p[i] ^ 0xff
		}
		t.Run(strconv.Itoa(size), func(t *testing.T) {
			for _, fn := range fns {
				t.Run(fn.name, func(t *testing.T) {
					cp := make([]byte, size)
					copy(cp, p)
					fn.fn(cp)
					if !bytes.Equal(cp, cmp) {
						t.Fatalf("%v != %v", cp, cmp)
					}
				})
			}
		})
	}

}

func BenchmarkInvert(b *testing.B) {
	fns := []struct {
		name string
		fn   func(b *testing.B, a []byte)
	}{
		{
			name: "default",
			fn: func(b *testing.B, a []byte) {
				for i := 0; i < b.N; i++ {
					InvertArray(a)
				}
			},
		},
		{
			name: "small",
			fn: func(b *testing.B, a []byte) {
				for i := 0; i < b.N; i++ {
					InvertArraySmall(a)
				}
			},
		},
		{
			name: "large",
			fn: func(b *testing.B, a []byte) {
				for i := 0; i < b.N; i++ {
					InvertArrayLarge(a)
				}
			},
		},
	}

	for _, size := range sizes {
		p := make([]byte, size)
		rand.Read(p)
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			for _, fn := range fns {
				b.Run(fn.name, func(b *testing.B) {
					b.ReportAllocs()
					b.SetBytes(int64(size))
					b.ResetTimer()
					fn.fn(b, p)
				})
			}
		})
	}
}
