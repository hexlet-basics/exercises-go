package solution

// ListNode is a node of a linked list.
type ListNode struct {
	Next *ListNode
	Val  int
}

// BEGIN

// Reverse reverses the linked list
func (head *ListNode) Reverse() *ListNode {
	if head == nil {
		return nil
	}

	var r *ListNode

	curr := &(*head)
	for curr != nil {
		r = &ListNode{
			Next: r,
			Val:  curr.Val,
		}
		curr = curr.Next
	}

	return r
}

// END
