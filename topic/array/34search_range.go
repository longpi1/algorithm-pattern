package main

import (
	"fmt"
	"sort"
)

/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。



示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]


提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109
*/

func search(nums []int, target int) int {
	left, right := 0, len(nums) -1
	for left  <= right {
		mid := left + (right - left)/2
		if nums[mid] < target {
			left = mid + 1
		}else  {
			right = mid -1
		}
	}
	return left
}

func searchRange(nums []int, target int) []int {
	leftIndex := search(nums,target)
	if leftIndex == len(nums) || nums[leftIndex] != target {
		return []int{-1,-1}
	}
	rightIndex := search(nums,target + 1) -1
	return []int{leftIndex,rightIndex}
}



//
//作者：力扣官方题解
//链接：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/solutions/504484/zai-pai-xu-shu-zu-zhong-cha-zhao-yuan-su-de-di-3-4/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main()  {
	nums := []int{5,7,7,8,8,10}
	target := 8
	fmt.Printf("result:%v",searchRange(nums,target))
}