package other

import "strconv"

func hammingDistance(x int, y int) int {
	xor := strconv.FormatInt(int64(x^y), 2)
	distance := 0
	for i := 0; i < len(xor); i++ {
		if xor[i] == '1' {
			distance++
		}
	}
	return distance
}
