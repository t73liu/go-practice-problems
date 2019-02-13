package linked_list

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	for current := head; current != nil; {
		temp := current.Next
		current.Next = newHead
		newHead = current
		current = temp
	}
	return newHead
}
