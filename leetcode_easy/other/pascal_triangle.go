package other

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		size := i + 1
		result[i] = make([]int, size)
		result[i][0] = 1
		result[i][size-1] = 1
		for j := 1; j < size-1; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}
