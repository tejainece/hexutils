package hexutils

import (
//"fmt"
)

//MakeUInt16 makes uint16 from byte array
func MakeUInt16(aBytes [2]byte) uint16 {
	lRet := uint16(0)

	lRet = uint16(aBytes[0])
	lRet = lRet << 8
	lRet |= uint16(aBytes[1])

	return lRet
}

//MakeUInt32FromUInt16 makes uint32 from two uint16
func MakeUInt32FromUInt16(aHi, aLo uint16) uint32 {
	lRet := uint32(0)

	lRet = uint32(aHi)
	lRet = lRet << 16
	lRet |= uint32(aLo)

	return lRet
}
