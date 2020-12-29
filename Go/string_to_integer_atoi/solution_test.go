package main

import "testing"

func Test01(t *testing.T) {
	result := myAtoi("42")

	if result != 42 {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	result := myAtoi("   -42")

	if result != -42 {
		t.Fail()
	}
}

func Test03(t *testing.T) {
	result := myAtoi("4193 with words")

	if result != 4193 {
		t.Fail()
	}
}

func Test04(t *testing.T) {
	result := myAtoi("words and 987")

	if result != 0 {
		t.Fail()
	}
}

func Test05(t *testing.T) {
	result := myAtoi("-91283472332")

	if result != -2147483648 {
		t.Fail()
	}
}
