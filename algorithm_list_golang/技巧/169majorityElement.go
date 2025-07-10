package main

import "sort"

/*

169. 多数元素
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
提示：
n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109


进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

示例 1：
输入：nums = [3,2,3]
输出：3
示例 2：
输入：nums = [2,2,1,1,1,2,2]
输出：2
*/

/*
解题思路：
此处撰写解题思路 先排序，然后返回排好序数组中间的那个数即可
*/

func majorityElement(nums []int) int {
	length := len(nums)
	// 先排序
	sort.Ints(nums)
	// 返回排好序数组中间的那个数即可
	return nums[length/2]
}

// 最优实现 使用 Boyer-Moore 投票算法实现
// 时间复杂度: O(n) - 因为我们只遍历数组一次
// 空间复杂度: O(1) - 因为我们只使用了两个额外的变量 (candidate, count)
/*
为什么这个算法有效？
因为众数的数量超过了所有其他元素数量的总和。所以，即使众数每次都和“非众数”进行“1对1消耗”，它最终也一定能剩下，成为最后留下的那个“候选人”。
*/
func majorityElement(nums []int) int {
	// 初始化候选人和计数器
	var candidate int
	count := 0

	// 遍历数组
	for _, num := range nums {
		// 如果计数器为0，说明之前的候选人已经被“抵消”完了
		// 我们选择当前元素作为新的候选人
		if count == 0 {
			candidate = num
		}

		// 判断当前元素是否是我们的候选人
		if num == candidate {
			// 是，票数+1
			count++
		} else {
			// 不是，票数-1 (同归于尽)
			count--
		}
	}

	// 因为题目保证众数总是存在，所以循环结束后，candidate 必然是众数
	return candidate
}
