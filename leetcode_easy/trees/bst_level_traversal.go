package trees

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := make([]*TreeNode, 0)
	resultLevel := make([]int, 0)
	for {
		queueSize := len(queue)
		if queueSize == 0 {
			break
		} else {
			for i := 0; i < queueSize; i++ {
				resultLevel = append(resultLevel, queue[i].Val)
				if queue[i].Left != nil {
					level = append(level, queue[i].Left)
				}
				if queue[i].Right != nil {
					level = append(level, queue[i].Right)
				}
			}
			queue = level
			result = append(result, resultLevel)
			level = nil
			resultLevel = nil
		}
	}

	return result
}
