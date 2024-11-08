package bytedance

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	x := Tree{
		left: &Tree{
			left: &Tree{
				left:  nil,
				right: nil,
				value: 2,
			},
			right: &Tree{
				left:  nil,
				right: nil,
				value: 0,
			},
			value: 1,
		},
		right: &Tree{
			left: &Tree{
				left:  nil,
				right: nil,
				value: 0,
			},
			right: &Tree{
				left:  nil,
				right: nil,
				value: 1,
			},
			value: 1,
		},
		value: 0,
	}
	fmt.Println(IsSymmetric(&x))

	y := make(map[string]string, 0)
	y["o"] = "l"

	k, v := y["o"]
	fmt.Println(k)
	fmt.Println(v)

}
