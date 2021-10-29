package errors

import "errors"

var (
	ErrUnsupportedType = errors.New("unsupported type")
	ErrInvalidEncoding = errors.New("invalid encoding")
	ErrUnknownTypeByte = errors.New("unknown type byte")
	ErrArrayTooSmall   = errors.New("array does not have sufficient space")
)
