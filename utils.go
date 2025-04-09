package epsolar

func checkBit(v uint16, bitNo int) bool {
	return ((v >> bitNo) & 1) == 1
}

func getBits(v uint16, bitNo int, mask uint16) uint16 {
	return (v >> bitNo) & mask
}
