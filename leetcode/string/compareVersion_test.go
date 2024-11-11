package string

import (
	"fmt"
	"testing"
)

func Test_compareVersion(t *testing.T) {
	fmt.Println(compareVersion("1.2", "1.10"))
}
