package main

import "fmt"

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。

示例 1：

输入：n = 4, k = 2
输出：
[

	[2,4],
	[3,4],
	[2,3],
	[1,2],
	[1,3],
	[1,4],

]
示例 2：

输入：n = 1, k = 1
输出：[[1]]

提示：

1 <= n <= 20
1 <= k <= n
*/
func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	dfsCombine := func(num int, combineArr []int) {
		if len(combineArr) == k {
			result = append(result, combineArr)
		}
		for i := num; i < n; i++ {
			// 需提前定义对应函数
			dfsCombine(i+1, combineArr)
			path = path[len(path)-1]
		}
	}
	dfsCombine(1, path)
	return result
}

func combine(n int, k int) [][]int {
	result := make([][]int, 0)              // 存储所有符合条件的组合
	currentCombination := make([]int, 0, k) // 存储当前正在构建的组合，预分配容量

	// 定义一个辅助的递归（回溯）函数
	// `start`: 表示当前可以从哪个数字开始选择（避免重复组合和生成重复的元素）
	// `currentCombination`: 已经选择的数字列表
	// `result`: 存储最终结果的引用
	var backtrack func(start int, currentCombination []int)
	backtrack = func(start int, currentCombination []int) {
		// 递归终止条件：当当前组合的长度达到 k 时，说明找到一个有效组合
		if len(currentCombination) == k {
			// **注意：这里必须创建 currentCombination 的副本**
			// 否则，当 currentCombination 在回溯过程中被修改时，result 中存储的引用会指向错误的数据
			temp := make([]int, k)
			copy(temp, currentCombination)
			result = append(result, temp)
			return // 结束当前路径的递归
		}

		// 剪枝优化：如果剩余可选的数字不足以填满 k 个位置，则无需继续
		// n - start + 1 是从 start 到 n 剩余的数字数量
		// k - len(currentCombination) 是还需要选择的数字数量
		if n-start+1 < k-len(currentCombination) {
			return
		}

		// 遍历所有可能的选择
		// 从 start 开始，到 n 结束
		for i := start; i <= n; i++ {
			// 1. 选择当前数字 i，加入组合
			currentCombination = append(currentCombination, i)

			// 2. 递归：从下一个数字 (i + 1) 开始，继续选择
			backtrack(i+1, currentCombination)

			// 3. 回溯：撤销选择，将当前数字 i 从组合中移除，尝试下一个可能的选择
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}

	// 从数字 1 开始，进行回溯
	backtrack(1, currentCombination)

	return result
}

func combine(n int, k int) [][]int {
	var result [][]int
	var path []int

	dfs := func(int) {}
	dfs = func(index int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := index; i < n+1; i++ {

			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return result

}

func main() {
	n := 4
	k := 2
	fmt.Printf("result: %v", combine(n, k))
}
