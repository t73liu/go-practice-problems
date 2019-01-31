package array

func maxProfit(prices []int) int {
	size := len(prices)

	if size < 2 {
		return 0
	}

	profit := 0
	low := prices[0]
	for i := 1; i < size; i++ {
		if prices[i] < low {
			low = prices[i]
		} else {
			profit += prices[i] - low
			low = prices[i]
		}
	}

	return profit
}
