package string

func calculate(s string) int {
	// numStack 用来保存已经处理过的数字
	//
	// 处理规则：
	// 1. 遇到 + ，当前数字直接入栈
	// 2. 遇到 - ，当前数字取负后入栈
	// 3. 遇到 * / ，直接和栈顶元素计算后再放回栈顶
	//
	// 这样就能保证：
	// - 加减最后统一求和
	// - 乘除在遍历过程中优先计算
	numStack := make([]int, 0)

	// num 表示当前正在解析的数字
	// 因为可能有多位数，所以不能每次只看一个字符
	// 例如 "123" 要解析成 123
	num := 0

	// sign 表示“当前这个 num 前面的运算符”
	// 初始默认是 '+'，相当于表达式最前面补了一个 + 号
	sign := byte('+')

	n := len(s)

	for i := 0; i < n; i++ {
		ch := s[i]

		// 如果当前字符是数字，就不断构造当前数字
		// 例如：
		// num = 1
		// 继续读到 '2' 后，num = 1*10 + 2 = 12
		// 再读到 '3' 后，num = 12*10 + 3 = 123
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		}

		// 当满足以下任一条件时，说明前面的数字 num 已经读完整了，需要结算：
		// 1. 当前字符不是数字，并且不是空格，说明遇到了运算符
		// 2. 当前已经到达字符串末尾，最后一个数字也必须处理
		if ((ch < '0' || ch > '9') && ch != ' ') || i == n-1 {

			// 根据“num 前面的运算符 sign”来处理这个数字
			switch sign {
			case '+':
				// 加法：直接把当前数字入栈
				numStack = append(numStack, num)

			case '-':
				// 减法：把当前数字取负后入栈
				numStack = append(numStack, -num)

			case '*':
				// 乘法：直接和栈顶元素计算
				// 例如原来栈顶是 2，当前 num 是 3
				// 那么把栈顶改成 2*3=6
				numStack[len(numStack)-1] = numStack[len(numStack)-1] * num

			case '/':
				// 除法：直接和栈顶元素计算
				// Go 中整数除法会自动向 0 截断
				// 比如：
				// 3/2  = 1
				// -3/2 = -1
				numStack[len(numStack)-1] = numStack[len(numStack)-1] / num
			}

			// 当前字符如果是运算符，那么它会成为下一个数字前面的运算符
			// 如果当前是最后一个数字，这里赋值虽然没实际作用，但也不影响
			sign = ch

			// 当前数字已经处理完了，重置为 0，准备解析下一个数字
			num = 0
		}
	}

	// 最后把栈里的所有数加起来
	// 因为乘除已经提前算过了，所以这里直接累加即可
	res := 0
	for _, v := range numStack {
		res += v
	}

	return res
}
