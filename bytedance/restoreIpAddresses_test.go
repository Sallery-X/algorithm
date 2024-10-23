package bytedance

import (
	"fmt"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}
