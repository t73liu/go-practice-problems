package leetcode_easy

func rotate(nums []int, k int) {
	size := len(nums)
	times := k % size

	for i := 0; i < times; i++ {
		next := nums[0]
		for j := 1; j < size; j++ {
			temp := nums[j]
			nums[j] = next
			next = temp
		}
		nums[0] = next
	}
}
