package leetcodelearn

import "fmt"

type NumArray struct {
	nums      []int
	prefixSum []int
}

func constructor(nums []int) NumArray {
	// 初始化前缀和数组
	prefixSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	return NumArray{nums: nums, prefixSum: prefixSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	obj := constructor(nums)
	param_1 := obj.SumRange(0, len(nums)-1)
	fmt.Println(param_1)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
