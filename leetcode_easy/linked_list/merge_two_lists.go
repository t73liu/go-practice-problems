package linked_list

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var newHead *ListNode
	var newTail *ListNode
	for {
		if l1 == nil && l2 == nil {
			break
		} else if l1 == nil {
			newTail.Next = l2
			break
		} else if l2 == nil {
			newTail.Next = l1
			break
		}
		if newHead == nil {
			if l1.Val <= l2.Val {
				newHead = l1
				newTail = l1
				l1 = l1.Next
			} else {
				newHead = l2
				newTail = l2
				l2 = l2.Next
			}
		} else {
			if l1.Val <= l2.Val {
				newTail.Next = l1
				l1 = l1.Next
			} else {
				newTail.Next = l2
				l2 = l2.Next
			}
			newTail = newTail.Next
		}
	}
	return newHead
}
