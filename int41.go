package float

// Decode the encoded Int41.  The Int41 type is an integer with 3 significant
// figures and an additional half.  See [PutUInt41] for additional details.
func UInt41(b []byte) uint {
	val := byteToUintKeep(b, 0xff)
	if val <= 10000 {
		return val
	}
	exp := uint(1)
	val = val - 10000
	for val > 18000 {
		exp, val = exp*10, val-18000
	}
	return (10000 + val*5) * exp
}

// Encode an int value into an encoded Int41 type.
//
// The purpose of this encoding is to record numbers with 3 (and the half step)
// significant figures in the least number of bits.  Often this is used for
// storage of numbers which do not need unit level precision when the numbers
// get larger thus enabling values which would otherwise take 8 bytes of
// storage in 2 bytes.
func PutUInt41(b []byte, val uint) {
	if val <= 10000 {
		putUintRight(b, val)
		return
	}
	if val < 100000 {
		// One digit rounding 2->0,3->5
		putUintRight(b, (val-9998)/5+10000)
		return
	}
	exp := uint(1)
	ret := uint(28000)
	for {
		if val < 1000000*exp {
			// Two digit rounding 24->00,25->50
			putUintRight(b, (val-99975*exp)/(50*exp)+ret)
			return
		}
		exp, ret = exp*10, ret+18000
	}
}
