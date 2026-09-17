package leetcodelearn

func getCommon(nums1 []int, nums2 []int) int {

	for _, num := range nums1 {
		if lowerBound(nums2, num) != -1 {
			return num
		}
	}

	return -1
}

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	mid := (left + right) / 2
	for left < right {
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
		mid = (left + right) / 2
	}
	return -1
}
