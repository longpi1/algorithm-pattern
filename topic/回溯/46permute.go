package main

import "fmt"

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。



示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]


提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同

解决思路：
回溯算法：
# 回溯算法

回溯 <---->递归1.递归的下面就是回溯的过程

2.回溯法是一个 纯暴力的 搜索

3.回溯法解决的问题：

	3.1组合 如：1234  两两组合

	3.2切割问题 如：一个字符串有多少个切割方式 ，或者切割出来是回文

	3.3子集 ： 1 2 3 4  的子集

	3.4排列问题（顺序）

	3.5棋盘问题：n皇后  解数独

4.回溯可抽象成树形结构

5.模板
```go
result = []
func backtrack(选择列表,路径):
    if 满足结束条件:
        result.add(路径)
        return
    for 选择 in 选择列表:
        做选择
        backtrack(选择列表,路径)
        撤销选择
```
*/

/*func permute(nums []int) (ans [][]int) {
	n := len(nums)
	var path []int
	dfs := func() {}
	used := make(map[int]bool, n)
	dfs = func() {
		if len(path) == n {
			tmp := make([]int, n)
			tmp = path
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[i]{
				used[i] = true
				path = append(path, nums[i])
				dfs()
				used[i] = false
				path = path[:len(path)-1]
			}

		}
	}
	dfs()

	return
}*/

/*
上述代码问题：
tmp = path
在创建要添加到ans数组的path副本时，赋值tmp = path是不正确的。这不会创建path的副本，而是使tmp指向与path相同的基础数组。要修复此问题，应使用copy函数创建正确的副本：
*/

func permute(nums []int) (ans [][]int) {
	n := len(nums)
	var path []int
	dfs := func() {}
	used := make(map[int]bool, n)
	dfs = func() {
		if len(path) == n {
			tmp := make([]int, n)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if !used[i]{
				used[i] = true
				path = append(path, nums[i])
				dfs()
				used[i] = false
				path = path[:len(path)-1]
			}

		}
	}
	dfs()

	return
}



func main()  {
	nums := []int{1,2,3}
	fmt.Printf("result:%v",permute(nums))
}



