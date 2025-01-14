package main

func main() {
	arr := [][]int{{2, 5, 8}, {4, 0, -1}}
	spiralOrder(arr)
}

func spiralOrder(matrix [][]int) []int {
	var res []int
	// Check if the matrix is empty, if it is return res
	for len(matrix) > 0 {
		// Take first array and remove the first element of the matrix
		res = append(res, matrix[0]...)
		matrix = matrix[1:]
		// Take last elements of next arrays
		for i := 0; i < len(matrix); i++ {
			if len(matrix[i]) > 0 {
				res = append(res, matrix[i][len(matrix[i])-1])
				matrix[i] = matrix[i][:len(matrix[i])-1]
			}
		}
		// Take all  elements from last array
		if len(matrix) > 0 {
			for j := len(matrix[len(matrix)-1]) - 1; j >= 0; j-- {
				res = append(res, matrix[len(matrix)-1][j])
			}
			matrix = matrix[:len(matrix)-1]
		}

		// Take first element of next arrays in reverse
		for i := len(matrix) - 1; i > 0; i-- {
			if len(matrix[i]) > 0 {
				res = append(res, matrix[i][0])
				matrix[i] = matrix[i][1:]
			}
		}
	}

	return res
}
