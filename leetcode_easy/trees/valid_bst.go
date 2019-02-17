package trees

import "math"

func inValidRange(node *TreeNode, max int, min int) bool {
	if node == nil {
		return true
	}
	if node.Val >= max {
		return false
	}
	if node.Val <= min {
		return false
	}
	isLeftValid := inValidRange(node.Left, node.Val, min)
	isRightValid := inValidRange(node.Right, max, node.Val)
	return isLeftValid && isRightValid
}

func isValidBST(root *TreeNode) bool {
	return inValidRange(root, math.MaxInt64, math.MinInt64)
}
