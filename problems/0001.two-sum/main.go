package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9))
}
