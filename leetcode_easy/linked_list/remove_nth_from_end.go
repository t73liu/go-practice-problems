package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	offset := head
	for i := 0; i < n; i++ {
		offset = offset.Next
	}
	if offset == nil {
		return head.Next
	}
	current := head
	for {
		if offset.Next == nil {
			current.Next = current.Next.Next
			break
		} else {
			offset = offset.Next
			current = current.Next
		}
	}
	return head
}
