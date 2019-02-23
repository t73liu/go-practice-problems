package dynamic_programming

func maxSubArray(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	max := nums[0]
	sum := max
	for i := 1; i < size; i++ {
		if sum+nums[i] > nums[i] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if max < sum {
			max = sum
		}
	}
	return max
}
