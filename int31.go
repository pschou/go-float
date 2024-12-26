package float

// Decode the encoded Int31.  The Int31 type is an integer with 3 significant
// figures and an additional half.  See [PutUInt31] for additional details.
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

// Encode an int value into an encoded Int31 type.
//
// The purpose of this encoding is to record numbers with 3 (and the half step)
// significant figures in the least number of bits.  Often this is used for
// storage of numbers which do not need unit level precision when the numbers
// get larger thus enabling values which would otherwise take 8 bytes of
// storage in 2 bytes.
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
