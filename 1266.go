package main

import (
	"math"
)

func main() {
	test := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	minTimeToVisitAllPoints(test)
}

func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	for i := 0; i < len(points)-1; i++ {
		res += int(math.Max(math.Abs(float64(points[i+1][1]-points[i][1])), math.Abs(float64(points[i+1][0]-points[i][0]))))
	}
	return res
}
