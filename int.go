package float

// Decode the encoded int.
func Int(b []byte, exponent int) uint {
	switch {
	case len(b) == 1:
		if b[0]>>(8-exponent) == 0 {
			return uint(b[0]<<exponent) >> exponent
		}
		return uint(FromU16(uint16(b[0])<<8, exponent, 7-exponent))
	case len(b) == 2:
		if b[0]>>(8-exponent) == 0 {
			return uint(b[0])<<8 | uint(b[1])
		}
		return uint(FromU16(uint16(b[0])<<8|uint16(b[1]), exponent, 15-exponent))
	}
	return 0
}

// Encode an int value into an encoded int.
func PutInt(b []byte, val uint, exponent int) {
	switch {
	case len(b) == 1:
		if val < (0xff >> exponent) {
			b[0] = byte(val)
			return
		}
		v := ToU16(float32(val), exponent, 7-exponent)
		b[0] = byte(v >> 8)
	case len(b) == 2:
		if val < (0xffff >> exponent) {
			b[0] = byte(val >> 8)
			b[1] = byte(val)
			return
		}
		v := ToU16(float32(val), exponent, 15-exponent)
		b[0] = byte(v >> 8)
		b[1] = byte(v)
	}
}
