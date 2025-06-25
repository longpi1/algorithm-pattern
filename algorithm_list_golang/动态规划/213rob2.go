package main

/*
这道题目和198.打家劫舍是差不多的，唯一区别就是成环了。
分两种情况考虑：
1.考虑包含首元素，不包含尾元素
2.考虑包含尾元素，不包含首元素
《代码随想录》算法视频公开课：动态规划，房间连成环了那还偷不偷呢？| LeetCode：213.打家劫舍II，相信结合视频再看本篇题解，更有助于大家对本题的理解。

分类讨论，考虑是否偷 nums[0]：
如果偷 nums[0]，那么 nums[1] 和 nums[n−1] 不能偷，问题变成从 nums[2] 到 nums[n−2] 的非环形版本，调用 198 题的代码解决；
如果不偷 nums[0]，那么问题变成从 nums[1] 到 nums[n−1] 的非环形版本，同样调用 198 题的代码解决。

*/

func dfs(nums []int) int {
	// 初始化两个变量，用于记录前两个房屋的最大抢劫价值
	first, second := nums[0], max(nums[0], nums[1])

	// 从第三个房屋开始遍历，依次计算每个房屋的最大抢劫价值
	for _, v := range nums[2:] {
		// 计算当前房屋的最大抢劫价值，同时更新first和second
		first, second = second, max(first+v, second)
	}

	// 返回最终结果，即抢劫所有房屋的最大价值
	return second
}

func rob(nums []int) int {
	n := len(nums)

	// 处理特殊情况，如果只有一个或两个房屋
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	// 将问题分解成两个子问题，一个包括第一个房屋，另一个包括最后一个房屋
	// 返回两个子问题的最大抢劫价值中的较大者
	return max(dfs(nums[:n-1]), dfs(nums[1:]))
}
