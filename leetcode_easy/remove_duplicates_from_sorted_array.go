package leetcode_easy

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	curr := nums[0]
	currIndex := 1
	i := 1
	for i < length {
		if curr == nums[i] {
			i++
		} else {
			nums[currIndex] = nums[i]
			curr = nums[i]
			currIndex++
			i++
		}
	}
	return currIndex
}
