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
