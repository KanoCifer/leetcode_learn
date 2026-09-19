package leetcodelearn

func longestPalindrome(s string) string {
	runes := []rune(s)
	n := len(runes)
	if n < 2 {
		return s
	}

	// expand 返回以 (l, r) 为中心向外扩展，能得到的最长回文区间 [L, R]
	expand := func(l, r int) (int, int) {
		for l >= 0 && r < n && runes[l] == runes[r] {
			l--
			r++
		}
		// 循环退出时 l/r 已经"越界一步"，回文真正范围是 [l+1, r-1]
		return l + 1, r - 1
	}

	bestL, bestR := 0, 0
	for i := 0; i < n; i++ {
		L1, R1 := expand(i, i)
		L2, R2 := expand(i, i+1)

		if R1-L1 > bestR-bestL {
			bestL, bestR = L1, R1
		}
		if R2-L2 > bestR-bestL {
			bestL, bestR = L2, R2
		}
	}
	return string(runes[bestL : bestR+1])
}
