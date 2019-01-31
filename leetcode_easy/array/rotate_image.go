package array

func rotate_matrix(matrix [][]int) {
	flip(matrix)
	transpose(matrix)
}

func flip(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length/2; i++ {
		for j := 0; j < length; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[length-1-i][j]
			matrix[length-1-i][j] = temp
		}
	}
}

func transpose(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
}
