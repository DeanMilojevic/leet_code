package main

import "strings"

func main() {}

func convert(s string, numRows int) string {
	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	if numRows == 1 {
		return s
	}

	var arr = make([]string, numRows)
	row := 0
	down := false

	for i := 0; i < len(s); i++ {
		arr[row] = arr[row] + s[i:i+1]

		if row == numRows-1 {
			down = false
		}

		if row == 0 {
			down = true
		}

		if down {
			row++
		} else {
			row--
		}
	}

	var sb strings.Builder
	for i := 0; i < len(arr); i++ {
		sb.WriteString(arr[i])
	}

	return sb.String()
}
