package float

// Decode the encoded int.
func UInt40(b []byte) uint {
	val := byteToUintKeep(b, 0xff)
	if val <= 10000 {
		return val
	}
	exp := uint(10)
	val = val - 10000
	for val > 9000 {
		exp, val = exp*10, val-9000
	}
	return (1000 + val) * exp
}

// Encode an int value into an encoded int.
func PutUInt40(b []byte, val uint) {
	if val <= 10000 {
		putUintRight(b, val)
		return
	}
	ret := uint(10000)
	val = (val + 5) / 10
	for val > 9999 {
		ret, val = ret+9000, (val+5)/10
	}
	ret += val - 1000
	putUintRight(b, ret)
}
