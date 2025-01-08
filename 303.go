package main

type NumArray struct {
	numbers []int
}

func Constructor(nums []int) NumArray {
	//return NumArray{nums}
	newArray := make([]int, len(nums))

}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.numbers[i]
	}
	return sum
}

func main() {

}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
