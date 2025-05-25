package main

/*题目 #
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:


Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]

题目大意 #
在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。

解题思路 #
这道题最优的做法时间复杂度是 O(n)。

顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个数字的下标即可。如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果。*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		ano := target - num
		if tmp, ok := m[ano]; ok {
			return []int{i, tmp}
		}
		m[nums[i]] = i
	}
	return nil
}

func twoSum(nums []int, target int) (int, int) {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return m[another], i
		}
		m[nums[i]] = i
	}
	return 0, 0
}

func twoSum1(nums []int, target int) []int {
	tmp := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {

		tmp1 := target - nums[i]
		if _, ok := tmp[tmp1]; ok {
			tmp2 := tmp[tmp1]
			return []int{tmp2, i}
		}
		tmp[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	print(twoSum1(nums, target))
}
