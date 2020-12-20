package main

import "testing"

func Test01(t *testing.T) {
	result := lengthOfLongestSubstring("abcabcbb")

	if result != 3 {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	result := lengthOfLongestSubstring("bbbbb")

	if result != 1 {
		t.Fail()
	}
}

func Test03(t *testing.T) {
	result := lengthOfLongestSubstring("pwwkew")

	if result != 3 {
		t.Fail()
	}
}
