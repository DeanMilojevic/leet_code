package main

import (
	"strings"
)

func main() {}

// Manacher’s Algorithm
func longestPalindrome(s string) string {

	t := transform(s)

	p := make([]int, len(t))

	id, m := 0, 0

	maxLen, maxLenCenter := 0, 0

	for i := range p {

		if i < m {
			p[i] = min(p[2*id-i], m-i)
		} else {
			p[i] = 1
		}

		for i-p[i] >= 0 && i+p[i] < len(t) && t[i-p[i]] == t[i+p[i]] {
			p[i]++
		}

		if p[i]+i-1 > m {
			m = p[i] + i - 1
			id = i
		}

		if maxLen < p[i]-1 {
			maxLen = p[i] - 1
			maxLenCenter = i
		}
	}

	return transformBack(t[maxLenCenter-maxLen : maxLenCenter+maxLen])
}

func transform(s string) string {
	var sb strings.Builder
	sb.WriteRune('#')
	for _, c := range s {
		sb.WriteRune(c)
		sb.WriteRune('#')
	}

	return sb.String()
}

func transformBack(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if c != '#' {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
