package float

import (
	"encoding/binary"
	"math"
)

// This function takes a byte slice in BigEndian order and restores it to a
// float value with the predetermined upper exponent, the maximum value a
// numberic number can be.
func Scaled(b []byte, exponent int) float64 {
	in := byteToUint(b)
	return FromScaled64(in, exponent)
}

// This function takes a byte slice in BigEndian order and restores it to a
// float value with the predetermined upper exponent, the maximum value a
// numberic number can be.
func UScaled(b []byte, exponent int) float64 {
	in := byteToUint(b)
	return FromUScaled64(in, exponent)
}

// This function takes a uint32 and restores it to an unsigned float value with
// the predetermined upper exponent, the upper numeric bound.
func FromUScaled32(in uint32, exponent int) float32 {
	// TODO: Consider re-writing this for 32bit processors
	return float32(FromUScaled64(uint64(in)<<32, exponent))
}

// This function takes a uint32 and restores it to a signed float value with
// the predetermined upper exponent, the upper numeric bound.
func FromScaled32(in uint32, exponent int) float32 {
	// TODO: Consider re-writing this for 32bit processors
	return float32(FromScaled64(uint64(in)<<32, exponent))
}

// This function takes a uint64 and restores it to a signed float value with
// the predetermined upper exponent, the upper numeric bound.
func FromScaled64(in uint64, exponent int) float64 {
	sign := in & 0x8000000000000000
	in = in & 0x7fffffffffffffff

	// Find the first bit
	shift := int(62)
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

// This function takes a uint64 and restores it to an unsigned float value with
// the predetermined upper exponent, the upper numeric bound.
func FromUScaled64(in uint64, exponent int) float64 {
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
		exponent = exponent + shift + 1011
	} else {
		shift = 52 - shift
		in = in << shift
		exponent = exponent - shift + 1011
	}
	in = in&0x000fffffffffffff | uint64(exponent)<<52
	return math.Float64frombits(in)
}

// Take a float32 value and scale the mantissa to fit within a defined exponent
// value.  The byte string is left aligned and in BigEndian format so chomping
// off bits will decrease the precision.
//
// When less than 4 bytes are provided the mantissa will be chomped down to fit
// within the slice provided.  This way one can make a Float24 or Float16 and
// use the [Scaled] to reverse.
func PutScaled32(out []byte, f float32, exponent int) {
	val := ToScaled32(f, exponent)
	putUint32(out, val)
	return
}

// Take a positive only float32 value and scale the mantissa to fit within a
// defined exponent value.  The byte string is left aligned and in BigEndian
// format so chomping off bits will decrease the precision.  By omitting the
// sign one can save a bit and use it to increase precision.
//
// When less than 4 bytes are provided the mantissa will be chomped down to fit
// within the slice provided.  This way one can make a Float24 or Float16 and
// use the [Scaled] to reverse.
func PutUScaled32(out []byte, f float32, exponent int) {
	val := UScaled32(f, exponent)
	putUint32(out, val)
	return
}

// Take a float32 value and scale the mantissa to fit within a defined exponent
// value.  The uint32 string is like a mantissa aligned at the largest big so
// shifting off lower bits (to the right) will decrease the precision.
func ToScaled32(f float32, exponent int) (val uint32) {
	val = math.Float32bits(f)
	sign := val & 0x80000000
	exp := int((val >> 23) & 0xff)

	// Wipe out the upper bits to leave the mantissa
	val = val&0x007fffff | 0x00800000
	shift := (-exp + 119 + exponent)
	if shift < -7 {
		return 0x7fffffff | sign
	}
	if shift < 0 {
		val = val << -shift
	} else if shift > 0 {
		val = val >> shift
	}
	val = val | sign
	return
}

// Take a float32 value and scale the mantissa to fit within a defined exponent
// value with out regard to sign.  The uint32 string is like a mantissa aligned
// at the largest big so shifting off lower bits (to the right) will decrease
// the precision.
func UScaled32(f float32, exponent int) (val uint32) {
	val = math.Float32bits(f)
	exp := int((val >> 23) & 0xff)

	shift := (-exp + 118 + exponent)
	if shift < -8 {
		return 0xffffffff
	}
	// Wipe out the upper bits to leave the mantissa
	val = val&0x007fffff | 0x00800000
	if shift < 0 {
		val = val << -shift
	} else if shift > 0 {
		val = val >> shift
	}
	return
}

// Take a float64 value and scale the mantissa to fit within a defined exponent
// value.  The byte string is left aligned and in BigEndian format so chomping
// off bits will decrease the precision.
//
// When less than 8 bytes are provided the mantissa will be chomped down to fit
// within the slice provided.  This way one can make a Float40 or higher and
// use the [Scaled] to reverse.
func PutUScaled64(out []byte, f float64, exponent int) {
	val := ToUScaled64(f, exponent)
	putUint(out, val)
	return
}

