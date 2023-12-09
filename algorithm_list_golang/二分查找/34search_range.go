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
//
//- 核心要素
//    - 注意区间开闭，三种都可以
//    - 循环结束条件：当前区间内没有元素
//    - 下一次二分查找区间：不能再查找(区间不包含)mid，防止死循环
//    - 返回值：大于等于target的第一个下标（注意循环不变量）
//
//- 有序数组中二分查找的四种类型（下面的转换仅适用于数组中都是整数）
//    1. 第一个大于等于x的下标： low_bound(x)
//    2. 第一个大于x的下标：可以转换为`第一个大于等于 x+1 的下标` ，low_bound(x+1)
//    3. 最后一个一个小于x的下标：可以转换为`第一个大于等于 x 的下标` 的`左边位置`, low_bound(x) - 1;
//    4. 最后一个小于等于x的下标：可以转换为`第一个大于等于 x+1 的下标` 的 `左边位置`, low_bound(x+1) - 1;
// 这一题求的是开始位置和结束位置，也就是大于等于和小于等于！！！！！
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