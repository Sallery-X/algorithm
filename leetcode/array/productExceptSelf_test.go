package array

import (
	"fmt"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	testCases := [][]int{
		{1, 2, 3, 4},
		{-1, 1, 0, -3, 3},
	}

	for _, tc := range testCases {
		result := productExceptSelf(tc)
		fmt.Printf("输入：%v\n输出：%v\n\n", tc, result)
	}
}
