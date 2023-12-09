package main


/*
74. 搜索二维矩阵
给你一个满足下述两条属性的 m x n 整数矩阵：
每行中的整数从左到右按非递减顺序排列。
每行的第一个整数大于前一行的最后一个整数。
给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

示例 1：
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true
示例 2：
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false

*/

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1 // 起始点是右上角
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++ // 则matrix[i][j]左边的数，均小于target, 则可排除此行, 则行号+1即可
		} else {
			j-- // 则matrix[i][j]上边的数，均小于target, 则可排除此列, 则列号-1即可
		}
	}
	return false
}