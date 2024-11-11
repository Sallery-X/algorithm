package dfs

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int

	dfs4(candidates, target, &res, path, 0)

	return res

}

func dfs4(candidates []int, target int, res *[][]int, path []int, start int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] <= target {
			path = append(path, candidates[i])
			dfs4(candidates, target-candidates[i], res, path, i)
			path = path[:len(path)-1]
		}
	}

}
