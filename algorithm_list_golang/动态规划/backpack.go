package main

/*
背包问题集合
一、0/1 背包问题
问题描述
有 N 件物品和一个容量为 W 的背包。第 i 件物品的重量是 weights[i]，价值是 values[i]。求解将哪些物品装入背包，可使这些物品的重量总和不超过背包容量，且价值总和最大。、
特点：每件物品只有一件，要么装入背包，要么不装入。

二、完全背包问题
问题描述
有 N 种物品和一个容量为 W 的背包。第 i 种物品的重量是 weights[i]，价值是 values[i]。求解将哪些物品装入背包，可使这些物品的重量总和不超过背包容量，且价值总和最大。
特点：每种物品都有无限件可用。
*/

/*

一、0/1 背包问题
*/
// 下述代码存在问题如下：
func backpack(n, m int, weights, values []int) int {
	result := make([][]int, 0) // 问题1: 创建了长度为0的切片
	for i := 0; i < n; i++ {
		result[i] = make([]int, m) // 问题2: 会导致panic，因为result长度为0
		for j := 0; j < m; j++ {
			if weights[i] < j { // 问题3: 条件判断错误，应该是 weights[i] > j
				result[i][j] = result[i-1][j] // 问题4: 当i=0时会越界
			} else {
				result[i][j] = max(result[i-1][j], result[i-1][j-weights[i]]+values[i])
				// 问题4: 当i=0时会越界
				// 问题5: 当j-weights[i]<0时会越界
			}
		}
	}
	return result[n][m] // 问题6: 数组越界，应该是result[n-1][m-1]
}

// 正确逻辑
func backpackFixed(numItems, capacity int, weights, values []int) int {
	if numItems == 0 || capacity == 0 {
		return 0
	}
	// dp[i][j] 表示从前 i 个物品中选取，放入容量为 j 的背包中的最大价值
	// 物品的索引在 weights/values 中是 0 到 numItems-1
	// dp 表的 i 维度从 0 (无物品)到 numItems (所有物品)
	// dp 表的 j 维度从 0 (容量0)到 capacity (最大容量)
	dp := make([][]int, numItems+1)
	for k := range dp {
		dp[k] = make([]int, capacity+1)
	}

	// 遍历物品 (从第1个物品到第numItems个物品)
	// i 代表考虑的是 "前i个物品"
	for i := 1; i <= numItems; i++ {
		// 当前考虑的物品是第 i 个，其在原始数组中的索引是 i-1
		itemWeight := weights[i-1]
		itemValue := values[i-1]

		// 遍历背包容量
		for j := 0; j <= capacity; j++ { // j 从0开始，因为容量为0也是一种状态
			// 不放入第 i 个物品 (即物品 weights[i-1])
			// 此时价值等于只考虑前 i-1 个物品，放入容量为 j 的背包的价值
			dp[i][j] = dp[i-1][j]

			// 如果当前背包容量 j 能够装下第 i 个物品
			if j >= itemWeight {
				// 我们可以选择放入第 i 个物品
				// 价值为：(放入第i个物品的价值) + (前i-1个物品放入剩余容量 j-itemWeight 的价值)
				// 与不放入第 i 个物品的情况比较，取较大者
				dp[i][j] = max(dp[i][j], dp[i-1][j-itemWeight]+itemValue)
			}
		}
	}

	// 最终结果是考虑了所有 numItems 个物品，在最大容量 capacity 下的最大价值
	return dp[numItems][capacity]
}

// ZeroOneKnapsack1D 使用一维DP (空间优化) 解决0/1背包问题
func ZeroOneKnapsack1D(weights []int, values []int, capacity int) int {
	numItems := len(weights)
	if numItems == 0 || capacity == 0 {
		return 0
	}

	// dp[j] 表示容量为 j 的背包能获得的最大价值
	dp := make([]int, capacity+1)

	// 遍历物品
	for i := 0; i < numItems; i++ {
		weight := weights[i]
		value := values[i]
		// 遍历背包容量 (必须倒序)
		// 确保每个物品只被放入一次。
		// 如果正序，dp[j-weight] 可能已经包含了当前物品 i，导致重复计算。
		for j := capacity; j >= weight; j-- {
			// 对于容量 j，我们可以选择不放入物品 i (价值为 dp[j] - 上一轮的值)
			// 或者放入物品 i (价值为 dp[j-weight] + value - dp[j-weight] 也是上一轮的值)
			dp[j] = max(dp[j], dp[j-weight]+value)
		}
	}
	return dp[capacity]
}

/*
二、完全背包问题
*/
func CompleteBackpack(n, m int, weights, values []int) int {

}
