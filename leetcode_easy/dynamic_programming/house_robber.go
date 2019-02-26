package dynamic_programming

func rob(nums []int) int {
	maxOne := 0
	maxTwo := 0
	for i := 0; i < len(nums); i++ {
		temp := maxOne
		if maxOne < maxTwo+nums[i] {
			maxOne = maxTwo + nums[i]
		}
		maxTwo = temp
	}
	return maxOne
}
