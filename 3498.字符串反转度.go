package leetcodelearn

func reverseDegree(s string) int {
	result := 0
	str := []rune(s)
	n := len(str)


	for i := 0; i < n ;i++ {
		num := int('z' - str[i] + 1 )
		result += num* (i+1)
	}

	return result
}
