package leetcode

func maxAreaOfIsland(grid [][]int) int {
	var maxArea int

	// 遍历整个地图，找到每个岛屿并计算其面积
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area := dfs6(i, j, grid)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func dfs6(i, j int, grid [][]int) int {
	if i <= 0 || i >= len(grid) || j <= 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1 + dfs6(i-1, j, grid) + dfs6(i+1, j, grid) + dfs6(i, j-1, grid) + dfs6(i, j+1, grid)

}
