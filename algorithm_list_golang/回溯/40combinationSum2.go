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
/*

下述代码问题：
问题 1：used 的初始化和使用没有考虑重复元素的情况。如果数组中有重复元素，used[i] 只能标记某个具体索引的元素是否被使用，而无法区分相同值的不同元素。

问题 2：used 的作用在当前代码中可以被优化掉，因为 combinationSum2 问题中每个元素只能使用一次，且我们通过 index 参数已经保证了不会重复使用之前的元素
问题原因：在遍历 candidates 时，没有跳过重复的元素。

解决方法：在遍历时，如果当前元素和前一个元素相同（candidates[i] == candidates[i-1]），并且前一个元素未被使用，则跳过当前元素。为了实现这一点，需要对 candidates 进行排序。

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]bool)
	dfsCombinationSum := func(path []int, index int, needNum int) {}
	dfsCombinationSum = func(path []int, index int, needNum int) {
		if needNum == 0 {
			combination := make([]int, len(path))
			copy(combination, path)
			result = append(result, combination)
			return
		}
		for i := index; i < len(candidates); i++ {
			if candidates[i] <= needNum && !used[i] {
				used[i] = true
				needNum = needNum - candidates[i]
				path = append(path, candidates[i])
				dfsCombinationSum(path, index+1, needNum)
				used[i] = false
				needNum = needNum + candidates[i]
				path = path[:len(path)-1]
			}
		}
	}
	dfsCombinationSum(path, 0, target)
	return result
}*/

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)

	// 排序：通过 sort.Ints(candidates) 对输入数组进行排序，便于剪枝和处理重复元素。
	sort.Ints(candidates)

	var dfsCombinationSum func(index int, needNum int)
	dfsCombinationSum = func(index int, needNum int) {
		// 找到一个有效组合
		if needNum == 0 {
			combination := make([]int, len(path))
			copy(combination, path)
			result = append(result, combination)
			return
		}

		// 从 index 开始遍历 candidates
		for i := index; i < len(candidates); i++ {
			// 剪枝：如果当前元素大于需要的目标值，直接跳出循环
			if candidates[i] > needNum {
				break
			}

			// 跳过重复元素：如果当前元素和前一个元素相同，则跳过
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}

			// 做选择：将当前元素加入路径
			path = append(path, candidates[i])

			// 递归：从下一个元素开始，继续寻找组合
			dfsCombinationSum(i+1, needNum-candidates[i])

			// 撤销选择：回溯
			path = path[:len(path)-1]
		}
	}

	dfsCombinationSum(0, target)
	return result
}

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

//func combinationSum2(candidates []int, target int) [][]int {
//	var ans [][]int
//	var path []int
//	sort.Ints(candidates)
//	n := len(candidates)
//	used := make(map[int]bool, n)
//	dfs := func(int, int) {}
//	dfs = func(tmp int, index int) {
//		if tmp < 0 {
//			return
//		}
//		if tmp == 0 {
//			tmp := make([]int, len(path))
//			copy(tmp, path)
//			ans = append(ans, tmp)
//		}
//		for i := index; i < n; i++ {
//			if !used[i] {
//				// 去重操作
//				// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
//				// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
//				// 一个是树枝一个是树层！！！
//				if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
//					continue
//				}
//				tmp = tmp - candidates[i]
//				path = append(path, candidates[i])
//
//				used[i] = true
//				dfs(tmp, i)
//				used[i] = false
//				tmp = tmp + candidates[i]
//				path = path[:len(path)-1]
//			}
//		}
//	}
//
//	dfs(target, 0)
//	return ans
//}

func main() {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	fmt.Printf("result:%v", combinationSum2(candidates, target))
}
