package linked_list

// O(n) time, O(1) space
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// Find mid index
	i := 0
	curr := head
	mid := head
	for {
		if curr == nil || curr.Next == nil {
			break
		}
		mid = mid.Next
		curr = curr.Next.Next
		i++
	}

	// Reverse list from mid to end
	reverse := reverseList(mid)

	// Compare reversed list against head
	curr = head
	for {
		if i > 0 {
			if reverse.Val != curr.Val {
				return false
			} else {
				reverse = reverse.Next
				curr = curr.Next
			}
			i--
		} else {
			break
		}
	}

	return true
}
