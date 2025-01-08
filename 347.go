package main

import (
	"sort"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	topKFrequent(nums, 2)
}

type freqPair struct {
	num  int
	freq int
}

func topKFrequent(nums []int, k int) []int {
	res := make(map[int]int)
	var kReturn []int
	for _, n := range nums {
		res[n]++
	}
	var freqSlice []freqPair
	for num, freq := range res {
		freqSlice = append(freqSlice, freqPair{num, freq})
	}

	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i].freq > freqSlice[j].freq
	})

	for i := 0; i < k && i < len(freqSlice); i++ {
		kReturn = append(kReturn, freqSlice[i].num)
	}

	return kReturn
}
