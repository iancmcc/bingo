package bytes

import (
	"encoding/binary"

	// import for linking
	_ "unsafe"
	// import for linking
	_ "crypto/cipher"
)

var word uint64 = 0xffffffffffffffff
var maxarray = make([]byte, 4096)

func init() {
	for i := 0; i < 4096; i++ {
		maxarray[i] = 0xff
	}
}

//go:linkname xorBytes crypto/cipher.xorBytes
func xorBytes(dst, a, b []byte) int

// InvertByte inverts a single byte.
func InvertByte(b byte) byte {
	return b ^ 0xff
}

// InvertArrayLarge inverts the byte array using the optimized xorBytes
// function from the crypto/cipher package. This is fastest for byte arrays
// larger than 128 bytes.
func InvertArrayLarge(a []byte) {
	for len(a) >= 4096 {
		xorBytes(a[:4096], a[:4096], maxarray)
		a = a[4096:]
	}
	xorBytes(a, a, maxarray)
}

// InvertArray inverts the byte array, choosing the optimal inversion function.
func InvertArray(a []byte) {
	if len(a) <= 128 {
		InvertArraySmall(a)
	} else {
		InvertArrayLarge(a)
	}
}

// InvertArraySmall inverts a byte array a word at a time, which is fastest up
// to 128 bytes.
func InvertArraySmall(b []byte) {
	if len(b) >= 8 {
		for len(b) >= 128 {
			v := binary.LittleEndian.Uint64(b)
			binary.LittleEndian.PutUint64(b, v^word)
			v = binary.LittleEndian.Uint64(b[8:])
			binary.LittleEndian.PutUint64(b[8:], v^word)
			v = binary.LittleEndian.Uint64(b[16:])
			binary.LittleEndian.PutUint64(b[16:], v^word)
			v = binary.LittleEndian.Uint64(b[24:])
			binary.LittleEndian.PutUint64(b[24:], v^word)
			v = binary.LittleEndian.Uint64(b[32:])
			binary.LittleEndian.PutUint64(b[32:], v^word)
			v = binary.LittleEndian.Uint64(b[40:])
			binary.LittleEndian.PutUint64(b[40:], v^word)
			v = binary.LittleEndian.Uint64(b[48:])
			binary.LittleEndian.PutUint64(b[48:], v^word)
			v = binary.LittleEndian.Uint64(b[56:])
			binary.LittleEndian.PutUint64(b[56:], v^word)
			v = binary.LittleEndian.Uint64(b[64:])
			binary.LittleEndian.PutUint64(b[64:], v^word)
			v = binary.LittleEndian.Uint64(b[72:])
			binary.LittleEndian.PutUint64(b[72:], v^word)
			v = binary.LittleEndian.Uint64(b[80:])
			binary.LittleEndian.PutUint64(b[80:], v^word)
			v = binary.LittleEndian.Uint64(b[88:])
			binary.LittleEndian.PutUint64(b[88:], v^word)
			v = binary.LittleEndian.Uint64(b[96:])
			binary.LittleEndian.PutUint64(b[96:], v^word)
			v = binary.LittleEndian.Uint64(b[104:])
			binary.LittleEndian.PutUint64(b[104:], v^word)
			v = binary.LittleEndian.Uint64(b[112:])
			binary.LittleEndian.PutUint64(b[112:], v^word)
			v = binary.LittleEndian.Uint64(b[120:])
			binary.LittleEndian.PutUint64(b[120:], v^word)
			b = b[128:]
		}
		for len(b) >= 64 {
			v := binary.LittleEndian.Uint64(b)
			binary.LittleEndian.PutUint64(b, v^word)
			v = binary.LittleEndian.Uint64(b[8:])
			binary.LittleEndian.PutUint64(b[8:], v^word)
			v = binary.LittleEndian.Uint64(b[16:])
			binary.LittleEndian.PutUint64(b[16:], v^word)
			v = binary.LittleEndian.Uint64(b[24:])
			binary.LittleEndian.PutUint64(b[24:], v^word)
			v = binary.LittleEndian.Uint64(b[32:])
			binary.LittleEndian.PutUint64(b[32:], v^word)
			v = binary.LittleEndian.Uint64(b[40:])
			binary.LittleEndian.PutUint64(b[40:], v^word)
			v = binary.LittleEndian.Uint64(b[48:])
			binary.LittleEndian.PutUint64(b[48:], v^word)
			v = binary.LittleEndian.Uint64(b[56:])
			binary.LittleEndian.PutUint64(b[56:], v^word)
			b = b[64:]
		}
		for len(b) >= 32 {
			v := binary.LittleEndian.Uint64(b)
			binary.LittleEndian.PutUint64(b, v^word)
			v = binary.LittleEndian.Uint64(b[8:])
			binary.LittleEndian.PutUint64(b[8:], v^word)
			v = binary.LittleEndian.Uint64(b[16:])
			binary.LittleEndian.PutUint64(b[16:], v^word)
			v = binary.LittleEndian.Uint64(b[24:])
			binary.LittleEndian.PutUint64(b[24:], v^word)
			b = b[32:]
		}
		for len(b) >= 16 {
			v := binary.LittleEndian.Uint64(b)
			binary.LittleEndian.PutUint64(b, v^word)
			v = binary.LittleEndian.Uint64(b[8:])
			binary.LittleEndian.PutUint64(b[8:], v^word)
			b = b[16:]
		}
	}
	for len(b) >= 8 {
		v := binary.LittleEndian.Uint64(b)
		binary.LittleEndian.PutUint64(b, v^word)
		b = b[8:]
	}
	for i := range b {
		b[i] ^= 0xff
	}
}
