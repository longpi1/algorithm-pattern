package main

import "fmt"

/*
73. 矩阵置零
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

示例 1：
输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]
示例 2：
输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
*/
func setZeroes(matrix [][]int) {
	rowsMap := make(map[int]bool)
	columnsMap := make(map[int]bool)
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				rowsMap[i] = true
				columnsMap[j] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		_, isRowHas := rowsMap[i]
		for j := 0; j < m; j++ {
			_, isColumnHas := columnsMap[j]
			if isColumnHas || isRowHas {
				matrix[i][j] = 0
			}
		}
	}
}

/*
func setZeroes(matrix [][]int)  {
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {

			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if
			matrix[i][j] = 0
		}
	}

}*/

func setZeroes2(matrix [][]int) {
	// 创建两个布尔数组，用于记录哪些行和列需要被置零
	row := make([]bool, len(matrix))    // 记录行的状态
	col := make([]bool, len(matrix[0])) // 记录列的状态

	// 第一次遍历矩阵，标记为零的行和列
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true // 标记当前行需要置零
				col[j] = true // 标记当前列需要置零
			}
		}
	}

	// 第二次遍历矩阵，根据标记置零行和列
	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0 // 根据标记将行和列置零
			}
		}
	}
}

// 优化方法，使用两个标志位，空间复杂度为O（1）
//我们可以对方法二进一步优化，只使用一个标记变量记录第一列是否原本存在 0。
//这样，第一列的第一个元素即可以标记第一行是否出现 000。但为了防止每一列的第一个元素被提前更新，
//我们需要从最后一行开始，倒序地处理矩阵元素。

func setZeroes3(matrix [][]int) {
	// 获取矩阵的行数和列数
	n, m := len(matrix), len(matrix[0])
	// 初始化两个标志位，用于标记第一行和第一列是否需要被置零
	row0, col0 := false, false

	// 检查第一行是否有零
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}

	// 检查第一列是否有零
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}

	// 遍历矩阵，将出现零的行和列的第一个元素置零
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 根据第一行和第一列的标记，将相应的行和列置零
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 根据标记，将第一行和第一列置零
	if row0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}

func main() {
	nums := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(nums)
	fmt.Printf("result: %v", nums)
}
