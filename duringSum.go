package main

import (
	"fmt"
)

// 计算区间和
func main() {
	// 读取n
	var n int
	fmt.Scan(&n)

	// 初始化数组和前缀和数组
	vec := make([]int, n)
	p := make([]int, n)

	// 计算前缀和
	presum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&vec[i]) // 读取每个元素
		presum += vec[i]  // 计算前缀和
		p[i] = presum     // 存储到前缀和数组
	}

	// 处理查询
	var a, b int
	for {
		// 读取区间[a, b]
		_, err := fmt.Scan(&a, &b)
		if err != nil {
			break // 如果没有更多的输入，退出循环
		}

		// 根据前缀和计算区间和
		var sum int
		if a == 0 {
			sum = p[b] // 区间 [0, b]
		} else {
			sum = p[b] - p[a-1] // 区间 [a, b]
		}

		// 输出结果
		fmt.Println(sum)
	}
}
