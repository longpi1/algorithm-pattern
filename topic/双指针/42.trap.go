package main
/*
42. 接雨水

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

*/


// 双指针解法：
func trap(height []int) int {
	sum := 0 // 存储积水的总量
	n := len(height)

	// 分别用 lh 和 rh 数组来记录当前位置左边最高的柱子高度和右边最高的柱子高度
	lh := make([]int, n)
	rh := make([]int, n)

	lh[0] = height[0] // 初始化 lh[0] 为第一个柱子的高度
	rh[n-1] = height[n-1] // 初始化 rh[n-1] 为最后一个柱子的高度

	// 计算每个位置左边最高的柱子高度
	for i := 1; i < n; i++ {
		lh[i] = max(lh[i-1], height[i])
	}

	// 计算每个位置右边最高的柱子高度
	for i := n-2; i >= 0; i-- {
		rh[i] = max(rh[i+1], height[i])
	}

	// 遍历每个位置，计算积水量并累加到 sum 中
	for i := 1; i < n-1; i++ {
		h := min(rh[i], lh[i]) - height[i] // 当前位置的积水高度
		if h > 0 {
			sum += h
		}
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
