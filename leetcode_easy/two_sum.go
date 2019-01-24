package leetcode_easy

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	seen := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		val, ok := seen[target-nums[i]]
		if ok {
			result[0] = val
			result[1] = i
			break
		} else {
			seen[nums[i]] = i
		}
	}

	return result
}
