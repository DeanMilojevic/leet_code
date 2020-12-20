package main

func main() {}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var longest = 0
	m := map[string]int{}

	var i = 0
	for j := 0; j < len(s); j++ {
		char := s[j : j+1]
		val, found := m[char]

		if found {
			if val > i {
				i = val
			}
		}

		diff := j - i + 1
		if diff > longest {
			longest = diff
		}

		m[char] = j + 1
	}

	return longest
}
