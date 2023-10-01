package main

/*
84. 柱状图中最大的矩形
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

示例 1:
输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10
示例 2：
输入： heights = [2,4]
输出： 4

*/



/*
单调栈
解题思路：

栈的应用： 使用栈来存储柱的索引。遍历柱的高度数组，如果当前柱的高度小于栈顶柱的高度，就将栈顶柱弹出，并以弹出的柱为高计算面积。

栈的特性： 栈内的柱高度是单调递增的，这样可以确保栈内的柱在出栈时，能找到其右边界（当前遍历到的柱）和左边界（栈内下一个柱）。

计算宽度： 在柱出栈时，计算矩形的宽度。如果栈为空，宽度为当前柱的索引；如果栈非空，宽度为当前柱的索引减去栈顶柱的索引再减一。

更新最大面积： 在每次计算矩形面积时，更新最大面积值。遍历完成后，栈中剩余的柱也可以按照相同的方法计算面积，确保没有被漏掉。

这个算法的时间复杂度是O(n)，其中n是柱的数量。算法使用了栈来辅助计算，每个柱最多被入栈和出栈一次，因此时间复杂度是线性的。
*/

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0) // 用栈来保存柱的索引
	maxArea := 0 // 最大矩形面积

	for i := 0; i <= len(heights); i++ {
		var h int
		if i < len(heights) {
			h = heights[i]
		} else {
			h = 0 // 在最后追加一个高度为0的柱，确保栈中所有柱都会被处理
		}

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			// 当前柱高度小于栈顶柱高度，计算以栈顶柱为高的矩形面积
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1] // 弹出栈顶柱
			width := i // 默认宽度为当前索引
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1 // 计算宽度，栈顶柱右边的索引减去新栈顶柱的索引再减一
			}
			maxArea = max(maxArea, height*width) // 计算并更新最大面积
		}

		// 当前柱高度大于等于栈顶柱高度，入栈
		stack = append(stack, i)
	}

	return maxArea // 返回最大面积
}

// 辅助函数，返回两个数中较大的数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}




func main() {
	nums := []int{2,1,5,6,2,3}
	print(largestRectangleArea(nums))
}