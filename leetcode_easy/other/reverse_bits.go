package other

import (
	"strconv"
	"strings"
)

func reverseBits(num uint32) uint32 {
	binary := strconv.FormatUint(uint64(num), 2)
	size := len(binary)
	var sb strings.Builder
	for i := size - 1; i >= 0; i-- {
		sb.WriteByte(binary[i])
	}
	for i := sb.Len(); i < 32; i++ {
		sb.WriteByte(48)
	}
	result, _ := strconv.ParseUint(sb.String(), 2, 32)
	return uint32(result)
}
