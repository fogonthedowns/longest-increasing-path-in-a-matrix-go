package main

func longestIncreasingPath(matrix [][]int) (ans int) {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(matrix, i, j, cache))
		}
	}
	return ans
}

func dfs(matrix [][]int, i int, j int, cache [][]int) int {
	if cache[i][j] != 0 {
		return cache[i][j]
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for index, _ := range dirs {

		x := i + dirs[index][0]
		y := j + dirs[index][1]

		if 0 <= x && x < len(matrix) && 0 <= y && y < len(matrix[0]) && matrix[x][y] > matrix[i][j] {
			cache[i][j] = max(cache[i][j], dfs(matrix, x, y, cache))
		}
	}
	cache[i][j] = cache[i][j] + 1
	return cache[i][j]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
