package leetcodelearn

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	for i, l := 0, len(s)-1; i < l; i, l = i+1, l-1 {
		if s[i] != s[l] {
			return false
		}
	}

	return true
}
