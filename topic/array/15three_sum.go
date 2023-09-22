package main

import (
	"reflect"
	"sort"
)

/*
15. 三数之和

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：
输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
示例 3：

输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。
题目大意 #
给定一个数组，要求在这个数组中找出 3 个数之和为 0 的所有组合。
解题思路 #
方法一：排序 + 双指针
题目中要求找到所有「不重复」且和为 000 的三元组，这个「不重复」的要求使得我们无法简单地使用三重循环枚举所有的三元组。这是因为在最坏的情况下，数组中的元素全部为 000，即

[0, 0, 0, 0, 0, ..., 0, 0, 0]
任意一个三元组的和都为 000。如果我们直接使用三重循环枚举三元组，会得到 O(N3)O(N^3)O(N
3
 ) 个满足题目要求的三元组（其中 NNN 是数组的长度）时间复杂度至少为 O(N3)O(N^3)O(N
3
 )。在这之后，我们还需要使用哈希表进行去重操作，得到不包含重复三元组的最终答案，又消耗了大量的空间。这个做法的时间复杂度和空间复杂度都很高，因此我们要换一种思路来考虑这个问题。

「不重复」的本质是什么？我们保持三重循环的大框架不变，只需要保证：

第二重循环枚举到的元素不小于当前第一重循环枚举到的元素；

第三重循环枚举到的元素不小于当前第二重循环枚举到的元素。

*/
func threeSum(nums []int) [][]int {
	result := make([][]int, 0) // 用于存储结果的切片
	n := len(nums)
	sort.Ints(nums) // 对数组进行升序排序

	for i := 0; i < n; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue // 避免重复处理相同的数字
		}
		left, right := i+1, n-1 // 左右指针分别指向当前数字后面的第一个和最后一个数字
		for left < right {
			if right < n-1 && nums[right] == nums[right+1] {
				right--
				continue // 避免重复处理相同的数字
			}
			if left > i+1 && nums[left-1] == nums[left] {
				left++
				continue // 避免重复处理相同的数字
			}
			if 0-nums[i]-nums[left] == nums[right] {
				result = append(result, []int{nums[i], nums[left], nums[right]})
			}
			if 0-nums[i]-nums[left] > nums[right] {
				left++
			} else {
				right--
			}
		}
	}

	return result
}


func main(){
	nums := []int{0,0,0,0}
	threeSum(nums)
}





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

func threeSum1(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i<len(nums); i++{
		if i > 0 && nums[i] == nums[i-1] {
			continue
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
			if nums[first] + nums[second] + nums[third] == 0{
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
//
//func main(){
//	nums := []int{34,55,79,28,46,33,2,48,31,-3,84,71,52,-3,93,15,21,-43,57,-6,86,56,94,74,83,-14,28,-66,46,-49,62,-11,43,65,77,12,47,61,26,1,13,29,55,-82,76,26,15,-29,36,-29,10,-70,69,17,49}
//	fmt.Printf("result: %v",threeSum(nums))
//}