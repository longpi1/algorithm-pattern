package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
题目 #
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:


Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

题目大意 #
给定一个数组，要求在这个数组中找出 3 个数之和为 0 的所有组合。
解题思路 #
用 map 提前计算好任意 2 个数字之和，保存起来，可以将时间复杂度降到 O(n^2)。这一题比较麻烦的一点在于，最后输出解的时候，要求输出不重复的解。数组中同一个数字可能出现多次，同一个数字也可能使用多次，但是最后输出解的时候
，不能重复。例如 [-1，-1，2] 和 [2, -1, -1]、[-1, 2, -1] 这 3 个解是重复的，即使 -1 可能出现 100 次，每次使用的 -1 的数组下标都是不同的。

这里就需要去重和排序了。map 记录每个数字出现的次数，然后对 map 的 key 数组进行排序，最后在这个排序以后的数组里面扫，找到另外 2 个数字能和自己组成 0 的组合。
*/


func containsList(lists [][]int, newList []int) bool {
	for _, existingList := range lists {
		if reflect.DeepEqual(existingList, newList) {
			return true
		}
	}
	return false
}

//
//func threeSum(nums []int) [][]int {
//	var result [][]int
//
//	tmp := make(map[int]int,len(nums))
//	for i := 0; i < len(nums) -1; i++ {
//		tmp[nums[i]] = i
//	}
//	for i := 0; i < len(nums) ; i++ {
//		for j := i+1; j < len(nums) ; j++ {
//			ano := 0 - nums[i] - nums[j]
//			flag := j != tmp[ano]  && i != tmp[ano]
//			_, ok := tmp[ano]
//			if flag && ok {
//				tmp1 := []int{nums[i],nums[j],ano}
//				sort.Ints(tmp1)
//				if !containsList(result,tmp1){
//					result = append(result, tmp1)
//				}
//
//
//			}
//		}
//	}
//	return result
//}
//
//// 上述答案超时
//
//// 优化后的答案
//// 排序+双指针
//func threeSum(nums []int) [][]int {
//	n := len(nums)
//	sort.Ints(nums)
//	ans := make([][]int, 0)
//
//	// 枚举 a
//	for first := 0; first < n; first++ {
//		// 需要和上一次枚举的数不相同
//		if first > 0 && nums[first] == nums[first - 1] {
//			continue
//		}
//		// c 对应的指针初始指向数组的最右端
//		third := n - 1
//		target := -1 * nums[first]
//		// 枚举 b
//		for second := first + 1; second < n; second++ {
//			// 需要和上一次枚举的数不相同
//			if second > first + 1 && nums[second] == nums[second - 1] {
//				continue
//			}
//			// 需要保证 b 的指针在 c 的指针的左侧
//			for second < third && nums[second] + nums[third] > target {
//				third--
//			}
//			// 如果指针重合，随着 b 后续的增加
//			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
//			if second == third {
//				break
//			}
//			if nums[second] + nums[third] == target {
//				ans = append(ans, []int{nums[first], nums[second], nums[third]})
//			}
//		}
//	}
//	return ans
//}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i<len(nums); i++{
		if i > 0 && nums[i] == nums[i-1] {
			break
		}
		first := i
		third := len(nums)-1
		for second := i+1; second<len(nums); second++{
			// 需要和上一次枚举的数不相同
			if second > first +1 &&  nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second] + nums[third]+ nums[first] > 0 {
				third--
			}
			if nums[first] + nums[second] +nums[third] == 0{
				// 如果指针重合，随着 b 后续的增加
				// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
				if second == third {
					break
				}
				temp := []int{nums[first],nums[second],nums[third]}
				result = append(result,temp)
			}

		}
	}
	return result
}

func main(){
	nums := []int{-1,0,1,2,-1,-4,-2,-3,3,0,4}
	fmt.Printf("result: %v",threeSum(nums))
}