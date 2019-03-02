package other

func missingNumber(nums []int) int {
	size := len(nums)
	sum := 0
	for i := 0; i < size; i++ {
		sum += nums[i]
	}
	return size*(size+1)/2 - sum
}
