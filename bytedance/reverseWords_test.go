package bytedance

import (
	"fmt"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	s := "you are a   b"
	fmt.Println(reverseWords(s))
}
