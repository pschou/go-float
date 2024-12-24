package float

import (
	"math"
)

// To determine the limits of the float16's [Encode16] and [Decode16] one can
// use [Limits16] to derive the limits of the notation.  Note that zero is never
// represented because Float16 expects the values to be on a Exponent with a 1
// prefixed Mantissa scale with sign bit.  It is up to the user to account for
// the minimum value by rounding down to zero when appropriate.
//
// NOTE: It is up to the user to account for the minimum value by rounding down
// to zero when appropriate.
func Limits16(expBits, expShift int) (lower, upper float32) {
	return Decode16(0x0000, expBits, expShift),
		Decode16(0x7fff, expBits, expShift)
}

// To determine the limits of the float16's [EncodeU16] and [DecodeU16] one
// can use [LimitsU16] to derive the limits of the notation.  Note that zero is
// never represented because Float16 expects the values to be on a Exponent
// with a 1 prefixed Mantissa scale.
//
// NOTE: It is up to the user to account for the minimum value by rounding down
// to zero when appropriate.
func LimitsU16(expBits, expShift int) (lower, upper float32) {
	return DecodeU16(0x0000, expBits, expShift),
		DecodeU16(0xffff, expBits, expShift)
}

// Encode a float32 down to fit within 2 bytes.  As this function can have
// very hard limits, one should predefine the expected range using the expBits,
// the number of bits used to express the mantissa.  The sign bit is included
// and will account for positive and negative values.  Use the [Decode16] to
// reverse this transformation.
//
// 0 bit -> numerical range of 1 (0.5 to 0.999)
// 1 bit -> 2 (0.5 to 1.999)
// 2 bit -> 4 (0.5 to 3.999)
// 3 bit -> 8 (0.5 to 7.99), etc...
//
// NOTE: It is up to the user to account for the minimum value by rounding down
// to zero when appropriate.
func Encode16(f float32, expBits, expShift int) uint16 {
	in := math.Float32bits(f)
	sign := in & 0x80000000
	max := 1 << (expBits - 1)
	exp := int(in&0x7ff80000)>>23 - 127 - expShift
	if exp > max {
		return 0x7fff | uint16(sign>>16)
	} else if exp < 0 {
		return uint16(sign >> 16)
	}

	in = (in&0x007fffff)>>(8+expBits) | sign>>16 | uint32(exp)<<(15-expBits)
	return uint16(in)
}

// Encode a float32 down to fit within 2 bytes.  As this function can have
// very hard limits, one should predefine the expected range using the expBits,
// the number of bits used to express the mantissa.  Use the [DecodeU16] to
// reverse this transformation.
//
// 0 bit -> numerical range of 1 (0.5 to 0.999)
// 1 bit -> 2 (0.5 to 1.999)
// 2 bit -> 4 (0.5 to 3.999)
// 3 bit -> 8 (0.5 to 7.99), etc...
//
// NOTE: It is up to the user to account for the minimum value by rounding down
// to zero when appropriate.
func EncodeU16(f float32, expBits, expShift int) uint16 {
	in := math.Float32bits(f)
	max := 1 << (expBits)
	exp := int(in&0x7ff80000)>>23 - 127 - expShift
	if exp > max {
		return 0xffff
	} else if exp < 0 {
		return 0x0000
	}

	in = (in&0x007fffff)>>(7+expBits) | uint32(exp)<<(16-expBits)
	return uint16(in)
}

// Use this to expand the [Encode16] and restore the number back to like the
// original, the represented value.
func Decode16(in uint16, expBits, expShift int) float32 {
	sign := uint32(in&0x8000) << 16
	exp := (int(0x7fff&in) >> (15 - expBits)) + 127 + expShift
	if exp > 189 {
		exp, in = 189, 0xffff
	}
	return math.Float32frombits(uint32(in<<(expBits+1))<<7 | sign | uint32(exp<<23))
}

// Use this to expand the [EncodeU16] and restore the number back to like the
// original, the represented value.
func DecodeU16(in uint16, expBits, expShift int) float32 {
	exp := int(in>>(16-expBits)) + 127 + expShift
	if exp > 189 {
		exp, in = 189, 0xffff
	}
	return math.Float32frombits(uint32(in<<(expBits))<<7 | uint32(exp)<<23)
}
