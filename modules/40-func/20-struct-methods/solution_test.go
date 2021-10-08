package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseList(t *testing.T) {
	a := assert.New(t)
	a.Nil(ReverseList(nil))
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
	}, ReverseList(&ListNode{
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
	}))
}
