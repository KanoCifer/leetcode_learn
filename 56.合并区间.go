package leetcodelearn

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	cur := []int{}
	next := []int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, interval := range intervals {
		next = interval
		if len(cur) == 0 {
			cur = next
		}

		if cur[1] >= next[0] {
			cur[1] = int(math.Max(float64(cur[1]), float64(next[1])))
		} else {
			ans = append(ans, cur)
			cur = next
		}
	}
	ans = append(ans, cur)
	return ans
}
