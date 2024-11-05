package bytedance

// subsets 生成给定集合的所有子集
func subsets(nums []int) [][]int {
	var result [][]int
	var path []int
	dfs5(nums, 0, path, &result)
	return result
}

func dfs5(nums []int, index int, path []int, res *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	for i := index; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs5(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}
