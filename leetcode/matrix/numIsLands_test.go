package matrix

import (
	"fmt"
	"testing"
)

func Test_nums(t *testing.T) {
	a := [][]byte{
		{'1', '0', '0'},
		{'0', '1', '0'},
		{'0', '0', '1'},
	}
	fmt.Println(numIslands(a))
}
