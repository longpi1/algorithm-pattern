package main

/*
11. 盛最多水的容器
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

示例 1：

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：

输入：height = [1,1]
输出：1

*/
func maxArea(height []int) int {
	left, right := 0, len(height) - 1 // 初始化左右指针分别指向数组的首尾
	max := 0 // 用于记录最大的容器面积

	for left < right { // 当左指针小于右指针时，继续循环
		tmp := 0 // 用于暂时存储当前容器的面积
		width := right - left // 计算容器的宽度

		if height[left] > height[right] { // 如果左边的高度大于右边的高度
			tmp = height[right] * width // 计算容器的面积
			right-- // 右指针向左移动，缩小容器宽度
		} else {
			tmp = height[left] * width // 如果右边的高度大于等于左边的高度，计算容器的面积
			left++ // 左指针向右移动，缩小容器宽度
		}

		if tmp > max { // 如果当前容器的面积大于最大面积
			max = tmp // 更新最大面积
		}
	}
	return max // 返回最大面积
}
