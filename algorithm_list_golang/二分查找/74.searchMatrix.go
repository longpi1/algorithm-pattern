package main

import "fmt"

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
	if matrix == nil {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	left1, right1 := 0, n-1
	left2, right2 := 0, m-1
	for left1 <= right1 {
		mid1 := left1 + (right1-left1)/2
		//fmt.Println("mid1:", mid1)
		if matrix[mid1][0] <= target && matrix[mid1][m-1] >= target {
			for left2 <= right2 {
				mid2 := left2 + (right2-left2)/2
				//fmt.Println("mid2:", mid2)
				if matrix[mid1][mid2] == target {
					return true
				}
				if matrix[mid1][mid2] > target {
					right2 = mid2 - 1
				} else {
					left2 = mid2 + 1
				}
			}

			return false
		} else if matrix[mid1][m-1] < target {
			left1 = mid1 + 1
		} else {
			right1 = mid1 - 1
		}
	}
	return false
}

//func searchMatrix(matrix [][]int, target int) bool {
//	i, j := 0, len(matrix[0])-1 // 起始点是右上角
//	for i < len(matrix) && j >= 0 {
//		if matrix[i][j] == target {
//			return true
//		} else if matrix[i][j] < target {
//			i++ // 则matrix[i][j]左边的数，均小于target, 则可排除此行, 则行号+1即可
//		} else {
//			j-- // 则matrix[i][j]上边的数，均小于target, 则可排除此列, 则列号-1即可
//		}
//	}
//	return false
//}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	falg := searchMatrix(matrix, 3)
	fmt.Println("result:", falg)
}
