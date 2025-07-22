package main

import "fmt"

/*
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为 target 的不同组合数少于 150 个。



示例 1：

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。
示例 2：

输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]
示例 3：

输入: candidates = [2], target = 1
输出: []


提示：

1 <= candidates.length <= 30
2 <= candidates[i] <= 40
candidates 的所有元素 互不相同
1 <= target <= 40

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
存在问题：
产生重复组合：这是最核心的问题。因为在递归的每一层，你的 for 循环都从索引 0 开始，
允许选择之前已经考虑过的、索引更小的数字。这会产生排列组合，而不是题目要求的组合。例如，对于 [2, 3]，它会生成 [2, 3] 和 [3, 2] 两种路径。

如何修正：避免重复
要解决重复组合的问题，我们需要施加一个约束：在递归搜索时，只允许选择当前或之后的元素，不允许回头选择之前的元素。

如何实现这个约束？很简单，我们给 dfs 函数增加一个参数 startIndex。
dfs(target, startIndex): 表示在这一层递归中，你的 for 循环只能从 candidates 数组的 startIndex 位置开始选择。
当你选择了一个数 candidates[i] 并向下递归时，你传递的下一个 startIndex 仍然是 i。为什么还是 i？因为题目允许同一个数字被重复使用。
通过这种方式，如果我们选择了 candidates[i]，下一层就不会再回头选择 candidates[i-1]、candidates[i-2] 等等，从而避免了重复组合的产生。

可选的优化： 先对 candidates 数组进行排序。这样做的好处是，可以在 for 循环中进行更有效的“剪枝”。如果 sum + candidates[i] 已经大于 target，
那么后面的 candidates[i+1], candidates[i+2] 等等（因为它们更大）也肯定会超，可以直接 break 循环。
*/
func combinationSum(candidates []int, target int) [][]int {
	// 边界条件处理正确
	if len(candidates) == 0 {
		return [][]int{}
	}

	result := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(sum int)

	dfs = func(sum int) {
		// 递归终止条件1：找到一个有效的组合
		if sum == target {
			// 创建副本并加入结果集，这部分是完全正确的，做得很好！
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		// 递归终止条件2：当前路径和已超过目标，剪枝
		if sum > target {
			return
		}

		// --- 核心问题在这里 ---
		// for i := 0; i < len(candidates); i++ {
		// 这个循环的逻辑是：在递归的【每一层】，都【从头开始】遍历整个 `candidates` 数组。
		// 这就导致了重复组合的产生。
		//
		// 让我们用一个例子来说明：candidates = [2, 3, 6, 7], target = 7
		//
		// 路径 1:
		// - 第一次调用 dfs(0)
		// - for 循环从 i=0 (值为 2) 开始
		// - path=[2], sum=2, 调用 dfs(2)
		//   - 第二次调用 dfs(2)
		//   - for 循环又【从头 i=0】开始！
		//   - 假设它走到了 i=1 (值为 3)
		//   - path=[2, 3], sum=5, 调用 dfs(5)
		//     - ... (后续会失败)
		//
		// 路径 2:
		// - 第一次调用 dfs(0)
		// - for 循环走到了 i=1 (值为 3)
		// - path=[3], sum=3, 调用 dfs(3)
		//   - 第二次调用 dfs(3)
		//   - for 循环又【从头 i=0】开始！
		//   - 走到了 i=0 (值为 2)
		//   - path=[3, 2], sum=5, 调用 dfs(5)
		//
		// 你看，我们通过不同的选择顺序，得到了 `[2, 3, ...]` 和 `[3, 2, ...]` 两种路径。
		// 尽管它们的元素集合是相同的，但因为你的算法关心顺序，所以它们会被当作不同的路径处理。
		// 而题目要求的是【组合】，组合是不关心顺序的。`[2,3]` 和 `[3,2]` 是同一个组合。
		//
		// --- 另一个问题：sum 参数的传递和修改方式 ---
		// 你将 sum 作为参数传递，同时在循环内部又修改它。
		// `sum += candidates[i]`
		// `dfs(sum)`
		// `sum -= candidates[i]`
		// 这种写法虽然在这里能正常工作，但通常不是回溯中处理“状态”的最清晰方式。
		// 更常见的做法是直接在递归调用时传递新的状态：`dfs(sum + candidates[i])`，
		// 这样就不需要在递归返回后手动减去 `sum`。这使得代码更简洁，不易出错。
		for i := 0; i < len(candidates); i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			dfs(sum)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return result
}

/*func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	n := len(candidates)
	dfs := func(int) {}
	var result []int
	temp := target
	dfs = func(temp int) {
		if temp < 0 {
			return
		}
		if temp ==0 {
			tmp := make([]int, len(result))
			copy(tmp ,result)
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < n ; i++ {
			if candidates[i] == 0 {
				continue
			}
			temp = temp - candidates[i]

			result = append(result, candidates[i])
			dfs(temp)
			temp = temp + candidates[i]
			result = result[:n-1]
		}
	}
	dfs(temp)

	return ans
}*/

/*
上述思路此外：原因如下：
1.要避免重复寻找，for循环i应该通过变量传递，也就是startIndex,for i := 也就是startIndex; i < n ; i++ {错误；
2.result = result[:n-1],应该为result = result[:len(result)-1] ，不然数组长度小于n时并不会成功回溯；

树形结构应该如下：
https://img-blog.csdnimg.cn/20201123202227835.png
相关讲解视频：
https://www.bilibili.com/video/BV1KT4y1M7HJ/?vd_source=34718180774b041b23050c8689cdbaf2
*/

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	n := len(candidates)

	// 定义深度优先搜索函数，用于生成组合
	var dfs func(temp int, startIndex int, result []int)

	dfs = func(temp int, startIndex int, result []int) {
		if temp < 0 {
			return
		}
		if temp == 0 {
			// 如果目标值为0，表示找到了一个组合，将结果复制到临时切片并添加到答案中
			tmp := make([]int, len(result))
			copy(tmp, result)
			ans = append(ans, tmp)
			return
		}
		for i := startIndex; i < n; i++ {
			if candidates[i] == 0 {
				continue // 跳过候选列表中的0
			}
			temp = temp - candidates[i]            // 尝试将当前候选值添加到组合中
			result = append(result, candidates[i]) // 记录当前候选值
			// 递归搜索下一个组合
			dfs(temp, i, result)
			temp = temp + candidates[i]     // 回溯，将当前候选值从组合中移除
			result = result[:len(result)-1] // 回溯，从结果中移除最后一个元素
		}
	}

	// 初始化深度优先搜索
	dfs(target, 0, []int{})

	return ans
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Printf("result:%v", combinationSum(candidates, target))
}
