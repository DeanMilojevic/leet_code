package main

func main() {}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i = 0
	var j = 0
	var k = 0

	length := len(nums1) + len(nums2)
	var merged = make([]int, length)

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged[k] = nums1[i]
			k = k + 1
			i = i + 1
		} else {
			merged[k] = nums2[j]
			k = k + 1
			j = j + 1
		}
	}

	for i < len(nums1) {
		merged[k] = nums1[i]
		k = k + 1
		i = i + 1
	}

	for j < len(nums2) {
		merged[k] = nums2[j]
		k = k + 1
		j = j + 1
	}

	l := len(merged)

	if l == 0 {
		return 0
	}

	if l%2 > 0 {
		return float64(merged[l/2])
	}

	first := float64(merged[(l/2 - 1)])
	second := float64(merged[l/2])

	return float64((first + second) / 2)
}
