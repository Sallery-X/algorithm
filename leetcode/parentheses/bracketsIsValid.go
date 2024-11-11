package parentheses

import "container/list"

func isValid(s string) bool {
	stack := list.New()

	sMap := make(map[rune]rune, 0)
	sMap['('] = ')'
	sMap['['] = ']'
	sMap['{'] = '}'

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack.PushFront(v)
		} else {
			if stack.Len() == 0 {
				return false
			}
			r := stack.Front().Value.(rune)
			if sMap[r] != v {
				return false
			}
			stack.Remove(stack.Front())
		}
	}
	return stack.Len() == 0
}
