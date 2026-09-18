package leetcodelearn

import (
	"strings"
	"unicode"
)

func calculate1(s string) int {
    s = strings.ReplaceAll(s, " ", "")
    pre := byte('+')
    num := 0
    stack := []int{0}
    apply := func() {
        switch pre {
        case '+':
            stack = append(stack, num)
        case '-':
            stack = append(stack, -num)
        case '*':
            top := stack[len(stack)-1]
            stack[len(stack)-1] = top * num
        case '/':
            top := stack[len(stack)-1]
            stack[len(stack)-1] = top / num
        }
    }

    for i := 0; i < len(s); i++ {
        ch := s[i]
        if ch >= '0' && ch <= '9' {
            num = int(ch-'0') + num*10
            if i != len(s)-1 {
                continue
            }
        }
        apply()
        pre = ch
        num = 0
    }

    result := 0
    for _, v := range stack {
        result += v
    }
    return result
}

func calculate(s string) int {
    s = strings.ReplaceAll(s, " ", "") // 中间空格也去掉
    parsed := []rune(s)
    pre := '+'
    num := 0
    stack := []int{0}

    apply := func() {
        switch pre {
        case '+':
            stack = append(stack, num)
        case '-':
            stack = append(stack, -num)
        case '*':
            top := stack[len(stack)-1]
            stack[len(stack)-1] = top * num
        case '/':
            top := stack[len(stack)-1]
            stack[len(stack)-1] = top / num
        }
    }

    for i, char := range parsed {
        if unicode.IsDigit(char) {
            num = int(char-'0') + num*10
        }
        if !unicode.IsDigit(char) || i == len(parsed)-1 {
            apply()
            pre = char
            num = 0
        }
    }

    result := 0
    for _, v := range stack {
        result += v
    }
    return result
}
