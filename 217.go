package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, hasNum := set[num]; hasNum {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
