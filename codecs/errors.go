package codecs

import "errors"

var (
	// ErrInvalidTarget is returned when the target is not a pointer.
	ErrInvalidTarget = errors.New("must pass a pointer to decode a value to")
	// ErrUnknownType is returned when the type is unknown.
	ErrUnknownType = errors.New("unknown type")
	// ErrByteSliceSize is returned when the byte slice provided doesn't have
	// enough capacity to hold the encoded value.
	ErrByteSliceSize = errors.New("receiving byte array doesn't have enough capacity")
	// ErrNullByte is returned when a string to be encoded contains a null byte.
	ErrNullByte = errors.New("can't encode strings that contain a null byte")
	// ErrInvalidTime is returned when an encoded time can't be marshaled to time.Time.
	ErrInvalidTime = errors.New("can't marshal to time.Time")
)
