package matrix

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	startX := 0
	startY := 0
	loop := n / 2
	center := n / 2
	count := 1
	bian := 1

	for loop > 0 {
		i, j := startX, startY
		for j = startY; j < n-bian; j++ {
			res[startX][j] = count
			count++
		}
		for i = startX; i < n-bian; i++ {
			res[i][j] = count
			count++
		}
		for ; j > startY; j-- {
			res[i][j] = count
			count++
		}
		for ; i > startX; i-- {
			res[i][j] = count
			count++
		}

		startX++
		startY++
		loop--
		bian++

	}

	if n%2 == 1 {
		res[center][center] = n * n
	}

	return res

}
