package main

import "fmt"

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。



示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]


提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同
*/

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	dfs := func(index int) {}
	dfs = func(index int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)

		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return result
}

func subsets(nums []int) [][]int {
	n := len(nums)
	var result [][]int
	var path []int
	dfs := func(int) {}
	dfs = func(index int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)

		for i := index; i < n; i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}

	}
	dfs(0)
	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Printf("result: %v", subsets(nums))
}
