package leetcodelearn

func s1(nums []int) int {
	var set = make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	for i := 1; i <= len(nums); i++ {
		if !set[i] {
			return i
		}
	}
	return len(nums) + 1
}

func firstMissingPositive(nums []int) int {
	i := 0

	for i < len(nums) {
		num := nums[i]
		n := len(nums)
		if num <= n && num >= 1 && nums[num-1] != nums[i] {
			nums[num-1], nums[i] = nums[i], nums[num-1]
		} else {
			i++
		}

	}

	for i, v := range nums {
		if i != v-1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
