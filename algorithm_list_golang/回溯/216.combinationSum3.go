package main

import "fmt"

/*
216. 组合总和 III
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：

只使用数字1到9
每个数字 最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。

示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
解释:
1 + 2 + 4 = 7
没有其他符合的组合了。
示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
解释:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
没有其他符合的组合了。
示例 3:

输入: k = 4, n = 1
输出: []
解释: 不存在有效的组合。
在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。
*/
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	dfs := func(path []int, index int, needNum int) {}
	dfs = func(path []int, index int, needNum int) {

		if len(path) == k && needNum == 0 {
			tmp := make([]int, k)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		//if len(path) == k || needNum < 0 {
		//	return
		//}
		for i := index; i <= 9; i++ {
			if i <= needNum && len(path) < k {
				path = append(path, i)
				needNum -= i
				dfs(path, i+1, needNum)
				needNum += i
				path = path[:len(path)-1]
			}
		}
	}
	dfs(path, 1, n)
	return result
}

// 最佳方式
// 最佳方式
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(index int, needNum int)
	dfs = func(index int, needNum int) {
		// 找到一个有效组合：路径长度等于 k 且剩余目标和为 0
		if len(path) == k && needNum == 0 {
			tmp := make([]int, k)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		// 剪枝：如果路径长度已经达到 k 或剩余目标和小于 0，直接返回
		if len(path) >= k || needNum <= 0 {
			return
		}

		// 从 index 开始遍历候选数字（1 到 9）
		for i := index; i <= 9; i++ {
			// 剪枝：如果当前数字大于剩余目标和，直接跳出循环
			if i > needNum {
				break
			}

			// 做选择：将当前数字加入路径
			path = append(path, i)

			// 递归：从下一个数字开始，继续寻找组合
			dfs(i+1, needNum-i)

			// 撤销选择：回溯
			path = path[:len(path)-1]
		}
	}

	dfs(1, n)
	return result
}

func main() {
	n := 7
	k := 3
	result := combinationSum3(k, n)
	fmt.Println("")
	fmt.Printf("result: %v", result)
}
