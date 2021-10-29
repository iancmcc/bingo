package decoder_test

import (
	"fmt"
	"testing"

	"github.com/iancmcc/bingo/internal/decoder"
	"github.com/iancmcc/bingo/internal/encoder"
)

func TestAThing(t *testing.T) {
	b := make([]byte, 10)
	_, err := encoder.EncodeInt8(b, int8(64), false)
	if err != nil {
		t.Fatalf("error was thrown: %v", err)
	}
	fmt.Printf("%v\n", b)
	var result int8
	_, err = decoder.Decode(b, &result)
	if err != nil {
		t.Fatalf("error was thrown: %v", err)
	}
	if result != int8(64) {
		t.Fatalf("wrong result: %v", result)
	}

}
