package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		// setting j to i + 1 makes it so that the same elemen't won't be used twice
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}

		}
	}
	return []int{}
}

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}
