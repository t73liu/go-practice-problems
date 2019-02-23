package dynamic_programming

func maxProfit(prices []int) int {
	profit := 0
	if len(prices) < 2 {
		return profit
	}
	min := prices[0]
	maxSoFar := prices[0]
	for i := 1; i < len(prices); i++ {
		curr := prices[i]
		if curr < min {
			min = curr
			maxSoFar = curr
		}
		if curr > maxSoFar {
			maxSoFar = curr
			if maxSoFar-min > profit {
				profit = maxSoFar - min
			}
		}
	}

	return profit
}
