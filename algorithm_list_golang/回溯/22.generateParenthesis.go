package main

/*

22. 括号生成
相关企业
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：
输入：n = 1
输出：["()"]


提示：
1 <= n <= 8
*/

// 第一次写思路错误
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	result := make([]string, 0)

	// --- 问题 1: 变量作用域和状态管理混乱 ---
	// 你在外部定义了 leftCount 和 rightCount，但又将它们作为参数传递给 dfs。
	// 这使得状态的管理变得非常混乱。在 dfs 内部，你修改的是【参数的副本】，
	// 而不是外部的变量。虽然你尝试了 `leftCount++` 和 `leftCount--`，
	// 但这在一个错误的框架下是无效的。
	leftCount, rightCount := 0, 0
	path := make([]byte, 0)

	var dfs func(leftCount, rightCount int)
	dfs = func(leftCount, rightCount int) {
		// --- 问题 2: 剪枝条件不正确 ---
		// if leftCount > n || rightCount > n {
		// 这个条件是正确的，一个有效的括号组合中，左括号和右括号的数量都不能超过 n。
		// 但它应该放在递归的入口处，而不是在找到解之后。
		if leftCount > n || rightCount > n {
			return
		}

		// 终止条件是正确的：当路径长度达到 2*n 时，我们找到了一个完整的组合。
		if leftCount+rightCount == 2*n {
			// 这里直接 append(result, string(path)) 是有风险的，
			// 因为 path 是一个共享的切片。虽然在你这个版本的逻辑下可能不会出错，
			// 但标准做法是创建一个副本。不过这不是最主要的问题。
			result = append(result, string(path))
			return
		}

		// --- 问题 3: 灾难性的 for 循环 ---
		// for i := 0; i < 2*n; i++ {
		// 这是整个代码中最根本的错误。回溯算法的核心是做“选择”，
		// 在这个问题中，每一步的选择只有两个：“添加一个左括号”或“添加一个右括号”。
		// 你在这里引入了一个 `2*n` 次的 for 循环，这是完全没有必要的，
		// 它破坏了回溯的决策树模型。
		// 你在循环的每次迭代中都尝试进行决策，这会导致生成大量重复且无效的路径。
		for i := 0; i < 2*n; i++ {

			// --- 问题 4: 决策逻辑（剪枝条件）错误 ---
			// if leftCount <= rightCount {
			// 这个条件是错误的。正确的约束是：
			// 1. 只要【左括号的数量 < n】，我们就可以添加一个左括号。
			// 2. 只要【右括号的数量 < 左括号的数量】，我们就可以添加一个右括号（确保括号是合法的）。
			//
			// 你的 `leftCount <= rightCount` 逻辑是说：
			// “如果左括号不多于右括号，就加一个左括号”。这会导致一开始 `0 <= 0`，
			// 加一个左括号，`path` 变成 `(`，`leftCount` 变成 1。
			// 然后在下一次 for 循环中，`1 <= 0` 不成立，进入 else...
			// 这整个逻辑是混乱的。
			if leftCount <= rightCount {
				path = append(path, '(')
				// leftCount++ 是在修改参数的副本，这个修改不会传递回上层。
				leftCount++
				dfs(leftCount, rightCount)
				// leftCount-- 同样是修改副本，然后这个修改就丢失了。
				leftCount--
				path = path[:len(path)-1]
			} else {
				path = append(path, ')')
				rightCount++
				dfs(leftCount, rightCount)
				rightCount--
				path = path[:len(path)-1]
			}
		}
	}

	// 初始调用，leftCount 和 rightCount 都是 0。
	dfs(leftCount, rightCount)
	return result
}

/*
正确的 Go 代码实现 (回溯法)
让我们抛弃错误的框架，从头开始构建一个清晰、正确的回溯解法。

核心思想（剪枝条件）：

在递归的每一步，我们都有两种选择：放一个 ( 或者放一个 )。但这些选择不是无条件的，必须满足：

可以放 ( 的条件：只要已经放置的左括号数量 left 还小于 n。
可以放 ) 的条件：只要已经放置的右括号数量 right 还小于已经放置的左括号数量 left。（这保证了括号序列的合法性，不会出现 )( 这样的情况）。
递归终止条件： 当生成的字符串长度达到 2*n 时，说明我们找到了一个完整的、合法的组合。
新旧代码对比
回溯结构：修正后的版本没有 for 循环，而是通过两个 if 语句来代表两个并列的“选择”，这才是正确的决策树模型。
剪枝条件：使用了 left < n 和 right < left 这两个正确的、核心的剪枝条件。
状态管理：状态（left, right）完全通过函数参数传递，清晰明了，没有外部变量和参数副本的混淆。
终止条件：left == n && right == n 作为终止条件更精确。
*/
func generateParenthesis(n int) []string {
	// 如果 n 为 0，可以直接返回一个空字符串的组合，但 LeetCode 通常要求返回空列表。
	if n == 0 {
		return []string{}
	}

	result := make([]string, 0)
	// path 用于存储当前正在构建的括号组合
	path := make([]byte, 0, 2*n) // 预分配容量以提高效率

	// dfs 辅助函数
	// left: 已使用的左括号数量
	// right: 已使用的右括号数量
	var dfs func(left, right int)

	dfs = func(left, right int) {
		// --- 终止条件 ---
		// 当已使用的左括号和右括号数量都等于 n 时，
		// 说明我们构建了一个长度为 2*n 的完整组合。
		if left == n && right == n {
			// 将当前路径的副本加入结果集
			result = append(result, string(path))
			return
		}

		// --- 递归和剪枝 ---
		// 选择 1: 放置一个左括号
		// 条件：已使用的左括号数量 < n
		if left < n {
			// 做选择
			path = append(path, '(')
			// 递归到下一层，左括号数量+1
			dfs(left+1, right)
			// 撤销选择 (回溯)
			path = path[:len(path)-1]
		}

		// 选择 2: 放置一个右括号
		// 条件：已使用的右括号数量 < 已使用的左括号数量
		// 这个条件保证了括号序列的合法性。
		if right < left {
			// 做选择
			path = append(path, ')')
			// 递归到下一层，右括号数量+1
			dfs(left, right+1)
			// 撤销选择 (回溯)
			path = path[:len(path)-1]
		}
	}

	// 从 0 个左括号和 0 个右括号开始
	dfs(0, 0)
	return result
}
