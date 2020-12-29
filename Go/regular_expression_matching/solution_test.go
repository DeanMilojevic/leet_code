package main

import "testing"

func Test01(t *testing.T) {
	result := isMatch("aa", "a")

	if result {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	result := isMatch("aa", "a*")

	if !result {
		t.Fail()
	}
}

func Test03(t *testing.T) {
	result := isMatch("ab", ".*")

	if !result {
		t.Fail()
	}
}

func Test04(t *testing.T) {
	result := isMatch("aab", "c*a*b")

	if !result {
		t.Fail()
	}
}

func Test05(t *testing.T) {
	result := isMatch("mississippi", "mississippi")

	if result {
		t.Fail()
	}
}
