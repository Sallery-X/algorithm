package dp

import (
	"fmt"
	"testing"
)

func Test_exist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println("Word found:", exist(board, word)) // Output: true

	word = "SEE"
	fmt.Println("Word found:", exist(board, word)) // Output: true

	word = "ABCB"
	fmt.Println("Word found:", exist(board, word)) // Output: false
}
