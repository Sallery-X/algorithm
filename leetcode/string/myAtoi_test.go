package string

import (
	"fmt"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	fmt.Println(myAtoi("   -42"))          // 输出: -42
	fmt.Println(myAtoi("4193 with words")) // 输出: 4193
	fmt.Println(myAtoi("words and 987"))   // 输出: 0
	fmt.Println(myAtoi("-91283472332"))
}
