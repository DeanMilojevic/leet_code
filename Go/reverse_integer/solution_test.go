package main

import "testing"

func Test01(t *testing.T) {
	result := reverse(123)

	if result != 321 {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	result := reverse(-123)

	if result != -321 {
		t.Fail()
	}
}

func Test03(t *testing.T) {
	result := reverse(120)

	if result != 21 {
		t.Fail()
	}
}

func Test04(t *testing.T) {
	result := reverse(0)

	if result != 0 {
		t.Fail()
	}
}
