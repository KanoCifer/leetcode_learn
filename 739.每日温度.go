package leetcodelearn

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := []int{}
	for i, cur := range temperatures {
		for len(stack) > 0 && cur > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return ans
}
