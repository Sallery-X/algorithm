package array

import (
	"fmt"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	p := []int{7, 1, 5, 3, 6, 4}
	m := maxProfit(p)
	fmt.Println(m)
}
