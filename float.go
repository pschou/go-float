package float

import (
	"encoding/binary"
	"math"
)

// This function takes a byte slice in BigEndian order and restores it to a
// float value with the predetermined upper exponent, the maximum value a
// numberic number can be.
func JoinBytesAt(b []byte, exponent int) float64 {
	in := byteToUint(b)
	return JoinAt(in, exponent)
}

// This function takes a uint64 and restores it to a float value with the
// predetermined upper exponent, the maximum value a numberic number can be.
func JoinAt(in uint64, exponent int) float64 {
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

// Take a float32 value and scale the mantissa to fit within a defined exponent
// value.  The byte string is left aligned and in BigEndian format so chomping
// off bits will decrease the precision.
func Split32BytesAt(f float32, exponent int) []byte {
	val := Split32At(f, exponent)
	out := make([]byte, 4)
	binary.BigEndian.PutUint32(out, val)
	return out
}

// Take a float32 value and scale the mantissa to fit within a defined exponent
// value.  The uint32 string is like a mantissa aligned at the largest big so
// shifting off lower bits (to the right) will decrease the precision.
func Split32At(f float32, exponent int) (val uint32) {
	val = math.Float32bits(f)
	sign := val & 0x80000000
	exp := int((val >> 23) & 0xff)

	// Wipe out the upper bits to leave the mantissa
	val = val&0x007fffff | 0x00800000
	shift := (-exp + 119 + exponent)
	if shift < -7 {
		panic("cannot fit mantissa into exponent")
	}
	if shift < 0 {
		val = val << -shift
	} else if shift > 0 {
		val = val >> shift
	}
	val = val | sign
	return
}

// Take a float64 value and scale the mantissa to fit within a defined exponent
// value.  The byte string is left aligned and in BigEndian format so chomping
// off bits will decrease the precision.
func Split64BytesAt(f float64, exponent int) []byte {
	val := Split64At(f, exponent)
	out := make([]byte, 8)
	binary.BigEndian.PutUint64(out, val)
	return out
}

// Take a float64 value and scale the mantissa to fit within a defined exponent
// value.  The uint64 string is like a mantissa aligned at the largest big so
// shifting off lower bits (to the right) will decrease the precision.
func Split64At(f float64, exponent int) (val uint64) {
	val = math.Float64bits(f)
	sign := val & 0x8000000000000000
	exp := int((val >> 52) & 0x07ff)

	// Wipe out the upper bits to leave the mantissa
	val = val&0x000fffffffffffff | 0x0010000000000000
	shift := (-exp + 1012 + exponent)
	if shift < -10 {
		panic("cannot fit mantissa into exponent")
	}
	if shift < 0 {
		val = val << -shift
	} else if shift > 0 {
		val = val >> shift
	}
	val = val | sign
	return
}

// Given two byte strings, overlap them so the center byte can be split and
// precision can be added to both numbers with a shared center byte.  This can
// be reversed by using the [Part] function.
func Overlap(a, b float64, exponentA, exponentB, length int) []byte {
	out := make([]byte, length)
	valA := Split64At(a, exponentA)
	putUint(out[:(length+1)/2], valA)
	valB := Split64At(b, exponentB)
	if length%2 > 0 {
		out[length/2] = out[length/2] & 0xf0
		valB = valB >> 4
	}
	orUint(out[length/2:], valB)
	return out
}

// Part will reverse an [Overlap] and restore the two floats.
func Part(in []byte, exponentA, exponentB int) (a, b float64) {
	var valA, valB uint64
	switch {
	case len(in) <= 8:
		val := byteToUint(in)
		mid := len(in) * 8 / 2
		valA = (val >> (64 - mid) << (64 - mid))
		valB = val << mid
	case len(in) <= 16:
		valA = byteToUint(in[:(len(in)+1)/2])
		valB = byteToUint(in[len(in)/2:])
		mid := len(in) * 8 / 2
		valA = (valA >> (64 - mid) << (64 - mid))
		valB = valB << 4
	default:
		valA = byteToUint(in[:(len(in)+1)/2])
		valB = byteToUint(in[len(in)/2:])
		valB = valB << 4
	}
	return JoinAt(valA, exponentA), JoinAt(valB, exponentB)
}

func byteToUint(b []byte) (val uint64) {
	switch len(b) {
	case 8:
		val = binary.BigEndian.Uint64(b)
	case 7:
		val = uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 | uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	case 6:
		val = uint64(b[5])<<16 | uint64(b[4])<<24 | uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	case 5:
		val = uint64(b[4])<<24 | uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	case 4:
		val = uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	case 3:
		val = uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	case 2:
		val = uint64(b[1])<<48 | uint64(b[0])<<56
	case 1:
		val = uint64(b[0]) << 56
	case 0:
		val = 0
	default:
		val = binary.BigEndian.Uint64(b[:8])
	}
	return
}
func orUint(b []byte, v uint64) {
	switch len(b) {
	case 8:
		b[0] |= byte(v >> 56)
		b[1] |= byte(v >> 48)
		b[2] |= byte(v >> 40)
		b[3] |= byte(v >> 32)
		b[4] |= byte(v >> 24)
		b[5] |= byte(v >> 16)
		b[6] |= byte(v >> 8)
		b[7] |= byte(v)
	case 7:
		b[0] |= byte(v >> 56)
		b[1] |= byte(v >> 48)
		b[2] |= byte(v >> 40)
		b[3] |= byte(v >> 32)
		b[4] |= byte(v >> 24)
		b[5] |= byte(v >> 16)
		b[6] |= byte(v >> 8)
	case 6:
		b[0] |= byte(v >> 56)
		b[1] |= byte(v >> 48)
		b[2] |= byte(v >> 40)
		b[3] |= byte(v >> 32)
		b[4] |= byte(v >> 24)
		b[5] |= byte(v >> 16)
	case 5:
		b[0] |= byte(v >> 56)
		b[1] |= byte(v >> 48)
		b[2] |= byte(v >> 40)
		b[3] |= byte(v >> 32)
		b[4] |= byte(v >> 24)
	case 4:
		b[0] |= byte(v >> 56)
		b[1] |= byte(v >> 48)
		b[2] |= byte(v >> 40)
		b[3] |= byte(v >> 32)
	case 3:
		b[0] |= byte(v >> 56)
		b[1] |= byte(v >> 48)
		b[2] |= byte(v >> 40)
	case 2:
		b[0] |= byte(v >> 56)
		b[1] |= byte(v >> 48)
	case 1:
		b[0] |= byte(v >> 56)
	}
	return
}

func putUint(b []byte, v uint64) {
	switch len(b) {
	case 8:
		b[0] = byte(v >> 56)
		b[1] = byte(v >> 48)
		b[2] = byte(v >> 40)
		b[3] = byte(v >> 32)
		b[4] = byte(v >> 24)
		b[5] = byte(v >> 16)
		b[6] = byte(v >> 8)
		b[7] = byte(v)
	case 7:
		b[0] = byte(v >> 56)
		b[1] = byte(v >> 48)
		b[2] = byte(v >> 40)
		b[3] = byte(v >> 32)
		b[4] = byte(v >> 24)
		b[5] = byte(v >> 16)
		b[6] = byte(v >> 8)
	case 6:
		b[0] = byte(v >> 56)
		b[1] = byte(v >> 48)
		b[2] = byte(v >> 40)
		b[3] = byte(v >> 32)
		b[4] = byte(v >> 24)
		b[5] = byte(v >> 16)
	case 5:
		b[0] = byte(v >> 56)
		b[1] = byte(v >> 48)
		b[2] = byte(v >> 40)
		b[3] = byte(v >> 32)
		b[4] = byte(v >> 24)
	case 4:
		b[0] = byte(v >> 56)
		b[1] = byte(v >> 48)
		b[2] = byte(v >> 40)
		b[3] = byte(v >> 32)
	case 3:
		b[0] = byte(v >> 56)
		b[1] = byte(v >> 48)
		b[2] = byte(v >> 40)
	case 2:
		b[0] = byte(v >> 56)
		b[1] = byte(v >> 48)
	case 1:
		b[0] = byte(v >> 56)
	}
	return
}
