package bytedance

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs7(board, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func dfs7(board [][]byte, word string, x, y, index int) bool {
	if index == len(word) {
		return true
	}

	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != word[index] {
		return false
	}

	temp := board[x][y]
	board[x][y] = '#'

	found := dfs7(board, word, x+1, y, index+1) || dfs7(board, word, x-1, y, index+1) || dfs7(board, word, x, y+1, index+1) || dfs7(board, word, x, y-1, index+1)

	board[x][y] = temp

	return found

}
