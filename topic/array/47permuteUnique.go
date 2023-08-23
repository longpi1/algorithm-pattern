package main

import (
	"fmt"
	"sort"
)

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。



示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]


提示：

1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/


func permuteUnique(nums []int) [][]int {
	var result [][]int
	var path []int
	// 题目中给的不一定有序，首先进行排序，否则nums[i-1] == nums[i]判断不成立，无法去重
	sort.Ints(nums)
	n := len(nums)
	used := make(map[int]bool, n)
	dfs := func() {}
	dfs = func() {
		if len(path) == n{
			tmp := make([]int,n)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := 0; i < n ; i++ {
			if !used[i]{
				// 去重操作
				// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
				// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
				// 一个是树枝一个是树层！！！
				if i>0 && nums[i-1] == nums[i] && used[i-1] == false{
					continue
				}
				used[i] = true
				path = append(path, nums[i])
				dfs()
				used[i] = false
				path = path[:len(path)-1]
			}
		}

	}
	dfs()
	return result
}


func main(){
	nums := []int{3,3,0,3}
	fmt.Printf("result:%v",permuteUnique(nums))
}

