package main

import "testing"

func Test01(t *testing.T) {
	result := convert("PAYPALISHIRING", 3)

	if result != "PAHNAPLSIIGYIR" {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	result := convert("A", 1)

	if result != "A" {
		t.Fail()
	}
}
