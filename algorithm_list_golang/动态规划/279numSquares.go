package main

import "math"

/*
279. 完全平方数
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。



示例 1：
输入：n = 12
输出：3
解释：12 = 4 + 4 + 4

示例 2：
输入：n = 13
输出：2
解释：13 = 4 + 9
*/
/*
解题思路：

可能刚看这种题感觉没啥思路，又平方和的，又最小数的。

我来把题目翻译一下：完全平方数就是物品（可以无限件使用），凑个正整数n就是背包，问凑满这个背包最少有多少物品？

*/
// 将题目转换为背包问题,转换后其实也就是零钱兑换的思路。。
func numSquares(n int) int {
	// 创建动态规划数组，dp[i]表示整数i的最小完全平方数个数
	dp := make([]int, n+1)

	// 初始化dp数组，dp[0]为0，其他元素初始为最大整数值
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}

	// 创建一个存储完全平方数的数组，nums[i]表示i的平方
	nums := make([]int, n+1)
	for i := 0; i <= n; i++ {
		nums[i] = i * i
	}

	// 遍历1到n的每个整数
	for i := 1; i <= n; i++ {
		// 遍历完全平方数数组
		for j := nums[i]; j <= n; j++ {
			// 动态规划递推公式：dp[j] = min(dp[j], dp[j-nums[i]] + 1)
			// dp[j]表示当前j的最小完全平方数个数，
			// dp[j-nums[i]]表示去掉一个完全平方数后的子问题的最小个数，
			// 加1表示当前的完全平方数加1个
			dp[j] = min(dp[j], dp[j-nums[i]] + 1)
		}
	}

	// 返回n的最小完全平方数个数
	return dp[n]
}

// 辅助函数，返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func min(a int, b int) int{
	if a < b {
		return a
	}
	return b
}