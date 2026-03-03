package parentheses

func isValid(s string) bool {
	var stack []rune

	sMap := make(map[rune]rune, 0)
	sMap['('] = ')'
	sMap['['] = ']'
	sMap['{'] = '}'

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			r := stack[len(stack)-1]
			if sMap[r] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
