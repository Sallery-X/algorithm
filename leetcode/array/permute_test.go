package array

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
	nums := []int{1, 2, 3}

	for _, p := range permute(nums) {
		fmt.Println(p)
	}

}
