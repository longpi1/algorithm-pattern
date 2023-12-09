package main

import (
	"math"
	"sort"
)

/*
题目 #
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:


Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
示例 1：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
示例 2：

输入：nums = [0,0,0], target = 1
输出：0
题目大意 #
给定一个数组，要求在这个数组中找出 3 个数之和离 target 最近。
*/
// 解法一 O(n^2)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := 10000000000000
	closest1 := 10000000000000.000
	for first:= 0; first <len(nums) ; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		//third := len(nums) -1

		for second := first+1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for third := len(nums) -1 ;third > second; third-- {
				temp := nums[first] + nums[second] + nums[third]
				// if temp == target {
				// 	closest1 = float64(target-temp)
				// 	closest = 0
				// 	continue
				// }
				tmp := math.Abs(float64(target-temp))

				if tmp < closest1 {
					closest1 = tmp
					closest = temp
				}
			}

		}

	}

	return closest
}

// 解法二 O(n^2)  用时时间更短，解法一 第二层遍历只需要遍历second<third就可以
func threeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < n-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			// 优化点
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				tmp := int(math.Abs(float64(sum-target)))
				if tmp < diff {
					res, diff = sum, tmp
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}

func main(){
	nums := []int{-1, 2, 1, -4}
	target := 1
	println(threeSumClosest(nums,target))
}