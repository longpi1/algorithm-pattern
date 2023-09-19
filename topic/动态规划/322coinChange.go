package main

import (
	"math"
)

/*
322. 零钱兑换
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。

示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0
*/

/*func coinChange(coins []int, amount int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	count := 0
	result := -1
	dfs := func(count int) {}
	dfs = func(amount int) {
		if amount == 0 {
			result = count
			return
		}
		if amount < 0 {
			return
		}
		for i := 0; i < len(coins); i++ {
			amount -= coins[i]
			count ++
			dfs(amount)
			amount += coins[i]
			count --
		}

	}

	dfs(amount)
	return result
}
*/

/*
上述采用回溯思路存在超时问题；
*/

func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}
	for i := 0; i< n; i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt64 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}

		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}





// 版本一, 先遍历物品,再遍历背包
func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化dp[0]
	dp[0] = 0
	// 初始化为math.MaxInt32
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}

	// 遍历物品
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// 版本二,先遍历背包,再遍历物品
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化dp[0]
	dp[0] = 0
	// 遍历背包,从1开始
	for j := 1; j <= amount; j++ {
		// 初始化为math.MaxInt32
		dp[j] = math.MaxInt32
		// 遍历物品
		for i := 0; i < len(coins); i++ {
			if j >= coins[i] && dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
				//fmt.Println(dp)
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}





func main() {
	coins := []int{1, 2, 5}
	amount := 11
	print(coinChange1(coins,amount))
}