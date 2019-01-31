package array

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	for zeroIndex := 0; zeroIndex < len(nums); zeroIndex++ {
		if nums[zeroIndex] == 0 {
			nonZeroIndex := zeroIndex + 1
			for ; nonZeroIndex < len(nums); nonZeroIndex++ {
				if nums[nonZeroIndex] != 0 {
					break
				}
			}
			if nonZeroIndex >= len(nums) {
				break
			} else {
				nums[zeroIndex] = nums[nonZeroIndex]
				nums[nonZeroIndex] = 0
			}
		}
	}
}
