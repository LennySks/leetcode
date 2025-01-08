package main

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	findDisappearedNumbers(nums)
}

func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	var missingNumbers []int
	for _, num := range nums {
		res[num-1] = num
	}

	for i := 0; i < length; i++ {
		if res[i] == 0 {
			missingNumbers = append(missingNumbers, i+1)
		}
	}
	return missingNumbers
}
