package main

import (
	"testing"
)

func Test01(t *testing.T) {
	result := longestPalindrome("babab")

	if result != "babab" {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	result := longestPalindrome("cbbd")

	if result != "bb" {
		t.Fail()
	}
}

func Test04(t *testing.T) {
	result := longestPalindrome("a")

	if result != "a" {
		t.Fail()
	}
}

func Test03(t *testing.T) {
	result := longestPalindrome("ac")

	if result != "a" {
		t.Fail()
	}
}
