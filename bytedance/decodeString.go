package bytedance

import (
	"strings"
)

// input := "3[a2[c]]"
func decodeString(s string) string {
	var stack []string
	var numStack []int
	var current string
	var count int

	for _, c := range s {
		if c >= '0' && c <= '9' {
			count = count*10 + int(c-'0')
		} else if c == '[' {

			stack = append(stack, current)
			numStack = append(numStack, count)
			//fmt.Println(stack)
			//fmt.Println(numStack)
			current = ""
			count = 0
		} else if c == ']' {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			repeatN := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			repeated := strings.Repeat(current, repeatN)
			current = prev + repeated
		} else {
			current += string(c)
		}
	}

	return current
}
