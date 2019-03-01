package math

func countPrimes(n int) int {
	primeArr := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if primeArr[i] == false {
			count++
			for j := 2; i*j < n; j++ {
				primeArr[i*j] = true
			}
		}
	}
	return count
}
