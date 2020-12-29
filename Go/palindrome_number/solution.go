package main

func main() {}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverse := 0
	temp := int(x)

	for temp > 0 {
		r := temp % 10
		reverse = reverse*10 + r
		temp = temp / 10
	}

	return x == reverse
}
