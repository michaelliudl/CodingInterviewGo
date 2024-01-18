package main

func sortListNaive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for slow != nil {
		fast = slow.Next
		for fast != nil {
			if fast.Val < slow.Val {
				slow.Val, fast.Val = fast.Val, slow.Val
			}
			fast = fast.Next
		}
		slow = slow.Next
	}
	return head
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	result := mergeSort(head)
	return result
}

func mergeSort(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	half1, half2 := split(head)
	half1 = mergeSort(half1)
	half2 = mergeSort(half2)
	result := merge(half1, half2)
	return result
}

func split(head *ListNode) (*ListNode, *ListNode) {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	t := slow.Next
	slow.Next = nil
	return head, t
}

func merge(half1, half2 *ListNode) *ListNode {
	dummy := new(ListNode)
	c := dummy
	for half1 != nil && half2 != nil {
		if half1.Val < half2.Val {
			c.Next = half1
			half1 = half1.Next
		} else {
			c.Next = half2
			half2 = half2.Next
		}
		c = c.Next
	}
	if half1 == nil {
		c.Next = half2
	}
	if half2 == nil {
		c.Next = half1
	}
	return dummy.Next
}
