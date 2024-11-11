package dp

import (
	"fmt"
	"testing"
)

func Test_coinChange(t *testing.T) {
	nums := []int{1, 2, 3, 3, 9}
	fmt.Println(coinChange(nums, 10))
}
