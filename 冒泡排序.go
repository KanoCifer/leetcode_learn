package leetcodelearn

func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := range n - 1 {
		swap := false
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swap = true
			}
		}

		if !swap {
			return nums
		}
	}

	return nums
}
