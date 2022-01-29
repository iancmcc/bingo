package bingo_test

import (
	"bytes"
	"fmt"
	"sort"
	"time"

	"github.com/iancmcc/bingo"
)

func ExamplePack() {
	// Pack some values into keys
	a, _ := bingo.Pack(1, "b", float32(4.2))
	b, _ := bingo.Pack(1, "b", float32(3.1))
	c, _ := bingo.Pack(1, "a", float32(7.5))

	// Sort the resulting byte slices
	packed := [][]byte{a, b, c}
	sort.SliceStable(packed, func(p, q int) bool {
		return bytes.Compare(packed[p], packed[q]) < 0
	})

	// Unpack again and witness the glorious order
	var (
		first  int
		second string
		third  float32
	)
	for _, bs := range packed {
		bingo.Unpack(bs, &first, &second, &third)
		fmt.Println(first, second, third)
	}
	// Output: 1 a 7.5
	// 1 b 3.1
	// 1 b 4.2
}

func ExampleWithDesc() {
	// Create a schema that packs the second value in descending order
	schema := bingo.WithDesc(false, true, false)

	// Pack some values into keys
	a, _ := schema.Pack(1, "b", float32(4.2))
	b, _ := schema.Pack(1, "b", float32(3.1))
	c, _ := schema.Pack(1, "a", float32(7.5))

	// Sort the resulting byte slices
	packed := [][]byte{a, b, c}
	sort.SliceStable(packed, func(p, q int) bool {
		return bytes.Compare(packed[p], packed[q]) < 0
	})

	// Unpack to see the order
	var (
		first  int
		second string
		third  float32
	)
	for _, bs := range packed {
		bingo.Unpack(bs, &first, &second, &third)
		fmt.Println(first, second, third)
	}
	// Output: 1 b 3.1
	// 1 b 4.2
	// 1 a 7.5
}

func ExamplePackTo() {
	// Pack requires an allocation for the resulting byte array, but you can
	// also pack to a byte array you already have

	values := []interface{}{1, "a", time.Now()}

	// Get the size you'll need
	size, _ := bingo.PackedSize(values)
	dest := make([]byte, size)

	bingo.PackTo(dest, values...)
}

func ExampleWritePackedTo() {
	var buf bytes.Buffer

	bingo.WritePackedTo(&buf, 1, "a")

	var (
		first  int
		second string
	)
	bingo.Unpack(buf.Bytes(), &first, &second)

	fmt.Println(first, second)
	// Output: 1 a
}
