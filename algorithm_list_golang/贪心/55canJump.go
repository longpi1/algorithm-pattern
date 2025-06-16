package main

//给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
// 示例 1：
//输入：nums = [2,3,1,1,4]
//输出：true
//解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

// 示例 2：
//输入：nums = [3,2,1,0,4]
//输出：false
//解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
// 提示：
// 1 <= nums.length <= 104
// 0 <= nums[i] <= 105
//
// Related Topics 贪心 数组 动态规划

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	target := len(nums) - 1
	step := 0
	for i := 0; i <= step; i++ {
		if step < i+nums[i] {
			step = i + nums[i]
		}
		if step >= target {
			return true
		}
	}
	return false
}

//func canJump(nums []int) bool {
//	n := len(nums)
//	for i := 0; i < n; i++ {
//		if i >0 && nums[i] == nums[i-1] {
//			continue
//		}
//	}
//	return false
//}
//
//func jump(num int, l int) bool{
//
//}

// 此题未做出来

/*
解题思路：
这道题是典型的贪心算法，通过局部最优解得到全局最优解。以下两种方法都是使用贪心算法实现，只是贪心的策略不同。
这道题目关键点在于：不用拘泥于每次究竟跳几步，而是看覆盖范围，覆盖范围内一定是可以跳过来的，不用管是怎么跳的。

大家可以看出思路想出来了，代码还是非常简单的。

一些同学可能感觉，我在讲贪心系列的时候，题目和题目之间貌似没有什么联系？

是真的就是没什么联系，因为贪心无套路！没有个整体的贪心框架解决一系列问题，只能是接触各种类型的题目锻炼自己的贪心思维！
*/
func canJump(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true // 如果数组只有一个元素，直接返回 true，无需跳跃
	}

	// 初始化覆盖范围
	cover := 0

	for i := 0; i <= cover; i++ {
		// 更新覆盖范围，选择能够跳到更远位置的情况
		if cover < i+nums[i] {
			cover = i + nums[i]
		}

		// 如果已经可以覆盖数组的最后一个位置，返回 true
		if cover >= n-1 {
			return true
		}
	}

	return false // 无法跳跃到数组的最后一个位置，返回 false
}

func main() {
	nums := []int{3, 0, 8, 2, 0, 0, 1}
	canJump(nums)
}
