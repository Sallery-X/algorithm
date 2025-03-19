package matrix

import "fmt"

// 旋转数组
func rotate(matrix [][]int) {
	n := len(matrix)

	/*
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n/2; j++ {
				matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
			}
		}*/
	/*res := make([][]int, n)
	for i, _ := range res {
		res[i] = make([]int, n)
	}*/

	var res [][]int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = matrix[n-i-1][n-j-1]
		}
	}
	fmt.Println(res)

}
