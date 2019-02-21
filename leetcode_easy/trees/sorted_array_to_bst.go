package trees

func bstBuilder(nums []int, lower int, upper int) *TreeNode {
	size := len(nums)
	if lower > size || upper > size || upper < 0 || lower < 0 || lower > upper {
		return nil
	}
	if lower == upper {
		return &TreeNode{nums[lower], nil, nil}
	}
	mid := (lower + upper) / 2
	return &TreeNode{
		nums[mid],
		bstBuilder(nums, lower, mid-1),
		bstBuilder(nums, mid+1, upper),
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	return bstBuilder(nums, 0, len(nums)-1)
}
