package main

import "fmt"

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
// 错误代码

func CompleteBackpack(n, m int, weights, values []int) int {
	// 问题1: result 初始化了外层切片，但长度是 n+1，而循环是从 i=1 到 n-1。
	// 如果 n 是物品种类的数量，那么 weights 和 values 的有效索引是 0 到 n-1。
	// dp 表的行通常对应物品，所以 dp[i] 的 i 应该能对应到 weights[i-1] 或 weights[i]。
	// dp 表的列通常对应容量，所以 m 如果是最大容量，需要 m+1 列。
	result := make([][]int, n+1) // 外层分配了 n+1 个 nil 切片

	// 问题2: 外层循环范围 `for i := 1; i < n; i++`
	// 这意味着它只考虑了物品索引从 1 到 n-1 (如果 weights/values 是0-indexed)。
	// 如果有 n 种物品，索引应该是 0 到 n-1。
	// 如果 dp[i] 表示考虑前 i 种物品，那么 i 应该从 1 到 n。
	// 当前循环会漏掉第一个物品 (索引0) 和最后一个物品 (如果 dp 表设计为 dp[n] 存储结果)。
	// 并且，result[0] 这一行没有被初始化 `make([]int, m)`。
	for i := 1; i < n; i++ {
		// 问题3: 内层切片初始化 `result[i] = make([]int, m)`
		// 如果 m 是最大容量，列数应该是 m+1 (容量0到m)。这里是 m 列 (容量0到m-1)。
		result[i] = make([]int, m)

		for j := 0; j < m; j++ { // j 代表当前容量，从0到m-1
			// 问题4: 数组越界 `weights[i]`
			// 如果 i 的设计是代表 "第i种物品" (1-indexed for dp table)，那么对应的
			// weights 数组索引应该是 i-1。当前代码直接用 i，可能越界 (当i=n时)或不对应。
			// 假设 i 是 weights/values 的直接索引 (0-indexed for items):
			// 此时 weights[i] 是正确的，但 i 循环范围是 1 to n-1，意味着物品0被跳过了。

			// 问题5: 状态转移方程错误
			// if weights[i] > j: (物品i的重量大于当前容量j，放不下)
			//   result[i][j] = result[i][j-1]  <--- 这是错误的。
			//   如果放不下当前物品，价值应该继承自 "不考虑当前物品" 或 "考虑了当前物品但当前容量太小"。
			//   对于完全背包的二维DP，通常是: result[i][j] = result[i-1][j] (不选第i种物品)
			//   或者如果j=0，result[i][0]=0。
			//   result[i][j-1] 的含义是“考虑第i种物品，容量为j-1时的最大价值”，这不适用于“放不下”的情况。
			if weights[i] > j {
				result[i][j] = result[i][j-1] // 逻辑错误
			} else {
				// 问题6: 应该加 values[i-1]
				result[i][j] = max(result[i][j-1], result[i][j-weights[i]]+values[i]) // 逻辑和价值错误
			}
		}
	}
	// 问题6: 返回值索引 `result[n][m]`
	// 由于循环 i 只到 n-1，result[n] 这一行根本没有被计算和填充。
	// 并且 result[...][m] 也是越界的，因为内层切片长度是 m。
	return result[n][m]
}

// 正确代码

func CompleteBackpackFixed(numItemTypes, capacity int, weights, values []int) int {
	if numItemTypes == 0 || capacity == 0 {
		return 0
	}

	// dp[i][j] 表示从前 i 种物品中选择 (每种可多次)，放入容量为 j 的背包中的最大价值
	// 物品的索引在 weights/values 中是 0 到 numItemTypes-1
	// dp 表的 i 维度从 0 (无物品)到 numItemTypes (所有种类的物品)
	// dp 表的 j 维度从 0 (容量0)到 capacity (最大容量)
	dp := make([][]int, numItemTypes+1)
	for k := range dp {
		dp[k] = make([]int, capacity+1)
		// dp[k][0] 默认为 0，表示容量为0时价值为0
	}
	// dp[0][j] 默认为 0，表示没有物品可选时价值为0

	// 遍历物品种类 (从第1种物品到第numItemTypes种物品)
	// i 代表考虑的是 "前i种物品"
	for i := 1; i <= numItemTypes; i++ {
		// 当前考虑的物品是第 i 种，其在原始数组中的索引是 i-1
		itemWeight := weights[i-1]
		itemValue := values[i-1]

		// 遍历背包容量
		for j := 0; j <= capacity; j++ { // j 从0开始
			// 选项1: 不选择第 i 种物品 (即物品 weights[i-1])
			// 价值等于只考虑前 i-1 种物品，放入容量为 j 的背包的价值
			dp[i][j] = dp[i-1][j]

			// 选项2: 选择至少一个第 i 种物品 (前提是当前容量 j 能装下它)
			if j >= itemWeight {
				// 价值为：(放入一个第i种物品的价值) +
				//         (从前i种物品中选择，放入剩余容量 j-itemWeight 的价值)
				// 注意这里是 dp[i][j-itemWeight]，因为我们仍然可以继续选择第i种物品
				// 与不选择第 i 种物品的情况比较，取较大者
				dp[i][j] = max(dp[i][j], dp[i][j-itemWeight]+itemValue)
			}
		}
	}

	// 最终结果是考虑了所有 numItemTypes 种物品，在最大容量 capacity 下的最大价值
	return dp[numItemTypes][capacity]
}
func main() { // Combined main
	// --- 0/1 Knapsack Example ---
	weights01 := []int{1, 3, 4}
	values01 := []int{15, 20, 30}
	capacity01 := 4

	fmt.Println("--- 0/1 Knapsack ---")
	fmt.Println("0/1 Knapsack (2D DP):")
	maxValue2D_01 := backpack(capacity01, len(weights01), weights01, values01)
	fmt.Printf("Max value: %d\n", maxValue2D_01) // Expected: 35

	fmt.Println("\n0/1 Knapsack (1D DP - Space Optimized):")
	maxValue1D_01 := backpack(capacity01, len(weights01), weights01, values01)
	fmt.Printf("Max value: %d\n", maxValue1D_01) // Expected: 35
	fmt.Println("----------------------")

	// --- Complete Knapsack Example ---
	weightsComplete := []int{1, 3, 4}   // 物品重量
	valuesComplete := []int{15, 20, 30} // 物品价值

	fmt.Println("\n--- Complete Knapsack ---")
	fmt.Println("Complete Knapsack (2D DP):")
	maxValue2DComplete := CompleteBackpackFixed(capacity01, len(weights01), weightsComplete, valuesComplete)
	// Possible combinations for capacity 4:
	// - 4x item1 (weight 1, value 15) -> total value 4*15 = 60
	// - 1x item2 (weight 3, value 20) + 1x item1 (weight 1, value 15) -> total value 20 + 15 = 35
	// - 1x item3 (weight 4, value 30) -> total value 30
	fmt.Printf("Max value: %d\n", maxValue2DComplete) // Expected: 60 (4 * item1)

	fmt.Println("\nComplete Knapsack (1D DP - Space Optimized):")
	maxValue1DComplete := CompleteBackpackFixed(capacity01, len(weights01), weightsComplete, valuesComplete)
	fmt.Printf("Max value: %d\n", maxValue1DComplete) // Expected: 60
	fmt.Println("-------------------------")
}
