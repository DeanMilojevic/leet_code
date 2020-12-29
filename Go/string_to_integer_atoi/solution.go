package main

import (
	"math"
)

func main() {}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	whitespace := byte(' ')
	minus := byte('-')
	plus := byte('+')
	zero := byte('0')
	nine := byte('9')

	result := 0
	sign := 1

	index := 0

	for index < len(s) && s[index] == whitespace {
		index++
	}

	if index < len(s) && (s[index] == plus || s[index] == minus) {
		if s[index] == minus {
			sign = -1
			index++
		} else if s[index] == plus {
			sign = 1
			index++
		}
	}

	for index < len(s) && s[index] >= zero && s[index] <= nine {
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && s[index]-'0' > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			}

			if sign == -1 {
				return math.MinInt32
			}
		}

		result = result*10 + int(s[index]-'0')
		index++
	}

	result = result * sign

	return result
}
