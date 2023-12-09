package main


/*
64. 最小路径和
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。



示例 1：
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例 2：
输入：grid = [[1,2,3],[4,5,6]]
输出：12

*/

func minPathSum(grid [][]int) int {
	// 获取网格的行数和列数
	m := len(grid[0]) // 列数
	n := len(grid)    // 行数

	// 创建一个二维动态规划数组dp，dp[i][j]表示从起点到(i, j)位置的最小路径和
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	// 初始化起点位置的路径和
	dp[0][0] = grid[0][0]

	// 遍历二维数组的每个元素
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				// 如果在第一行，只能向右走，所以当前位置的路径和等于左边位置的路径和加上当前位置的值
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}
			if j == 0 {
				// 如果在第一列，只能向下走，所以当前位置的路径和等于上方位置的路径和加上当前位置的值
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}
			// 对于其他位置(i, j)，路径和等于上方位置和左边位置的路径和中较小的一个加上当前位置的值
			dp[i][j] = min1(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	// 返回右下角位置的最小路径和，即dp数组的最后一个元素
	return dp[n-1][m-1]
}

// 辅助函数，返回两个整数中的较小值
func min1(a, b int) int {
	if a < b {
		return a
	}
	return b
}


// 辅助函数，返回两个整数中的较小值
func min1(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main(){
	grid := [][]int{{1,3,1},{1,5,1},{4,2,1}}
	print(minPathSum(grid))
}

