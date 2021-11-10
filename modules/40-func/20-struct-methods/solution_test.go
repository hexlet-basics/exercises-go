package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseList(t *testing.T) {
	a := assert.New(t)

	var l *ListNode

	a.Nil(l.Reverse())

	l = &ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: &ListNode{
					Next: nil,
					Val:  40,
				},
				Val: 30,
			},
			Val: 20,
		},
		Val: 10,
	}

	rev := l.Reverse()

	// initial list should be the same after the reversion
	a.Equal(&ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: &ListNode{
					Next: nil,
					Val:  40,
				},
				Val: 30,
			},
			Val: 20,
		},
		Val: 10,
	}, l)

	a.Equal(&ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: &ListNode{
					Next: nil,
					Val:  10,
				},
				Val: 20,
			},
			Val: 30,
		},
		Val: 40,
	}, rev)
}
