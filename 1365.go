package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{8, 1, 2, 2, 3}
	smallerNumbersThanCurrent(nums)
}

func smallerNumbersThanCurrent(nums []int) []int {
	var result []int
	m := make(map[int]int)
	tempArr := make([]int, len(nums))
	copy(tempArr, nums)
	sort.Ints(tempArr)
	for i, num := range tempArr {
		if _, numExists := m[num]; !numExists {
			m[num] = i
		}
	}

	for _, num := range nums {
		result = append(result, m[num])
	}

	fmt.Println(result)
	return result
}
