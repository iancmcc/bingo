package codecs

import (
	"errors"
	"time"
	"unsafe"

	"github.com/goccy/go-reflect"
	"github.com/iancmcc/bingo/bytes"
)

const (
	typeByteTime        byte = 0x35
	typeByteTimeInverse      = typeByteTime ^ 0xff
	sizeTime                 = 16
)

func encodeTime(b []byte, v time.Time, inverse bool) (int, error) {
	if cap(b) < sizeTime {
		return 0, ErrByteSliceSize
	}
	b = b[:sizeTime]
	b[0] = typeByteTime
	if err := timeMarshalBinary(b[1:], v); err != nil {
		return 0, ErrInvalidTime
	}
	if inverse {
		bytes.InvertArraySmall(b)
	}
	return sizeTime, nil
}

func decodeTime(b []byte, v reflect.Value) (int, error) {
	encoded := b[1:sizeTime]
	if b[0] == typeByteTimeInverse {
		encoded = make([]byte, 15)
		copy(encoded, b[1:16])
		bytes.InvertArraySmall(encoded)
	}
	var t time.Time
	err := t.UnmarshalBinary(encoded)
	if err != nil {
		return 0, err
	}
	ptr := v.Pointer()
	**(**time.Time)(unsafe.Pointer(&ptr)) = *(*time.Time)(unsafe.Pointer(&t))
	return sizeTime, nil
}

//go:linkname time_sec time.(*Time).sec
//go:noescape
func time_sec(*time.Time) int64

//go:linkname time_nsec time.(*Time).nsec
//go:noescape
func time_nsec(*time.Time) int32

const timeBinaryVersion byte = 1

func timeMarshalBinary(b []byte, t time.Time) error {
	var offsetMin int16 // minutes east of UTC. -1 is UTC.

	if t.Location() == time.UTC {
		offsetMin = -1
	} else {
		_, offset := t.Zone()
		if offset%60 != 0 {
			return errors.New("Time.MarshalBinary: zone offset has fractional minute")
		}
		offset /= 60
		if offset < -32768 || offset == -1 || offset > 32767 {
			return errors.New("Time.MarshalBinary: unexpected zone offset")
		}
		offsetMin = int16(offset)
	}

	sec := time_sec(&t)
	nsec := time_nsec(&t)

	b[0] = timeBinaryVersion // byte 0 : version
	b[1] = byte(sec >> 56)   // bytes 1-8: seconds
	b[2] = byte(sec >> 48)
	b[3] = byte(sec >> 40)
	b[4] = byte(sec >> 32)
	b[5] = byte(sec >> 24)
	b[6] = byte(sec >> 16)
	b[7] = byte(sec >> 8)
	b[8] = byte(sec)
	b[9] = byte(nsec >> 24) // bytes 9-12: nanoseconds
	b[10] = byte(nsec >> 16)
	b[11] = byte(nsec >> 8)
	b[12] = byte(nsec)
	b[13] = byte(offsetMin >> 8) // bytes 13-14: zone offset in minutes
	b[14] = byte(offsetMin)

	return nil
}
