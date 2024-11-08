package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_rand10(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	counts := make(map[int]int)
	for i := 0; i < 1000000; i++ {
		num := rand10()
		counts[num]++
	}
	fmt.Println("Distribution of Rand10():", counts)
}
