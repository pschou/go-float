package float

// Decode the encoded int.
func UInt30(b []byte) uint {
	val := byteToUintKeep(b, 0xff)
	if val <= 1000 {
		return val
	}
	exp := uint(10)
	val = val - 1000
	for val > 900 {
		exp, val = exp*10, val-900
	}
	return (100 + val) * exp
}

// Encode an int value into an encoded int.
func PutUInt30(b []byte, val uint) {
	if val <= 1000 {
		putUintRight(b, val)
		return
	}
	ret := uint(1000)
	val = (val + 5) / 10
	for val > 999 {
		ret, val = ret+900, (val+5)/10
	}
	ret += val - 100
	putUintRight(b, ret)
}
