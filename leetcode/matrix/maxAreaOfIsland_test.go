package matrix

import (
	"fmt"
	"testing"
)

func Test_maxAreaOfIsland(t *testing.T) {
	// 定义一个二维数组表示地图，其中1表示陆地，0表示水域
	var grid = [][]int{
		{0, 1, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 1, 0},
	}
	maxArea := maxAreaOfIsland(grid)
	fmt.Println("最大岛屿面积:", maxArea)
}
