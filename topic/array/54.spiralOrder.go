package main

import "fmt"

/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。



示例 1：


输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：


输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/
/*

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	var result []int

	for i := 0; i < m ; i++ {
		n := len(matrix[i])
		if i == 0 {
			for j :=0; j < n; j ++ {
				result = append(result, matrix[i][j])
			}
		}else {
			result = append(result, matrix[i][n-1])
		}
	}
	if m <= 1{
		return  result
	}
	for m == 2{
		if  len(matrix[1]) >1 {
			for j := len(matrix[1])-2 ; j >=0  ; j-- {
				result = append(result, matrix[1][j])
			}
		}
		return result
	}

	for i := m-1; i > 1 ; i-- {
		if len(matrix[i]) > 1 {
			for j := len(matrix[i]) -2; j>=0; j-- {
				result = append(result, matrix[i][j])
			}
		}

	}


	if  len(matrix[1]) >1 {
		for j := 0 ; j < len(matrix[1])-1 ; j++ {
			result = append(result, matrix[1][j])
		}
	}


		return result
}
*/

// 上述答案错误
/*
原因如下：
未总结出正确规律
以及对应的循环不变量
养成画图的好习惯

大佬思路如下：
https://leetcode.cn/problems/spiral-matrix/solutions/275716/shou-hui-tu-jie-liang-chong-bian-li-de-ce-lue-kan-/
如果一条边从头遍历到底，则下一条边遍历的起点随之变化

选择不遍历到底，可以减小横向、竖向遍历之间的影响

一轮迭代结束时，4条边的两端同时收窄 1

一轮迭代所做的事情变得很清晰：遍历一个“圈”，遍历的范围收缩为内圈

一层层向里处理，按顺时针依次遍历：上、右、下、左。

不再形成“环”了，就会剩下：一行或一列，然后单独判断

四个边界
上边界 top : 0
下边界 bottom : matrix.length - 1
左边界 left : 0
右边界 right : matrix[0].length - 1
矩阵不一定是方阵
top < bottom && left < right 是循环的条件
无法构成“环”了，就退出循环，退出时可能是这 3 种情况之一：
top == bottom && left < right —— 剩一行
top < bottom && left == right —— 剩一列
top == bottom && left == right —— 剩一项（也算 一行/列）
处理剩下的单行或单列
因为是按顺时针推入结果数组的，所以
剩下的一行，从左至右 依次推入结果数组
剩下的一列，从上至下 依次推入结果数组
代码
每个元素访问一次，时间复杂度 O(m*n)，m、n 分别是矩阵的行数和列数
空间复杂度 O(m*n)

*/
/*func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for top < bottom && left < right {
		for i := left; i < right; i++ { res = append(res, matrix[top][i]) }
		for i := top; i < bottom; i++ { res = append(res, matrix[i][right]) }
		for i := right; i > left; i-- { res = append(res, matrix[bottom][i]) }
		for i := bottom; i > top; i-- { res = append(res, matrix[i][left]) }
		right--
		top++
		bottom--
		left++
	}
	if top == bottom {
		for i := left; i <= right; i++ { res = append(res, matrix[top][i]) }
	} else if left == right {
		for i := top; i <= bottom; i++ { res = append(res, matrix[i][left]) }
	}
	return res
}*/
/*
换一种遍历的策略：遍历到底

循环的条件改为： top <= bottom && left <= right
每遍历一条边，下一条边遍历的起点被“挤占”，要更新相应的边界
注意到，可能在循环途中，突然不再满足循环的条件，即top > bottom或left > right，其中一对边界彼此交错了
这意味着所有项都遍历完了，要break了，如果没有及时 break ，就会重复遍历
解决办法
每遍历完一条边，更新相应的边界后，都加上一句if (top > bottom || left > right) break;，避免因没有及时退出循环而导致重复遍历。
且，“遍历完成”这个时间点，要么发生在遍历完“上边”，要么发生在遍历完“右边”
所以只需在这两步操作之后，加 if (top > bottom || left > right) break 即可

*/

/*func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ { res = append(res, matrix[top][i]) }
		top++
		for i := top; i <= bottom; i++ { res = append(res, matrix[i][right]) }
		right--
		if top > bottom || left > right { break }
		for i := right; i >= left; i-- { res = append(res, matrix[bottom][i]) }
		bottom--
		for i := bottom; i >= top; i-- { res = append(res, matrix[i][left]) }
		left++
	}

	return res
}*/


/*
以下是自己再做一遍的结果：
错误点如下：
1. bottom是等于len(matrix)-1，right是等于len(matrix[0])-1,这个点看图就可以得出来。。。。。
2. 添加元素错误matrix[i][top]和 matrix[right][i]等都搞反了。。。。。。。。
3.矩阵不一定是方阵
top < bottom && left < right 是循环的条件
无法构成“环”了，就退出循环，退出时可能是这 3 种情况之一：
top == bottom && left < right —— 剩一行
top < bottom && left == right —— 剩一列
top == bottom && left == right —— 剩一项（也算 一行/列）
处理剩下的单行或单列
因为是按顺时针推入结果数组的，所以
剩下的一行，从左至右 依次推入结果数组
剩下的一列，从上至下 依次推入结果数组
这里应该是<=才行
*/
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	var result []int
	// bottom是等于len(matrix)-1，right是等于len(matrix[0])-1,这个点看图就可以得出来
	//top,bottom,left,right := 0, n-1, 0, m-1
	top,bottom,left,right := 0, m-1, 0, n-1
	for top < bottom && left < right {
		for i := left; i < right; i++ {
			//添加元素错误matrix[i][top]和 matrix[right][i]等都搞反了。。。。。。。。
			//result = append(result, matrix[i][top])
			result = append(result, matrix[top][i])
		}
		for i := top; i < bottom; i++ {
			//result = append(result, matrix[right][i])
			result = append(result, matrix[i][right])
		}
		for i := right; i > left; i-- {
			//result = append(result, matrix[i][top])
			result = append(result, matrix[bottom][i])
		}
		for i := bottom; i > top; i-- {
			//result = append(result, matrix[left][i])
			result = append(result, matrix[i][left])
		}

		left ++
		right --
		top ++
		bottom --
	}

	if top == bottom{
		//这里应该是<=才行
		for i := left; i <= right; i++ {
			//result = append(result, matrix[i][top])
			result = append(result, matrix[top][i])
		}
	}else if left == right {
	// 这里要改为else if 两者情况不可能同时方式
	//	if left == right {
		for i := top; i <= bottom; i++ {
			//result = append(result, matrix[right][i])
			result = append(result, matrix[i][right])
		}
	}

	return result
}



func main()  {
	//nums := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	nums1 := [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	fmt.Printf("result: %v",spiralOrder(nums1))
}