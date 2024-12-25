package float

// Decode the encoded int.
func Int(b []byte, exponent int) int {
	keep := byte(0x7f) >> exponent
	val := byteToUintKeep(b, keep)
	if b[0]&0x7f&(^keep) == 0 {
		if 0x80&b[0] > 0 {
			return -int(val)
		}
		return int(val)
	}
	val = val | 1<<(len(b)*8-exponent-1)
	if 0x80&b[0] > 0 {
		return -int(val << ((0x7f & b[0]) >> (7 ^ exponent)) >> 1)
	}
	return int(val << ((0x7f & b[0]) >> (7 ^ exponent)) >> 1)
}

// Decode the encoded int.
func UInt(b []byte, exponent int) uint {
	keep := byte(0xff) >> exponent
	val := byteToUintKeep(b, keep)
	if b[0]&(^keep) == 0 {
		return val
	}
	val = val | 1<<((len(b)<<3)-exponent)
	return val << (b[0] >> (7 ^ exponent) >> 1) >> 1
}

// Encode an int value into an encoded int.
func PutInt(b []byte, val int, exponent int) {
	sign := val < 0
	if sign {
		val = -val
	}
	mantmax := (int(1) << ((len(b) << 3) - exponent))
	if mantmax > val {
		putUintRight(b, uint(val))
		if sign {
			b[0] |= 0x80
		}
		return
	}
	exp := byte(2)
	val = val >> 1
	for val >= mantmax {
		exp, val = exp+1, val>>1
	}
	val = val &^ (mantmax >> 1)
	putUintRight(b, uint(val))
	b[0] |= exp << (7 ^ exponent)
	if sign {
		b[0] |= 0x80
	}
}

// Encode an int value into an encoded int.
func PutUInt(b []byte, val uint, exponent int) {
	mantmax := (uint(1) << ((len(b) << 3) - exponent)) << 1
	if mantmax > val {
		putUintRight(b, val)
		return
	}
	exp := byte(2)
	val = val >> 1
	for val >= mantmax {
		exp, val = exp+1, val>>1
	}
	val = val &^ (mantmax >> 1)
	putUintRight(b, val)
	b[0] |= exp << (7 ^ exponent) << 1
}
