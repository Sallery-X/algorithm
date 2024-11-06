package bytedance

func calculate(s string) int {
	numStack := make([]int, 0)
	num := 0

	sign := '+'
	n := len(s)

	for i := 0; i < n; i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		}

		if ((ch < '0' || ch > '9') && ch != ' ') || i == n-1 {
			if sign == '+' {
				numStack = append(numStack, num)
			} else if sign == '-' {
				numStack = append(numStack, -num)
			} else if sign == '*' {
				numStack[len(numStack)-1] = num * numStack[len(numStack)-1]
			} else if sign == '/' {
				numStack[len(numStack)-1] = numStack[len(numStack)-1] / num
			}
			sign = int32(ch)
			num = 0

		}

	}

	res := 0

	for _, v := range numStack {
		res = res + v
	}
	return res

}
