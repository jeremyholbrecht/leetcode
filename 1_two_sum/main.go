package main

import "fmt"

// brute force
/*
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
*/

// map

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// key = number, value = index of number
		m[nums[i]] = i
	}

	// check if key is present in map
	val, ok := m[7]
	if ok {
		fmt.Println(val)
	}

	return []int{}
}
func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}
