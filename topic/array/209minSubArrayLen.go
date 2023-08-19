package main

import "math"

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。



示例 1：

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
示例 2：

输入：target = 4, nums = [1,4,4]
输出：1
示例 3：

输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0


提示：

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 105


进阶：

如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
*/

/*func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	min := math.MaxInt64
	start := 0
	last := start +1
	sum := nums[0]
	for start <= last && last < n{
		num := nums[last]
		if nums[start] >= target || num >= target{
			return 1
		}
		sum += num
		if sum >= target {
			if last - start +1 < min {
				min = last - start + 1
			}
			sum = nums[start+1]

			start ++
			last = start +1
		}else{
			last ++
		}

	}
	if min == math.MaxInt64 {
		return 0 // 如果没有找到满足条件的子数组，返回0
	}
	return min
}*/

/*
上述答案存在超时问题；

接下来通过滑动窗口的思路实现：

*/

//
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	min := math.MaxInt64
	start := 0

	sum := 0
	for last := 0; last < n; last++ {
		sum += nums[last]
		for sum >= target {
			size := last - start +1
			if size < min {
				min = size
			}
			// 缩小子数组的长度，移除开头的元素
			sum -= nums[start]
			start ++
		}
	}
	if min == math.MaxInt64 {
		return 0 // 如果没有找到满足条件的子数组，返回0
	}
	return min
}



func main(){
	nums := []int{1,1,1,1,7}
	target := 7
	println(minSubArrayLen(target,nums))
}