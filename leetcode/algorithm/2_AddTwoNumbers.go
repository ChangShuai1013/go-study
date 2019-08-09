package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val: 0,
	}
	curr := dummyHead
	carry := 0
	for l1 != nil || l2 != nil {
		var x, y int = 0, 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{
			Val: sum % 10,
		}
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{
			Val: carry,
		}
	}
	return dummyHead.Next
}
