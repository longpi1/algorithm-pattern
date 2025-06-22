package main

import "fmt"

/*
63. 不同路径 II

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
网格中的障碍物和空位置分别用 1 和 0 来表示。


示例 1：
输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
输出：2
解释：3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
示例 2：
输入：obstacleGrid = [[0,1],[0,0]]
输出：1
*/

/*
下述代码错误逻辑：

1.移除了冗余的判断：原代码中检查前一个格子是否是障碍物的逻辑是多余的，因为如果前一个格子是障碍物，它的值已经被设为0了。

2.修正了第一行和第一列的初始化：
第一行的格子应该等于左边格子的值（如果左边是障碍物或无法到达，值为0）
第一列的格子应该等于上面格子的值（如果上面是障碍物或无法到达，值为0）这样可以正确处理路径被阻断的情况
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	result := make([][]int, m)

	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				result[i][j] = 0
				//fmt.Printf("m,n,result", i, j, result[i][j])
				continue
			}
			if j > 0 && i == 0 && obstacleGrid[i][j-1] == 1 {
				//fmt.Printf("m,n,result", i, j, result[i][j])
				result[i][j] = 0
				continue
			}
			if i > 0 && j == 0 && obstacleGrid[i-1][j] == 1 {
				//fmt.Printf("m,n,result", i, j, result[i][j])
				result[i][j] = 0
				continue
			}
			if i == 0 || j == 0 {
				result[i][j] = 1
				continue
			}

			result[i][j] = result[i-1][j] + result[i][j-1]
			// fmt.Printf("m,n,result", i, j, result[i][j])
		}
	}
	return result[m-1][n-1]
}

// 修正后代码
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	result := make([][]int, m)

	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			// 如果当前位置是障碍物，路径数为0
			if obstacleGrid[i][j] == 1 {
				result[i][j] = 0
				continue
			}

			// 起点位置
			if i == 0 && j == 0 {
				result[i][j] = 1
				continue
			}

			// 第一行：只能从左边来
			if i == 0 {
				result[i][j] = result[i][j-1]
				continue
			}

			// 第一列：只能从上面来
			if j == 0 {
				result[i][j] = result[i-1][j]
				continue
			}

			// 其他位置：可以从上面或左边来
			result[i][j] = result[i-1][j] + result[i][j-1]
		}
	}
	return result[m-1][n-1]
}

func main() {
	obstacleGrid1 := [][]int{{1, 0}}
	//obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	result := uniquePathsWithObstacles(obstacleGrid1)
	fmt.Printf("result: %v", result)
}

/*
本题是62.不同路径的障碍版，整体思路大体一致。

但就算是做过62.不同路径，在做本题也会有感觉遇到障碍无从下手。

其实只要考虑到，遇到障碍dp[i][j]保持0就可以了。

也有一些小细节，例如：初始化的部分，很容易忽略了障碍之后应该都是0的情况。
*/
//func uniquePathsWithObstacles(obstacleGrid [][]int) int {
//	m, n := len(obstacleGrid), len(obstacleGrid[0])
//	// 定义一个dp数组
//	dp := make([][]int, m)
//	for i, _ := range dp {
//		dp[i] = make([]int, n)
//	}
//	// 初始化, 如果是障碍物, 后面的就都是0, 不用循环了
//	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
//		dp[i][0] = 1
//	}
//	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
//		dp[0][i] = 1
//	}
//	// dp数组推导过程
//	for i := 1; i < m; i++ {
//		for j := 1; j < n; j++ {
//			// 如果obstacleGrid[i][j]这个点是障碍物, 那么dp[i][j]保持为0
//			if obstacleGrid[i][j] != 1 {
//				// 否则我们需要计算当前点可以到达的路径数
//				dp[i][j] = dp[i-1][j] + dp[i][j-1]
//			}
//		}
//	}
//	return dp[m-1][n-1]
//}
