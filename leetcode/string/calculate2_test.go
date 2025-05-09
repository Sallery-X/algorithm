package string

import "testing"
import "fmt"

func TestCalculateExpression(t *testing.T) {
	expressions := []string{
		"ADD(1,SUB(1, 2))",                  // (1-2)+(3+4) = -1+7 = 6
		"SUB(ADD(5, 7), DIV(6, 2))",         // (5+7)-(6/2) = 12-3 = 9
		"MUL(ADD(1, 2), SUB(3, 4))",         // (1+2)*(3-4) = 3*-1 = -3
		"ADD(1, SUB(2, ADD(3, DIV(8, 4))))", // 1+(2-(3+(8/4))) = 1+(2-(3+2)) = 1+(2-5) = 1+(-3) = -2
		"5",                                 // 直接数字
	}

	for _, expr := range expressions {
		result, err := CalculateExpression(expr)
		if err != nil {
			fmt.Printf("Error calculating '%s': %v\n", expr, err)
			continue
		}
		fmt.Printf("%s = %d\n", expr, result)
	}
}
