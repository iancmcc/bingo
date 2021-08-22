package codecs

import (
	"math"
)

func invert(b byte) byte {
	return b ^ math.MaxUint8
}

func invertArray(a []byte) {
	for i, b := range a {
		a[i] = invert(b)
	}
}
