package main

import "fmt"

/*
718. 最长重复子数组
提示
给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。

示例 1：
输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
输出：3
解释：长度最长的公共子数组是 [3,2,1] 。
示例 2：
输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
输出：5

提示：

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 100
*/
// 错误代码
/*
	// 问题：于求最长公共子数组（连续）
	// 但下面的逻辑实现的是最长公共子序列（不需要连续）
*/
func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	n1 := len(nums1) + 1
	n2 := len(nums2) + 1
	result := make([][]int, n1)
	for i := 0; i < n1; i++ {
		result[i] = make([]int, n2)
	}

	for i := 1; i < n1; i++ {
		for j := 1; j < n2; j++ {
			if nums1[i-1] == nums2[j-1] {
				result[i][j] = result[i-1][j-1] + 1
			} else {
				// 错误：最长公共子数组要求连续，当元素不相等时应该重置为0
				// 而不是取 max(result[i-1][j], result[i][j-1])
				// 这行代码是最长公共子序列的逻辑
				result[i][j] = max(result[i-1][j], result[i][j-1])
			}
		}
	}
	// 错误：对于最长公共子数组，应该返回整个dp数组中的最大值
	// 而不是右下角的值
	return result[n1-1][n2-1]
}

// 正确的最长公共子数组实现应该是：
func findLengthCorrect(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}

	n1 := len(nums1)
	n2 := len(nums2)

	// dp[i][j] 表示以 nums1[i-1] 和 nums2[j-1] 结尾的最长公共子数组长度
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	// 记录最长公共子数组的长度
	maxLen := 0

	// 填充dp表
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if nums1[i-1] == nums2[j-1] {
				// 当前元素相等，最长公共子数组长度加1
				dp[i][j] = dp[i-1][j-1] + 1
				// 更新最大长度
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			} else {
				// 当前元素不相等，重置为0
				// 因为我们要找的是连续的子数组
				dp[i][j] = 0
			}
		}
	}

	return maxLen
}

func main() {
	/*
		nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
	*/
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	length := findLength(nums1, nums2)
	fmt.Printf("result, %v", length)
}
