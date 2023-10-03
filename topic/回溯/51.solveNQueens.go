package main

import (
	"fmt"
	"strings"
)

/*
51. N 皇后

按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

示例 1：
输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：
输入：n = 1
输出：[["Q"]]
*/

func solveNQueens(n int) (ans [][]string) {
	col := make([]int, n)       // 列上是否有皇后的标志数组
	onPath := make([]bool, n)    // 当前路径上是否有皇后的标志数组
	diag1 := make([]bool, n*2-1) // 主对角线上是否有皇后的标志数组
	diag2 := make([]bool, n*2-1) // 副对角线上是否有皇后的标志数组

	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			board := make([]string, n)
			for i, c := range col {
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)
			}
			ans = append(ans, board) // 找到一个解，加入结果集
			return
		}
		for c, on := range onPath {
			rc := r - c + n - 1
			if !on && !diag1[r+c] && !diag2[rc] {
				col[r] = c                    // 在第r行第c列放置皇后
				onPath[c], diag1[r+c], diag2[rc] = true, true, true // 更新标志数组
				dfs(r + 1)                  // 继续搜索下一行
				onPath[c], diag1[r+c], diag2[rc] = false, false, false // 恢复现场
			}
		}
	}

	dfs(0) // 从第0行开始搜索
	return ans
}

func main() {
	n := 4
	result := solveNQueens(n)
	fmt.Println(result)
}
