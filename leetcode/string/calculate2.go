package string

import (
	"fmt"
	"strconv"
	"strings"
)

// 定义支持的运算符
var operators = map[string]func(int, int) int{
	"ADD": func(a, b int) int { return a + b },
	"SUB": func(a, b int) int { return a - b },
	"MUL": func(a, b int) int { return a * b },
	"DIV": func(a, b int) int { return a / b },
}

// 主计算函数
func CalculateExpression(s string) (int, error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0, fmt.Errorf("empty expression")
	}

	// 如果是数字直接返回
	if num, err := strconv.Atoi(s); err == nil {
		return num, nil
	}

	// 解析操作符和参数
	op, args, err := parseExpression(s)
	if err != nil {
		return 0, err
	}

	// 验证操作符是否支持
	operation, ok := operators[op]
	if !ok {
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}

	// 计算参数
	if len(args) != 2 {
		return 0, fmt.Errorf("operator %s expects 2 arguments, got %d", op, len(args))
	}

	arg1, err := CalculateExpression(args[0])
	if err != nil {
		return 0, err
	}

	arg2, err := CalculateExpression(args[1])
	if err != nil {
		return 0, err
	}

	return operation(arg1, arg2), nil
}

// 解析表达式为操作符和参数
func parseExpression(s string) (string, []string, error) {
	// 找到操作符和括号的位置
	opEnd := strings.Index(s, "(")
	if opEnd == -1 {
		return "", nil, fmt.Errorf("invalid expression format: %s", s)
	}

	op := s[:opEnd]
	argsStr := s[opEnd+1 : len(s)-1] // 去掉外层括号

	// 分割参数
	args, err := splitArguments(argsStr)
	if err != nil {
		return "", nil, err
	}

	return op, args, nil
}

// 分割参数，处理嵌套表达式
func splitArguments(s string) ([]string, error) {
	var args []string
	parenDepth := 0
	start := 0

	for i, ch := range s {
		if ch == '(' {
			parenDepth++
		} else if ch == ')' {
			parenDepth--
		} else if ch == ',' && parenDepth == 0 {
			arg := s[start:i]
			args = append(args, strings.TrimSpace(arg))
			start = i + 1
		}
	}

	// 添加最后一个参数
	lastArg := s[start:]
	args = append(args, strings.TrimSpace(lastArg))

	if parenDepth != 0 {
		return nil, fmt.Errorf("unbalanced parentheses in: %s", s)
	}
	fmt.Println(args)

	return args, nil
}
