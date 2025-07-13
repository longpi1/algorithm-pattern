package main

import "sort"

/*
56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/

func merge(intervals [][]int) [][]int {

}

func merge(intervals [][]int) [][]int {
	// 按照区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, len(intervals))         // 用于存储合并后的区间结果
	left, right := intervals[0][0], intervals[0][1] // 初始化当前合并区间的左右边界

	// 遍历排序后的区间列表
	for i := 1; i < len(intervals); i++ {
		if right < intervals[i][0] { // 如果当前合并区间的右边界小于下一个区间的左边界
			res = append(res, []int{left, right})          // 将当前合并区间添加到结果中
			left, right = intervals[i][0], intervals[i][1] // 更新合并区间的左右边界为下一个区间
		} else {
			right = max(right, intervals[i][1]) // 如果有重叠，更新合并区间的右边界
		}
	}
	res = append(res, []int{left, right}) // 将最后一个合并区间添加到结果中
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
