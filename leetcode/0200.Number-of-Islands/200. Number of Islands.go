package leetcode

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n, res := len(grid), len(grid[0]), 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(grid, m, n, i, j)
				res++
			}
		}
	}
	return res
}
func dfs(grid [][]byte, m int, n int, x int, y int) {
	if !isSafe(grid, m, n, x, y) {
		return
	}
	grid[x][y] = '0'
	dx, dy := []int{0, 0, -1, 1}, []int{-1, 1, 0, 0}
	for i := 0; i < 4; i++ {
		dfs(grid, m, n, x+dx[i], y+dy[i])
	}
}
func bfs(grid [][]byte, m int, n int, x int, y int) {
	dx, dy, queue := [4]int{0, 0, -1, 1}, [4]int{-1, 1, 0, 0}, [][2]int{{x, y}}
	grid[x][y] = '0'
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			px, py := p[0]+dx[i], p[1]+dy[i]
			if isSafe(grid, m, n, px, py) {
				queue = append(queue, [2]int{px, py})
				grid[px][py] = '0'
			}
		}
	}
}
func isSafe(grid [][]byte, m int, n int, x int, y int) bool {
	return !(x < 0 || y < 0 || x >= m || y >= n || grid[x][y] == '0')
}
