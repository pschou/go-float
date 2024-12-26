package float

// Decode the encoded int.
func UInt31(b []byte) uint {
	val := byteToUintKeep(b, 0xff)
	if val <= 1000 {
		return val
	}
	exp := uint(1)
	val = val - 1000
	for val > 1800 {
		exp, val = exp*10, val-1800
	}
	return (1000 + val*5) * exp
}

// Encode an int value into an encoded int.
func PutUInt31(b []byte, val uint) {
	if val <= 1000 {
		putUintRight(b, val)
		return
	}
	if val < 10000 {
		// One digit rounding 2->0,3->5
		putUintRight(b, (val-998)/5+1000)
		return
	}
	exp := uint(1)
	ret := uint(2800)
	for {
		if val < 100000*exp {
			// Two digit rounding 24->00,25->50
			putUintRight(b, (val-9975*exp)/(50*exp)+ret)
			return
		}
		exp, ret = exp*10, ret+1800
	}
}
