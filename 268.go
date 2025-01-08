package main

import "fmt"

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(len(nums))
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	var arrsum, allsum int
	for i := 0; i < len(nums); i++ {
		arrsum += nums[i] //37
		allsum += i
	}
	return allsum + len(nums) - arrsum
}
