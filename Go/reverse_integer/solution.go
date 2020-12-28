package main

import "math"

func main() {}

func reverse(x int) int {
	var result = 0

	for x != 0 {
		var reminder = x % 10
		result = result*10 + reminder

		if result > math.MaxInt32 {
			return 0
		}

		if result < math.MinInt32 {
			return 0
		}

		x = x / 10
	}

	return result
}
