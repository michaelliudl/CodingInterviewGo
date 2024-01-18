package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var rHead, rCur, p1, p2 *ListNode
	p1 = l1
	p2 = l2
	var carry int
	for p1 != nil || p2 != nil {
		var c *ListNode = new(ListNode)
		var v int
		if p1 == nil {
			v = p2.Val + carry
		} else if p2 == nil {
			v = p1.Val + carry
		} else {
			v = p1.Val + p2.Val + carry
		}
		c.Val = v % 10
		carry = v / 10
		if rHead == nil {
			rHead = c
			rCur = c
		} else {
			rCur.Next = c
			rCur = c
		}
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
	if carry > 0 {
		var c *ListNode = new(ListNode)
		c.Val = carry
		rCur.Next = c
	}
	return rHead
}
