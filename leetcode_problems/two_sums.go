package main

import (
	"fmt"
)

// Here I will compile this task
// https://leetcode.com/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	for element := 0; element < len(nums); element++ {
		for counter := element + 1; counter < len(nums); counter++ {
			if nums[element]+nums[counter] == target {
				return []int{element, counter}
			}
		}
	}
	return nil
}

func main() {
	result := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(result)
}
