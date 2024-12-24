package float

// Decode the Int16 encoded int.
func DecodeInt16(b []byte, exponent int) uint {
	if b[0]>>(8-exponent) == 0 {
		return uint(b[0])<<8 | uint(b[1])
	}
	return uint(DecodeU16(uint16(b[0])<<8|uint16(b[1]), exponent, 15-exponent))
}

// Encode an int value into an Int16 encoded int.
func EncodeInt16(b []byte, val uint, exponent int) {
	if val < (0xffff >> exponent) {
		b[0] = byte(val >> 8)
		b[1] = byte(val)
		return
	}
	v := EncodeU16(float32(val), exponent, 15-exponent)
	b[0] = byte(v >> 8)
	b[1] = byte(v)
}
