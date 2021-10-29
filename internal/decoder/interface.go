package decoder

// Decoder describes a function that can decode bingo-encoded bytes.
type Decoder func([]byte, uintptr) (int, error)
