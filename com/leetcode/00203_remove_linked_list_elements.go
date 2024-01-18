package main

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dh := new(ListNode)
	dh.Next = head
	c := dh
	for c.Next != nil {
		if c.Next.Val == val {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}
	return dh.Next
}
