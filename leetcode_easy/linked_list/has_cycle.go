package linked_list

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	curr := head
	fast := head.Next
	for {
		if curr == fast {
			return true
		} else {
			if curr == nil || fast == nil {
				break
			} else {
				curr = curr.Next
				if fast.Next == nil {
					break
				}
				fast = fast.Next.Next
			}
		}
	}
	return false
}
