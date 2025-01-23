package epsolar

import "encoding/binary"

func checkBit(v uint16, bitNo int) bool {
	return ((v >> bitNo) & 1) == 1
}

func getBits(v uint16, bitNo int, mask uint16) uint16 {
	return (v >> bitNo) & mask
}

func convert16BitRegister(v uint16, divisor float64) float64 {
	return float64(v) / divisor
}

func convert32BitRegister(high uint16, low uint16, divisor float64) float64 {
	var b []byte
	b = binary.BigEndian.AppendUint16(b, high)
	b = binary.BigEndian.AppendUint16(b, low)
	return float64(int32(binary.BigEndian.Uint32(b))) / divisor
}
