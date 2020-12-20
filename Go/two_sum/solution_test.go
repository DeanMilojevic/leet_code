package main

import "testing"

func Test01(t *testing.T) {
	result := twoSum([]int{1, 2, 3, 4, 5}, 5)

	if result[0] != 2 && result[1] != 3 {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	result := twoSum([]int{1, 2, 3, 4, 5}, 50)

	if len(result) > 0 {
		t.Fail()
	}
}
