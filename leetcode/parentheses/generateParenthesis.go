package parentheses

func generateParenthesis(n int) []string {
	var res []string
	dfs8(n, &res, "", 0, 0)
	return res
}

func dfs8(n int, res *[]string, path string, left, right int) {
	if len(path) == 2*n {
		*res = append(*res, path)
		return
	}
	if left < n {
		dfs8(n, res, path+"(", left+1, right)
	}
	if right < left {
		dfs8(n, res, path+")", left, right+1)
	}

}
