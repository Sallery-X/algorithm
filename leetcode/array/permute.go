package array

func permute(nums []int) [][]int {
	var res [][]int

	dfs2([]int{}, make([]bool, len(nums)), nums, &res)

	return res
}

func dfs2(path []int, used []bool, nums []int, res *[][]int) {

	//fmt.Println(path)
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		//fmt.Println(nums[i], used[i])
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		//fmt.Println(nums[i], used[i])
		dfs2(path, used, nums, res)
		path = path[:len(path)-1]
		used[i] = false
	}
}
