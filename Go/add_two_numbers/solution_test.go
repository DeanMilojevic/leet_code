package main

import (
	"strconv"
	"testing"
)

func Test01(t *testing.T) {
	first := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	second := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(first, second)

	var str = ""
	for result != nil {
		str = str + strconv.Itoa(result.Val)
		result = result.Next
	}

	if str != "708" {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	first := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	second := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	result := addTwoNumbers(first, second)

	var str = ""
	for result != nil {
		str = str + strconv.Itoa(result.Val)
		result = result.Next
	}

	if str != "89990001" {
		t.Fail()
	}
}
