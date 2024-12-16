package float

import (
	"encoding/binary"
	"math"
)

func JoinBytesAt(b []byte, exponent int) float64 {
	var in uint64
	switch len(b) {
	case 8:
		in = binary.BigEndian.Uint64(b)
	case 7:
		in = uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 | uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	case 6:
		in = uint64(b[5])<<16 | uint64(b[4])<<24 | uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	case 5:
		in = uint64(b[4])<<24 | uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	case 4:
		in = uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	case 3:
		in = uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	case 2:
		in = uint64(b[1])<<48 | uint64(b[0])<<56
	case 1:
		in = uint64(b[0]) << 56
	}

	sign := in & 0x8000000000000000
	in = in & 0x7fffffffffffffff

	// Find the first bit
	shift := int(63)
	for in>>shift == 0 {
		if shift == 0 {
			return 0
		}
		shift--
	}

	if shift > 52 {
		shift = shift - 52
		in = in >> shift
		exponent = exponent + shift + 1012
	} else {
		shift = 52 - shift
		in = in << shift
		exponent = exponent - shift + 1012
	}
	in = in&0x000fffffffffffff | sign | uint64(exponent)<<52
	return math.Float64frombits(in)
}

func Split32BytesAt(f float32, exponent int) []byte {
	in := math.Float32bits(f)
	sign := in & 0x80000000
	exp := int((in >> 23) & 0xff)

	// Wipe out the upper bits to leave the mantissa
	in = in&0x007fffff | 0x00800000
	shift := (-exp + 119 + exponent)
	if shift < -7 {
		panic("cannot fit mantissa into exponent")
	}
	if shift < 0 {
		in = in << -shift
	} else if shift > 0 {
		in = in >> shift
	}
	in = in | sign
	out := make([]byte, 4)
	binary.BigEndian.PutUint32(out, in)
	return out
}

func Split64BytesAt(f float64, exponent int) []byte {
	in := math.Float64bits(f)
	sign := in & 0x8000000000000000
	exp := int((in >> 52) & 0x07ff)

	// Wipe out the upper bits to leave the mantissa
	in = in&0x000fffffffffffff | 0x0010000000000000
	shift := (-exp + 1012 + exponent)
	if shift < -10 {
		panic("cannot fit mantissa into exponent")
	}
	if shift < 0 {
		in = in << -shift
	} else if shift > 0 {
		in = in >> shift
	}
	in = in | sign
	out := make([]byte, 8)
	binary.BigEndian.PutUint64(out, in)
	return out
}
