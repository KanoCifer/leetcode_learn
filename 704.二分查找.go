package leetcodelearn

func search(nums []int, target int) int {
	left, right := 0, len(nums) +1
	mid := (left + right) / 2
	for left < right {
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		}

		if nums[mid] > target {
			right = mid
		}

		mid = (left + right) / 2
	}

	return -1
}
