package main

import "testing"

func Test01(t *testing.T) {
	result := findMedianSortedArrays([]int{1, 3}, []int{2})

	if result != 2 {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	result := findMedianSortedArrays([]int{1}, []int{})

	if result != 1 {
		t.Fail()
	}
}

func Test03(t *testing.T) {
	result := findMedianSortedArrays([]int{}, []int{1})

	if result != 1 {
		t.Fail()
	}
}

func Test05(t *testing.T) {
	result := findMedianSortedArrays([]int{1, 2}, []int{3, 4})

	if result != 2.5 {
		t.Fail()
	}
}

func Test06(t *testing.T) {
	result := findMedianSortedArrays([]int{}, []int{})

	if result != 0 {
		t.Fail()
	}
}

func Test07(t *testing.T) {
	result := findMedianSortedArrays([]int{}, []int{1, 2, 3, 4, 5})

	if result != 3 {
		t.Fail()
	}
}
