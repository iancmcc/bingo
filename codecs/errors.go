package codecs

import "errors"

var (
	ErrNotAPointer   = errors.New("must pass a pointer to decode a value to")
	ErrUnknownType   = errors.New("unknown type")
	ErrByteArraySize = errors.New("receiving byte array doesn't have enough capacity")
)
