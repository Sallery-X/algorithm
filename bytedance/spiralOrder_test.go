package bytedance

import (
	"fmt"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	x := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(x))
}
