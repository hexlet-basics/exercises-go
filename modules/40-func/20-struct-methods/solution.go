package solution

// ListNode is a node of a linked list.
type ListNode struct {
	Next *ListNode
	Val  int
}

// BEGIN

// ReverseList reverses given linked list
func ReverseList(head *ListNode) *ListNode {
	var r *ListNode

	curr := head
	for curr != nil {
		next := curr.Next

		curr.Next = r
		r = curr

		curr = next
	}

	return r
}

// END
