package main

import (
	"fmt"
	"slices"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	nums := []int{1, 2, 2}
	k := 2
	minOperations(nums, k)
}

func minOperations(nums []int, k int) int {
	count := 0
	var newNums []int
	for i := len(nums) - 1; i >= 0; i-- {
		if len(newNums) <= k {
			if nums[i] <= k && !slices.Contains(newNums, nums[i]) {
				newNums = append(newNums, nums[i])
			}
			count++
		}
		if len(newNums) == k {
			break
		}
	}
	fmt.Println(count)
	return count
}
