package decoder

import (
	"unsafe"

	"github.com/iancmcc/bingo/internal/errors"
	"github.com/iancmcc/bingo/internal/types"
)

func Decode(b []byte, v interface{}) (int, error) {
	decoder, err := getDecoder(b[0])
	if err != nil {
		return 0, err
	}
	size, err := decoder(b, uintptr(unsafe.Pointer(&v)))
	if err != nil {
		return 0, err
	}
	return size, nil
}

// getDecoder returns the appropriate decoder for a given type byte, or an
// error if the type byte is unknown.
func getDecoder(b byte) (d Decoder, err error) {
	switch b {
	/*
		case typeByteNil, typeByteNilInverse:
			d = Encoder
		case typeByteBool, typeByteBoolInverse:
			n = SizeBool
		case typeByteUint8, typeByteUint8Inverse:
			n = SizeUint8
		case typeByteUint16, typeByteUint16Inverse:
			n = SizeUint16
		case typeByteUint32, typeByteUint32Inverse:
			n = SizeUint32
		case typeByteUint64, typeByteUint64Inverse:
			n = SizeUint64
	*/
	case types.TypeByteInt8, types.TypeByteInt8Inverse:
		d = DecodeInt8
	case types.TypeByteInt16, types.TypeByteInt16Inverse:
		d = DecodeInt16
		/*
			case typeByteInt32, typeByteInt32Inverse:
				n = SizeInt32
			case typeByteInt64, typeByteInt64Inverse:
				n = SizeInt64
			case typeByteFloat32, typeByteFloat32Inverse:
				n = SizeFloat32
			case typeByteFloat64, typeByteFloat64Inverse:
				n = SizeFloat64
			case typeByteTime, typeByteTimeInverse:
				n = SizeTime
			case typeByteString:
				n = bytes.IndexByte(b, terminatorByte) + 1
			case typeByteStringInverse:
				n = bytes.IndexByte(b, terminatorByteInverse) + 1
		*/
	default:
		err = errors.ErrUnknownTypeByte
	}
	return
}
