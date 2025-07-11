package main

import "sort"

/*
455. 分发饼干
假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是满足尽可能多的孩子，并输出这个最大数值。

示例 1:
输入: g = [1,2,3], s = [1,1]
输出: 1
解释:
你有三个孩子和两块小饼干，3 个孩子的胃口值分别是：1,2,3。
虽然你有两块小饼干，由于他们的尺寸都是 1，你只能让胃口值是 1 的孩子满足。
所以你应该输出 1。
示例 2:
输入: g = [1,2], s = [1,2,3]
输出: 2
解释:
你有两个孩子和三块小饼干，2 个孩子的胃口值分别是 1,2。
你拥有的饼干数量和尺寸都足以让所有孩子满足。
所以你应该输出 2。


提示：

1 <= g.length <= 3 * 104
0 <= s.length <= 3 * 104
1 <= g[i], s[j] <= 231 - 1
*/

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var result int
	for len(g) > 0 && len(s) > 0 {
		sLen := len(s)
		gLen := len(g)
		if s[sLen-1] >= g[gLen-1] {
			s = s[:sLen-1]
			g = g[:gLen-1]
			result++
		} else {
			g = g[:gLen-1]
		}

	}
	return result
}

// 优化后的代码：
func findContentChildren(g []int, s []int) int {
	// 优化 1：提前处理空数组的情况
	// 说明：如果 g 或 s 为空，则无法满足任何孩子，直接返回 0
	// 这样可以避免不必要的排序和循环
	if len(g) == 0 || len(s) == 0 {
		return 0
	}

	sort.Ints(g)
	sort.Ints(s)
	var result int

	// 优化 2：使用索引变量替代切片操作
	// 原代码：for len(g) > 0 && len(s) > 0 { ... }
	// 说明：通过切片操作移除元素会导致额外的内存分配和数据拷贝，效率较低
	// 优化：使用两个索引变量 i 和 j，从数组末尾开始遍历，避免切片操作
	i := len(g) - 1 // 指向 g 的末尾
	j := len(s) - 1 // 指向 s 的末尾
	for i >= 0 && j >= 0 {
		// 优化 3：移除不必要的长度计算
		// 原代码：sLen := len(s); gLen := len(g)
		// 说明：在循环中反复计算长度是多余的，因为我们已经使用索引变量控制循环
		// 优化：直接使用索引变量 i 和 j 访问元素
		if s[j] >= g[i] {
			// 优化 4：移除切片操作
			// 原代码：s = s[:sLen-1]; g = g[:gLen-1]
			// 说明：切片操作会导致内存分配和数据拷贝，效率较低
			// 优化：通过递减索引变量 i 和 j 移动指针，避免切片操作
			result++
			i--
			j--
		} else {
			// 优化 5：移除切片操作
			// 原代码：g = g[:gLen-1]
			// 说明：同上，切片操作效率较低
			// 优化：通过递减索引变量 i 移动指针
			i--
		}
	}

	return result
}
