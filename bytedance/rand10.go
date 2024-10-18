package bytedance

import (
	"math/rand"
)

func rand7() int {
	return rand.Intn(7) + 1
}

func rand10() int {
	for {
		sum := (rand7()-1)*7 + rand7()

		if sum <= 10 {
			return sum%10 + 1
		}
	}
}
