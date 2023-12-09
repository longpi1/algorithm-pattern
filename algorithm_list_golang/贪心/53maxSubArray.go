package main

import "math"

/*
53. 最大子数组和

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：
输入：nums = [1]
输出：1
示例 3：
输入：nums = [5,4,-1,7,8]
输出：23

*/

func maxSubArray(nums []int) int {
	n := len(nums)
	if n ==1 {
		return nums[0]
	}
	result := math.MinInt64
	for first := 0; first < n ; first++ {
		sum := nums[first]
		if sum > result {
			result = sum
		}
		for second := first +1; second < n; second++ {
			sum += nums[second]
			if sum > result {
				result = sum
			}
		}
	}

	return result
}

/*
上述答案存在超时问题
下面基于动态规划与贪心实现
贪心解法
贪心贪的是哪里呢？

如果 -2 1 在一起，计算起点的时候，一定是从 1 开始计算，因为负数只会拉低总和，这就是贪心贪的地方！

局部最优：当前“连续和”为负数的时候立刻放弃，从下一个元素重新计算“连续和”，因为负数加上下一个元素 “连续和”只会越来越小。

全局最优：选取最大“连续和”

局部最优的情况下，并记录最大的“连续和”，可以推出全局最优。

从代码角度上来讲：遍历 nums，从头开始用 count 累积，如果 count 一旦加上 nums[i]变为负数，那么就应该从 nums[i+1]开始从 0 累积 count 了，因为已经变为负数的 count，只会拖累总和。

这相当于是暴力解法中的不断调整最大子序和区间的起始位置。

那有同学问了，区间终止位置不用调整么？ 如何才能得到最大“连续和”呢？

区间的终止位置，其实就是如果 count 取到最大值了，及时记录下来了。例如如下代码：

if (count > result) result = count;
这样相当于是用 result 记录最大子序和区间和（变相的算是调整了终止位置）。
*/

func maxSubArray(nums []int) int {
	maxSum := nums[0] // 初始化最大子数组和为第一个元素
	for i := 1; i < len(nums); i++ {
		// 如果当前元素加上前一个子数组和大于当前元素本身，
		// 则将当前元素更新为当前子数组和
		// 也就是nums[i-1]  要大于0
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		// 更新最大子数组和
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}
	return maxSum // 返回最大子数组和
}




func main(){
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	print(maxSubArray(nums))
}