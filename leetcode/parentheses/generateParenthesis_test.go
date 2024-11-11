package parentheses

import (
	"fmt"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	n := 3
	combinations := generateParenthesis(n)
	fmt.Println("All combinations of valid parentheses for n =", n)
	for _, combination := range combinations {
		fmt.Println(combination)
	}
}
