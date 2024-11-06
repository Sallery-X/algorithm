package bytedance

import (
	"fmt"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := dailyTemperatures(temperatures)
	fmt.Println("Days to wait for a warmer temperature:", result)
}
