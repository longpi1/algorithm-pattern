package main

import "fmt"

/*
300. 最长递增子序列
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
示例 1：
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
示例 2：
输入：nums = [0,1,0,3,2,3]
输出：4
示例 3：
输入：nums = [7,7,7,7,7,7,7]
输出：1

提示：
1 <= nums.length <= 2500
-104 <= nums[i] <= 104
进阶：
你能将算法的时间复杂度降低到 O(n log(n)) 吗?
*/
// 第一版错误答案
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	// 错误1: 使用二维数组是不必要的，最长递增子序列只需要一维dp数组
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		// 错误2: 变量名min容易误导，实际上应该记录当前子序列的最大值
		min := nums[i]
		for j := i + 1; j < n; j++ {
			// 错误3: 逻辑错误，应该是 nums[j] > 前一个元素，而不是 > min
			// 错误4: min在这里应该记录子序列的最大值，但更新逻辑有问题
			if nums[j] > min {
				min = nums[j]
				fmt.Println("result:", dp[i][j-1])
				// 错误5: dp[i][j]的定义不明确，不符合LIS问题的标准定义
				// 错误6: 状态转移方程错误，没有考虑之前的所有可能子序列
				dp[i][j] = max(dp[i][j], dp[i][j-1]+1)
				fmt.Println("result1:", dp[i][j])
			}
			// 错误7: 当nums[j] <= min时，dp[i][j]没有被赋值，保持为0
		}
	}
	// 错误8: 返回值错误，dp[n-1][n-1]不代表整个数组的最长递增子序列长度
	// 错误9: 没有遍历所有可能的结果来找最大值
	return dp[n-1][n-1]
}

func main() {
	nuns := []int{0, 1, 0, 3, 2, 3}
	lis := lengthOfLIS(nuns)
	fmt.Printf("result:", lis)
}
