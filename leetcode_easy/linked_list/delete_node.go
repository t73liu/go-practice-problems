package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Write a function to delete a node (except the tail) in a singly linked list,
// given only access to that node.
func deleteNode(node *ListNode) {
	current := node
	next := current.Next
	for {
		if next.Next == nil {
			current.Val = next.Val
			current.Next = nil
			break
		} else {
			current.Val = next.Val
			current = next
			next = next.Next
		}
	}
}
