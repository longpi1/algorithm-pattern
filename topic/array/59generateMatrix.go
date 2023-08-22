package main

import "fmt"

/*
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。



示例 1：


输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
示例 2：

输入：n = 1
输出：[[1]]


提示：

1 <= n <= 20
*/

/*func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	} 
	result := make([][]int, n)
	top,left,bottom,right := 0,0,n-1,n-1
	// 二维切片，需要初始化一维 ！！！
	for i := range result {
		result[i] = make([]int, n)
	}
	// 这里应该初始化为1
	//var count int
	count := 1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {

			result[top][i] = count
			count ++
		}
		for i := top; i <=bottom ; i++ {

			result[i][right] = count
			count ++
		}
		for i := right; i >= left; i-- {

			result[bottom][i] = count
			count ++
		}
		for i := bottom; i >= top; i-- {

			result[i][left] = count
			count ++
		}
		left ++
		top ++
		bottom --
		right --
	}
	return result
}*/

/*

上述代码存在以下几个问题：
1.二维切片，需要初始化一维 ！！！
2.这里改为<=和>= 需要遍历完
构建 n * n 的矩阵

确定矩阵的四个边界，它是初始遍历的边界。
按 上 右 下 左，一层层向内，遍历矩阵填格子
每遍历一个格子，填上对应的 num，num 自增
直到 left <= right && top <= bottom ，遍历结束

链接：https://leetcode.cn/problems/spiral-matrix-ii/solutions/659084/ru-guo-ni-yuan-yi-yi-ceng-yi-ceng-yi-cen-cm9h/

*/

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	result := make([][]int, n)
	// 1.二维切片，需要初始化一维 ！！！
	for i := range result {
		result[i] = make([]int, n)
	}

	top, left, bottom, right := 0, 0, n-1, n-1
	count := 1
	// 2.这里需要为<= 需要遍历完
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ { // top layer
			result[top][i] = count
			count++
		}
		top++                            // move to next layer
		for i := top; i <= bottom; i++ { // right layer
			result[i][right] = count
			count++
		}
		right--
		for i := right; i >= left; i-- { // bottom layer
			result[bottom][i] = count
			count++
		}
		bottom--
		for i := bottom; i >= top; i-- { // left layer
			result[i][left] = count
			count++
		}
		left++

	}

	return result
}





func main()  {
	//nums := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	nums1 := 3
	fmt.Printf("result: %v",generateMatrix(nums1))
}