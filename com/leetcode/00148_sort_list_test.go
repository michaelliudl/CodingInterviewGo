package main

import "testing"

func TestSortList(t *testing.T) {
	t.Logf("Test %+v", sortList(&ListNode{Val: 4,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 1,
				Next: &ListNode{Val: 3,
					Next: nil}}}}))
	t.Logf("Test %+v", sortList(&ListNode{Val: -1,
		Next: &ListNode{Val: 5,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 0,
						Next: nil}}}}}))
	// t.Logf("Test %v", addTwoNumbers(nil, nil))
	// t.Logf("Test %v", addTwoNumbers(nil, nil))
}
