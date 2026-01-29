package matrix

/*
func numIslands(grid [][]byte) int {
	result := 0
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				result++
			}
		}
	}

	return result

}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
}


*/

func numIslands(grid [][]int) int {
	res := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs1(grid, i, j)
				res++
			}
		}
	}
	return res

}
func dfs1(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs1(grid, i, j-1)
	dfs1(grid, i, j+1)
	dfs1(grid, i-1, j)
	dfs1(grid, i+1, j)
}
