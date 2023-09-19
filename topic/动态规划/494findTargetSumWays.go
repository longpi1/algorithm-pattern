package main

import "math"

/*
494. 目标和
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

示例 1：
输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
示例 2：
输入：nums = [1], target = 1
输出：1

*/


func findTargetSumWays(nums []int, target int) int {
	// 获取输入数组的长度
	n := len(nums)

	// 计算输入数组中所有元素的总和
	count := 0
	for i := 0; i < n; i++ {
		count += nums[i]
	}
	// 这里需要将target转换为绝对值进行比较 ！！！ int(math.Abs((float64(x))))
	// 如果目标值的绝对值大于所有元素总和，无法达到目标值，返回0
	if int(math.Abs((float64(target)))) > count {
		return 0
	}

	// 如果目标值和所有元素总和的和是奇数，无法找到合法的组合，返回0
	if (count+target)%2 == 1 {
		return 0
	}

	// 计算目标值和所有元素总和的一半，作为背包容量
	bag := (count + target) / 2

	// 创建动态规划数组，dp[i]表示达到和为i的方法数
	dp := make([]int, bag+1)

	// 初始化dp[0]为1，表示和为0的方法数为1
	dp[0] = 1

	// 遍历输入数组中的每个元素
	for i := 0; i < n; i++ {
		// 更新dp数组，从后向前遍历，确保不重复计算
		for j := bag; j >= nums[i] ; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	// 返回达到目标值的方法数
	return dp[bag]
}


func main() {
	nums := []int{1,1,1,1,1}
	target := 3
	print(findTargetSumWays(nums, target))
}