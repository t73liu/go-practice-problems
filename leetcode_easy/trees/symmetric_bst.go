package trees

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	} else if left.Val != right.Val {
		return false
	}
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}
