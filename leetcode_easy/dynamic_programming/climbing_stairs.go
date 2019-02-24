package dynamic_programming

func climbStairs(n int) int {
	count := make([]int, n+1)
	count[0] = 1
	count[1] = 1
	for i := 2; i <= n; i++ {
		count[i] = count[i-1] + count[i-2]
	}
	return count[n]
}
