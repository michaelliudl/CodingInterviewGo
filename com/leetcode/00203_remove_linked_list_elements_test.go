package main

import "testing"

func TestRemoveElements(t *testing.T) {
	t.Logf("Test %+v", removeElements(&ListNode{}, 1))
	t.Logf("Test %+v", removeElements(&ListNode{Val: 7,
		Next: &ListNode{Val: 7,
			Next: &ListNode{Val: 7,
				Next: &ListNode{Val: 7,
					Next: nil}}}}, 7))
	t.Logf("Test %+v", removeElements(&ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 6,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 0,
							Next: &ListNode{Val: 6,
								Next: nil}}}}}}}, 6))
}
