package bytedance

import (
	"fmt"
	"testing"
)

func Test_calculate(t *testing.T) {
	expression := "3+2*2"
	fmt.Println("Result:", calculate(expression)) // Output: 7

	expression = " 3/2 "
	fmt.Println("Result:", calculate(expression)) // Output: 1

	expression = " 3+5 / 2 "
	fmt.Println("Result:", calculate(expression)) // Output: 5
}
