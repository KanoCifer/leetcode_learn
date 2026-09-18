package leetcodelearn

import (
	"slices"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	stack := []string{}
	numStack := []int{}

	num := 0

	var popstack func() string
	popstack = func() string {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return top
	}

	for _, char := range s {
		if unicode.IsDigit(char) {
			num = num*10 + int(char-'0')
		} else if char == '[' {
			numStack = append(numStack, num)
			stack = append(stack, string(char))
			num = 0
		} else if char == ']' {
			str := []string{}
			for stack[len(stack)-1] != "[" {
				str = append(str, popstack())
			}
			slices.Reverse(str)
			inner_str := strings.Join(str, "")
			_ = popstack()
			repeat := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			stack = append(stack, strings.Repeat(inner_str, repeat))
		} else {
			stack = append(stack, string(char))
		}
	}

	result := strings.Join(stack, "")
	return result
}
