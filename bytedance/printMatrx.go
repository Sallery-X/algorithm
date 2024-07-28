package bytedance

import "fmt"

func NMatrix(n int) {
	// 创建一个 n*n 的矩阵
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// 填充矩阵
	value := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i+j)%2 == 0 {
				matrix[(i+j)%n][i] = value
			} else {
				matrix[i][(i+j)%n] = value
			}
			value++
		}
	}

	// 输出矩阵
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
