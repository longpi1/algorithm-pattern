package main

/*
41. 缺失的第一个正数

给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。


示例 1：
输入：nums = [1,2,0]
输出：3
示例 2：
输入：nums = [3,4,-1,1]
输出：2
示例 3：
输入：nums = [7,8,9,11,12]
输出：1
*/


func firstMissingPositive(nums []int) int {
	n := len(nums) // 获取数组的长度

	// 第一次遍历数组，将每个正整数尽可能地放在正确的位置上
	for i := 0; i < n; i++ {
		// 当前元素 nums[i] 大于 0、小于等于数组长度，并且不在正确的位置上时
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// 交换当前元素和应该在的位置上的元素
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	// 第二次遍历数组，查找第一个不在正确位置上的正整数
	for i := 0; i < n; i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}

	// 如果整个数组都在正确的位置上，返回下一个正整数
	return n + 1
}
