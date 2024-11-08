package matrix

import (
	"fmt"
	"testing"
)

func Test_maximalSquare(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println("Maximal Square Area:", maximalSquare(matrix)) // Output: 4
}
