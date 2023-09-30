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


/*
从右上角开始搜索的策略是为了利用矩阵的特殊排列顺序，提高搜索效率。在这种矩阵中，每一行从左到右递增，每一列从上到下递增。由于这种特性，右上角的元素是该行最大的，同时也是该列最小的。

通过从右上角开始搜索，我们可以根据目标值与当前元素的大小关系，快速地缩小搜索范围：

如果当前元素等于目标值，搜索成功，直接返回true。
如果当前元素小于目标值，由于当前行的元素都比该元素小，可以排除该行，所以行号加1，继续在剩余的矩阵中搜索。
如果当前元素大于目标值，由于当前列的元素都比该元素大，可以排除该列，所以列号减1，继续在剩余的矩阵中搜索。
这样，每一步都能够通过比较当前元素与目标值的大小，排除一整行或一整列，从而迅速地缩小搜索范围。这种策略在时间复杂度上比较高效，是一种常用的二维矩阵搜索方法。
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

