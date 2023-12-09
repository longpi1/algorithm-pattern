package main

/*
560. 和为 K 的子数组

给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。

子数组是数组中元素的连续非空序列。

示例 1：
输入：nums = [1,1,1], k = 2
输出：2
示例 2：
输入：nums = [1,2,3], k = 3
输出：2
*/


func subarraySum(nums []int, k int) int {
	count := 0 // 用于记录符合条件的子数组数量

	for start := 0; start < len(nums); start++ { // 外层循环遍历所有可能的子数组起始位置
		sum := 0 // 用于记录当前子数组的累计和

		// 内层循环逆向遍历子数组
		for end := start; end >= 0; end-- { // 从起始位置往前遍历子数组
			sum += nums[end] // 累计当前子数组的元素和

			if sum == k { // 如果当前子数组的元素和等于目标值 k
				count++ // 增加符合条件的子数组数量
			}
		}
	}

	return count // 返回符合条件的子数组数量
}

func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

