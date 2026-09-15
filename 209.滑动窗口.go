package leetcodelearn

func minSubArrayLen(target int, nums []int) int {
	slow := 0
	fast := 0
	sum := 0
	min_length := len(nums) + 1

	for fast < len(nums) {
		sum += nums[fast]
		for sum >= target {
			current_length := fast - slow + 1
			if current_length < min_length {
				min_length = current_length
			}
			sum -= nums[slow]
			slow++
		}
		fast++
	}

	if min_length == len(nums)+1 {
		return 0
	}
	return min_length
}

