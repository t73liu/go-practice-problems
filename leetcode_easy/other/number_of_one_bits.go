package other

import "strconv"

func hammingWeight(num uint32) int {
	binary := strconv.FormatInt(int64(num), 2)
	count := 0
	for i := 0; i < len(binary); i++ {
		if binary[i] == 49 {
			count++
		}
	}
	return count
}
