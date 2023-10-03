package main

/*
79. 单词搜索
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例 1：
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出：true
示例 3：
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出：false

*/

/*
实现思路：
这道题目可以使用深度优先搜索（DFS）算法来解决。首先，我们遍历整个矩阵，
找到第一个字符匹配的位置，然后从这个位置开始使用 DFS 递归地搜索剩余部分的单词。
在 DFS 中，我们递归地搜索当前位置的上、下、左、右四个相邻位置，如果能够在相邻位置中找到剩余部分的单词，
返回 true。如果在某个位置无法找到匹配的字符，
或者递归搜索时超出了矩阵的边界，或者当前位置已经被使用过，都返回 false。在搜索的过程中，
使用一个二维数组 used 来记录每个位置的使用状态，避免重复使用相同的位置。遍历整个矩阵后，
如果都没有找到匹配的单词，返回 false。
*/
// exist 函数判断在给定的字符矩阵 board 中是否能够找到单词 word
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	// canFind 函数递归地在 board 中查找单词的剩余部分
	var canFind func(r, c, i int) bool
	canFind = func(r, c, i int) bool {
		// 如果已经找到了整个单词，返回 true
		if i == len(word) {
			return true
		}
		// 如果当前位置超出矩阵边界，返回 false
		if r < 0 || r >= m || c < 0 || c >= n {
			return false
		}
		// 如果当前位置已经被使用过，或者当前位置的字符不匹配单词的当前字符，返回 false
		if used[r][c] || board[r][c] != word[i] {
			return false
		}
		// 将当前位置标记为已使用
		used[r][c] = true
		// 递归搜索当前位置的上、下、左、右四个相邻位置
		canFindRest := canFind(r+1, c, i+1) || canFind(r-1, c, i+1) ||
			canFind(r, c+1, i+1) || canFind(r, c-1, i+1)
		// 如果能够在相邻位置中找到剩余部分的单词，返回 true
		if canFindRest {
			return true
		} else {
			// 如果不能找到，将当前位置的使用状态恢复为未使用，返回 false
			used[r][c] = false
			return false
		}
	}

	// 遍历整个矩阵，找到第一个字符匹配的位置，然后调用 canFind 函数查找剩余部分
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && canFind(i, j, 0) {
				return true
			}
		}
	}
	// 如果整个矩阵都搜索完毕都没有找到匹配的单词，返回 false
	return false
}
