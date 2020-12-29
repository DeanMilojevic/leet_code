package main

import "testing"

func Test01(t *testing.T) {
	result := isPalindrome(121)

	if !result {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	result := isPalindrome(-121)

	if result {
		t.Fail()
	}
}

func Test03(t *testing.T) {
	result := isPalindrome(10)

	if result {
		t.Fail()
	}
}

func Test04(t *testing.T) {
	result := isPalindrome(-101)

	if result {
		t.Fail()
	}
}
