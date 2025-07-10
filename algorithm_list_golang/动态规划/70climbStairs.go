package main

/*
70. 爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

示例 1：
输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶
*/

// 第一版存在的问题
/*
BUG：存在大量重复计算，导致时间复杂度过高（指数级）
这是典型的重叠子问题 (Overlapping Subproblems) 现象，也是动态规划问题的一个标志。
例如，要计算 climbStairs(5)：
climbStairs(5) = climbStairs(4) + climbStairs(3)
climbStairs(4) = climbStairs(3) + climbStairs(2)
climbStairs(3) = climbStairs(2) + climbStairs(1)
你会发现 climbStairs(3)、climbStairs(2)、climbStairs(1) 等函数会被多次重复调用，每次都从头开始计算，而不是利用之前已经计算过的结果。
这种纯递归的时间复杂度是 O(2^n)，会随着 n 的增大呈指数级增长，很快就会导致栈溢出或计算超时。对于 LeetCode 等在线判题系统，当 n 超过 40 左右时，通常就会超时。
*/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
func climbStairs(n int) int {
	// 基础情况 (Base Cases):
	// 如果 n <= 3，直接返回 n。
	// 这包含了 n=1 (返回1种), n=2 (返回2种), n=3 (返回3种)
	// 对于 n=1: 1种 (1)
	// 对于 n=2: 2种 (1+1, 2)
	// 对于 n=3: 3种 (1+1+1, 1+2, 2+1) -- 注意这里 n=3 返回 3 是对的
	// 所以，这个基础情况处理是正确的。
	if n <= 3 {
		return n
	}

	// 递归关系 (Recurrence Relation):
	// 爬到 n 级台阶的方法数 = 爬到 n-1 级台阶的方法数 + 爬到 n-2 级台阶的方法数。
	// 这是正确的斐波那契数列模式。
	return climbStairs(n-1) + climbStairs(n-2) // BUG: 存在大量重复计算，导致时间复杂度过高（指数级）
}

/*
优化方案：动态规划 (记忆化搜索 或 迭代)
解决这种重复计算问题，正是动态规划（或记忆化搜索）的用武之地。

优化方案一：记忆化搜索 (Top-down with Memoization)
使用一个缓存（例如哈希表或数组）来存储已经计算过的 climbStairs(k) 的结果。在每次计算前，先检查缓存中是否已有结果；如果没有，则计算并存入缓存。
*/

var memo map[int]int // 全局或传入参数的缓存

func climbStairsMemo(n int) int {
	// 初始化缓存
	if memo == nil {
		memo = make(map[int]int)
	}

	// 1. 检查缓存：如果已经计算过，直接返回结果，避免重复计算。
	if val, ok := memo[n]; ok {
		return val
	}

	// 2. 基础情况 (Base Cases):
	if n <= 3 {
		memo[n] = n // 将基础情况的结果也存入缓存
		return n
	}

	// 3. 递归关系 (Recurrence Relation):
	// 计算结果并存入缓存
	result := climbStairsMemo(n-1) + climbStairsMemo(n-2)
	memo[n] = result
	return result
}

/*
// 如果不想使用全局变量，可以将 memo 作为参数传递
func climbStairsMemoClean(n int, memo map[int]int) int {
    if val, ok := memo[n]; ok {
        return val
    }
    if n <= 3 {
        memo[n] = n
        return n
    }
    result := climbStairsMemoClean(n-1, memo) + climbStairsMemoClean(n-2, memo)
    memo[n] = result
    return result
}
*/
//
//优化方案二：迭代 (Bottom-up Dynamic Programming)
//从基础情况（小问题）开始，迭代计算出大问题的结果。通常使用一个数组来存储中间结果。

func climbStairsDP(n int) int {
	// 1. 基础情况 (Base Cases):
	if n <= 3 {
		return n
	}

	// 2. 创建 DP 数组：dp[i] 表示爬到 i 级台阶的方法数
	// 数组长度为 n+1，因为我们需要 dp[n]
	dp := make([]int, n+1)

	// 3. 初始化 DP 数组的基础值
	dp[1] = 1 // 爬 1 级台阶有 1 种方法
	dp[2] = 2 // 爬 2 级台阶有 2 种方法
	dp[3] = 3 // 爬 3 级台阶有 3 种方法

	// 4. 迭代计算：从 dp[4] 开始，根据状态转移方程计算
	// 状态转移方程：dp[i] = dp[i-1] + dp[i-2]
	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	// 5. 返回最终结果
	return dp[n]
}

/*
总结
原始的纯递归 climbStairs 代码在功能上是正确的，但由于其指数级的时间复杂度（O(2^n)），存在严重的性能问题。这是由于重叠子问题导致的重复计算。

通过引入记忆化搜索（Top-down DP）或迭代动态规划（Bottom-up DP），可以将时间复杂度优化到 O(n)，大大提高了算法效率，使其在面对较大 n 值时也能快速得出结果。迭代动态规划通常在空间复杂度上也有优势（O(n)），甚至可以进一步优化到 O(1) 空间复杂度，因为它只依赖于前两个值。
*/
