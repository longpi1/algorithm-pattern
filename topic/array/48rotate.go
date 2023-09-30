package main

/*
48. 旋转图像
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]

示例 2：
输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
*/

/*
解题思路：
画图后得出对应思路：
使用额外矩阵
*/
func rotate(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			tmp[j][n-1-i] = v
		}
	}
	copy(matrix, tmp) // 拷贝 tmp 矩阵每行的引用
}

// 空间复杂度优化
func rotate(matrix [][]int) {
	n := len(matrix)
	// 遍历矩阵的上半部分（包括中间行）和左半部分（包括中间列）
	for i := 0; i < n/2; i++ {
		// 遍历当前层的每一列
		for j := 0; j < (n+1)/2; j++ {
			// 将四个对应位置的元素进行旋转
			// 左上角的元素(matrix[i][j])与右上角的元素(matrix[n-j-1][i])交换
			// 右上角的元素与右下角的元素(matrix[n-i-1][n-j-1])交换
			// 右下角的元素与左下角的元素(matrix[j][n-i-1])交换
			// 左下角的元素与左上角的元素交换(matrix[i][j])
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

