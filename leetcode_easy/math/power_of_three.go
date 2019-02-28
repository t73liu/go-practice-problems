package math

// 1162261467 = 3^19 within int limits
func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}
