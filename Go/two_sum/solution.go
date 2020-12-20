package main

func main() {}

func twoSum(nums []int, target int) []int {
	var m = map[int]int{}

	for i := 0; i < len(nums); i++ {
		var reminder = target - nums[i]

		if v, found := m[reminder]; found {
			return []int{i, v}
		}

		m[nums[i]] = i
	}

	return []int{}
}
