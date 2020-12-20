package main

func main() {}

// ListNode needs a comment
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = ListNode{0, nil}
	var first = l1
	var second = l2
	var curr = &dummy

	var carry = 0

	for first != nil || second != nil {
		var x = 0
		if first != nil {
			x = first.Val
		}

		var y = 0
		if second != nil {
			y = second.Val
		}

		var sum = carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next

		if first != nil {
			first = first.Next
		}

		if second != nil {
			second = second.Next
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}

	return dummy.Next
}
