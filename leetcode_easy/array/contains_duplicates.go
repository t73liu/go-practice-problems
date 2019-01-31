package array

func containsDuplicate(nums []int) bool {
	exists := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		v, ok := exists[nums[i]]
		if ok {
			return v
		} else {
			exists[nums[i]] = true
		}
	}
	return false
}