// Take a float64 value and scale the mantissa to fit within a defined exponent
// value.  The byte string is left aligned and in BigEndian format so chomping
// off bits will decrease the precision.
//
// When less than 8 bytes are provided the mantissa will be chomped down to fit
// within the slice provided.  This way one can make a Float40 or higher and
// use the [Scaled] to reverse.
func PutScaled64(out []byte, f float64, exponent int) {
	val := ToScaled64(f, exponent)
	putUint(out, val)
	return
}

// Take a float64 value and scale the mantissa to fit within a defined exponent
// value.  The uint64 string is like a mantissa aligned at the largest big so
// shifting off lower bits (to the right) will decrease the precision.
func ToUScaled64(f float64, exponent int) (val uint64) {
	val = math.Float64bits(f)
	exp := int((val >> 52) & 0x07ff)

	shift := (-exp + 1011 + exponent)
	if shift < -11 {
		return 0xffffffffffffffff
	}
	// Wipe out the upper bits to leave the mantissa
	val = val&0x000fffffffffffff | 0x0010000000000000
	if shift < 0 {
		val = val << -shift
	} else if shift > 0 {
		val = val >> shift
	}
	return
}

// Take a float64 value and scale the mantissa to fit within a defined exponent
// value.  The uint64 string is like a mantissa aligned at the largest big so
// shifting off lower bits (to the right) will decrease the precision.
func ToScaled64(f float64, exponent int) (val uint64) {
	val = math.Float64bits(f)
	sign := val & 0x8000000000000000
	exp := int((val >> 52) & 0x07ff)

	shift := (-exp + 1012 + exponent)
	if shift < -10 {
		return 0x7fffffffffffffff | sign
	}
	// Wipe out the upper bits to leave the mantissa
	val = val&0x000fffffffffffff | 0x0010000000000000
	if shift < 0 {
		val = val << -shift
	} else if shift > 0 {
		val = val >> shift
	}
	val = val | sign
	return
}

func byteToUintKeep(b []byte, keep byte) (val uint) {
	switch len(b) {
	case 7:
		val = uint(b[0]&keep)<<48 | uint(b[1])<<40 | uint(b[2])<<32 | uint(b[3])<<24 | uint(b[4])<<16 | uint(b[5])<<8 | uint(b[6])
	case 6:
		val = uint(b[0]&keep)<<40 | uint(b[1])<<32 | uint(b[2])<<24 | uint(b[3])<<16 | uint(b[4])<<8 | uint(b[5])
	case 5:
		val = uint(b[0]&keep)<<32 | uint(b[1])<<24 | uint(b[2])<<16 | uint(b[3])<<8 | uint(b[4])
	case 4:
		val = uint(b[0]&keep)<<24 | uint(b[1])<<16 | uint(b[2])<<8 | uint(b[3])
	case 3:
		val = uint(b[0]&keep)<<16 | uint(b[1])<<8 | uint(b[2])
	case 2:
		val = uint(b[0]&keep)<<8 | uint(b[1])
	case 1:
		val = uint(b[0] & keep)
	case 0:
		val = 0
	default:
		val = uint(b[7]) | uint(b[6])<<8 | uint(b[5])<<16 | uint(b[4])<<24 | uint(b[3])<<32 | uint(b[2])<<40 | uint(b[1])<<48 | uint(b[0]&keep)<<56
	}
	return
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

func putUintRight(b []byte, v uint) {
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
		b[0] = byte(v >> 48)
		b[1] = byte(v >> 40)
		b[2] = byte(v >> 32)
		b[3] = byte(v >> 24)
		b[4] = byte(v >> 16)
		b[5] = byte(v >> 8)
		b[6] = byte(v)
	case 6:
		b[0] = byte(v >> 40)
		b[1] = byte(v >> 32)
		b[2] = byte(v >> 24)
		b[3] = byte(v >> 16)
		b[4] = byte(v >> 8)
		b[5] = byte(v)
	case 5:
		b[0] = byte(v >> 32)
		b[1] = byte(v >> 24)
		b[2] = byte(v >> 16)
		b[3] = byte(v >> 8)
		b[4] = byte(v)
	case 4:
		b[0] = byte(v >> 24)
		b[1] = byte(v >> 16)
		b[2] = byte(v >> 8)
		b[3] = byte(v)
	case 3:
		b[0] = byte(v >> 16)
		b[1] = byte(v >> 8)
		b[2] = byte(v)
	case 2:
		b[0] = byte(v >> 8)
		b[1] = byte(v)
	case 1:
		b[0] = byte(v)
	}
	return
}

func putUint32(b []byte, v uint32) {
	switch len(b) {
	case 4:
		b[0] = byte(v >> 24)
		b[1] = byte(v >> 16)
		b[2] = byte(v >> 8)
		b[3] = byte(v)
	case 3:
		b[0] = byte(v >> 24)
		b[1] = byte(v >> 16)
		b[2] = byte(v >> 8)
	case 2:
		b[0] = byte(v >> 24)
		b[1] = byte(v >> 16)
	case 1:
		b[0] = byte(v >> 24)
	}
	return
}
