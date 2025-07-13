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

// 思路错误
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	result := 0
	left := 0
	n := len(height)
	tmp := 0

	for i := 1; i < n; i++ {
		/*
				1. 只累加 min(left, height[i]) * 宽度，却没有减掉柱子的高度
				计算雨水体积时应该是
				(两端较矮高度) × (宽度) － (宽度区间内所有柱子的高度和)
			2. “左高右低”的情况完全失效
			你的算法只有在 height[i] > height[left] 时才会把 tmp 加入 result。那如果最高的柱子在最左边呢？

			让我们用一个例子来走一遍你的代码：height = [4, 1, 2]

			left = 0, height[left] = 4, result = 0, tmp = 0。
			i = 1:
			height[1] (1) 不大于 height[left] (4)。
			tmp += min(4, 1) * (1 - 0 - 1) = 1 * 0 = 0。tmp 仍然是 0。
			i = 2:
			height[2] (2) 不大于 height[left] (4)。
			tmp += min(4, 2) * (2 - 0 - 1) = 2 * 1 = 2。 tmp 变为 2。
			循环结束。
			因为 height[i] > height[left] 这个条件始终没有满足，result += tmp 从未被执行。最终函数返回 result = 0。而对于 [4, 1, 2]，正确答案是在 height[1] 上方可以接 min(4, 2) - 1 = 1 的雨水，所以正确结果应该是 1。
		*/
		tmp += (min(height[left], height[i])) * (i - left - 1)
		if height[i] > height[left] {
			result += tmp
			tmp = 0
			left = i
		}
	}
	return result
}

func trap(height []int) int {
	if len(height) < 3 { // 少于3根柱子无法接水
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	result := 0

	for left < right {
		// 水坑的储水能力取决于较短的那一边
		if height[left] < height[right] {
			// 如果左边的墙比目前记录的leftMax要高，它就无法储水，更新leftMax
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				// 否则，它可以储水，储水量为 leftMax - height[left]
				result += leftMax - height[left]
			}
			left++
		} else { // height[left] >= height[right]
			// 如果右边的墙比目前记录的rightMax要高，它就无法储水，更新rightMax
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				// 否则，它可以储水，储水量为 rightMax - height[right]
				result += rightMax - height[right]
			}
			right--
		}
	}
	return result
}

// 双指针解法：
func trap(height []int) int {
	sum := 0 // 存储积水的总量
	n := len(height)

	// 分别用 lh 和 rh 数组来记录当前位置左边最高的柱子高度和右边最高的柱子高度
	lh := make([]int, n)
	rh := make([]int, n)

	lh[0] = height[0]     // 初始化 lh[0] 为第一个柱子的高度
	rh[n-1] = height[n-1] // 初始化 rh[n-1] 为最后一个柱子的高度

	// 计算每个位置左边最高的柱子高度
	for i := 1; i < n; i++ {
		lh[i] = max(lh[i-1], height[i])
	}

	// 计算每个位置右边最高的柱子高度
	for i := n - 2; i >= 0; i-- {
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
