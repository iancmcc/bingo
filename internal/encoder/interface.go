package encoder

// Encoder describes a function that can encode to bytes.
type Encoder func([]byte, interface{}, bool) (int, error)
