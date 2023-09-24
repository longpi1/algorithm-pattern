package main

import "sort"

/*
240. 搜索二维矩阵 II
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。


示例 1：
输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true
示例 2：
输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
输出：false

*/

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

/*
二分法优化
*/
func searchMatrix(matrix [][]int, target int) bool {
	// 遍历矩阵的每一行
	for _, row := range matrix {
		// 使用二分查找在当前行中查找目标值
		i := sort.SearchInts(row, target)

		// 如果找到目标值并且在当前行中存在，返回 true
		if i < len(row) && row[i] == target {
			return true
		}
	}

	// 如果遍历完所有行都没有找到目标值，返回 false
	return false
}

