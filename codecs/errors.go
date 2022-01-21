package codecs

import "errors"

var (
	ErrNotAPointer   = errors.New("must pass a pointer to decode a value to")
	ErrUnknownType   = errors.New("unknown type")
	ErrByteArraySize = errors.New("receiving byte array doesn't have enough capacity")
	ErrNullByte      = errors.New("can't encode strings that contain a null byte")
	ErrInvalidTime   = errors.New("can't marshal to time.Time")
)
