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

// 动态规划正确答案
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	// dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度
	dp := make([]int, n)
	// 初始化：每个元素自身构成长度为1的子序列
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	maxLen := 1
	// 对于每个位置i
	for i := 1; i < n; i++ {
		// 检查所有在i之前的位置j
		for j := 0; j < i; j++ {
			// 如果nums[i]可以接在nums[j]后面形成递增序列
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		// 更新全局最大长度
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

/*
算法思路
维护一个数组 tails，其中 tails[i] 表示长度为 i+1 的递增子序列的最小结尾元素。这个数组始终保持递增，因此可以使用二分查找。
算法步骤：

初始化一个空的 tails 数组。这个数组的长度将代表当前找到的最长上升子序列的长度。

遍历输入数组 nums 中的每个元素 num：

如果 num 大于 tails 数组中的所有元素（即 num > tails.back()，或者如果 tails 为空）： 将 num 追加到 tails 数组的末尾。这表示我们找到了一个更长的上升子序列。
如果 num 不大于 tails 数组中的所有元素： 在 tails 数组中，使用二分查找找到第一个大于或等于 num 的元素 tails[k]。 用 num 替换 tails[k]。 为什么这样做？ 因为我们找到了一个元素 num，它可以形成一个长度为 k+1 的上升子序列，并且这个子序列的末尾元素 num 比之前记录的 tails[k] 要小（或等于，但我们通常找第一个大于或等于的来替换，以得到更“有潜力”的末尾）。这意味着我们用一个更小的末尾元素达到了同样的长度，这为后续元素构成更长的上升子序列提供了更好的机会。
遍历完所有 nums 中的元素后，tails 数组的长度就是整个序列的最长上升子序列的长度。
*/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	// tails[i] 表示长度为 i+1 的递增子序列的最小结尾元素
	tails := make([]int, 0, n)

	for _, num := range nums {
		// 在 tails 中二分查找第一个大于等于 num 的位置
		left, right := 0, len(tails)
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 如果 num 比所有元素都大，追加到末尾
		if left == len(tails) {
			tails = append(tails, num)
		} else {
			// 否则，更新对应位置的值
			tails[left] = num
		}
	}

	return len(tails)
}

func main() {
	nuns := []int{0, 1, 0, 3, 2, 3}
	lis := lengthOfLIS(nuns)
	fmt.Printf("result:", lis)
}
