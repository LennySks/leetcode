package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))

}

func twoSum(nums []int, target int) []int {
	newMap := make(map[int]int)
	for i, num := range nums {
		result := target - nums[i]

		if index, found := newMap[result]; found {
			return []int{index, i}
		} else {
			newMap[num] = i
		}
	}
	return nil
}
