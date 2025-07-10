package main

/*
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和。
示例 1:
输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
示例 2:
输入: numRows = 1
输出: [[1]]
*/
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if i == 0 || j == 0 || j == i {
				result[i][j] = 1
				continue
			}
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	// 切记这里要初始化二维数组中的一维数组
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if i < 2 {
				result[i][j] = 1
			} else if j == 0 || j == i {
				result[i][j] = 1
			} else {
				// 推理范式
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}
	return result
}

func main() {
	numRows := 5
	generate(numRows)
}
