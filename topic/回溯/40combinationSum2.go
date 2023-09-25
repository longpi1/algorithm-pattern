package main

import (
	"fmt"
	"sort"
)

/*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。

示例 1:
输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]


提示:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30

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



/*func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	var path []int
	n := len(candidates)
	used := make(map[int]bool, n)
	dfs := func(int) {}
	dfs = func(tmp int) {
		if tmp < 0 {
			return
		}
		if tmp == 0{
			var tmp []int
			copy(tmp, path)
			ans = append(ans, tmp)
		}
		for i := 0; i < n; i++ {
			if !used[i] {
				tmp = tmp - candidates[i]
				path = append(path, candidates[i])
				used[i] = true
				dfs(tmp)
				used[i] = false
				tmp = tmp + candidates[i]
				path = path[:len(path)-1]
			}
		}
	}
	dfs(target)
	return ans
}*/
/*
上述代码问题：
	1.var tmp []int
在创建要添加到ans数组的tmp副本时，赋值var tmp []int并随后的copy(tmp, path)是不正确的。你不能重新声明同名的变量，并且在调用copy之前，应该使用make创建一个新的切片。另外，tmp应该是一个局部变量。以下是修复后的代码：
2.要避免重复寻找，for循环i应该通过变量传递，也就是startIndex,for i := 也就是startIndex; i < n ; i++ {错误；

3.去重操作 used[i - 1] == false，说明同一树层candidates[i - 1]使用过，切记是树层

*/

func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	var path []int
	sort.Ints(candidates)
	n := len(candidates)
	used := make(map[int]bool, n)
	dfs := func(int,int) {}
	dfs = func(tmp int, index int) {
		if tmp < 0 {
			return
		}
		if tmp == 0{
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
		}
		for i := index; i < n; i++ {
			if !used[i] {
				// 去重操作
				// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
				// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
				// 一个是树枝一个是树层！！！
				if i > 0 && candidates[i] == candidates[i-1]  && used[i-1] == false {
					continue
				}
				tmp = tmp - candidates[i]
				path = append(path, candidates[i])

				used[i] = true
				dfs(tmp,i)
				used[i] = false
				tmp = tmp + candidates[i]
				path = path[:len(path)-1]
			}
		}
	}

	dfs(target, 0)
	return ans
}

func main()  {
	candidates := []int{2,5,2,1,2}
	target := 5
	fmt.Printf("result:%v",combinationSum2(candidates,target))
}
