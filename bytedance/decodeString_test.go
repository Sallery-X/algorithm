package bytedance

import (
	"fmt"
	"testing"
)

func Test_decodeString(t *testing.T) {
	input := "3[a2[c]]"
	fmt.Println(decodeString(input))

}
