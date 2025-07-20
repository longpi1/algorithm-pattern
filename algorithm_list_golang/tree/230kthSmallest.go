package main

/*
230. 二叉搜索树中第K小的元素
相关企业
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

示例 1：
输入：root = [3,1,4,null,2], k = 1
输出：1

示例 2：
输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3
*/
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	// --- 问题 1: visit map 的使用方式和时机错误 ---
	// 在标准的中序遍历迭代法中，我们不需要一个 `visit` map。
	// 栈的 push/pop 机制本身就能保证节点的访问顺序。
	// 你引入了 `visit` map，但它的使用逻辑也是错误的。
	visit := make(map[*TreeNode]bool)

	stack := []*TreeNode{root}
	result := make([]int, 0)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		for node.Left != nil && !visit[node.Left] {
			visit[node.Left] = true
			stack = append(stack, node.Left)
			// 切记记得切换至左子树以及标记已访问
			node = node.Left
			// --- 第一次写的错误: 灾难性的 visit 标记 ---
			// visit[node.Left] = true
			// 假设当前 node 是 5，它的左子节点是 3。
			// 你标记的是 `visit[3.Left]` 为 true，而不是 `visit[3]`。
			// 如果 3.Left 是 nil，你就在 map 中创建了一个 `visit[nil] = true` 的条目，这没有意义。
			// 如果 3.Left 不是 nil (比如是2)，你标记的是 `visit[2] = true`。
			// 这意味着你永远是标记“当前节点的左孩子的左孩子”，逻辑完全错了。
			// 正确的标记应该是针对已处理过的节点。
			// visit[node.Left] = true
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return result[k-1]
}

func kthSmallest(root *TreeNode, k int) int {
	var result []int
	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		result = append(result, n.Val)
		dfs(n.Right)
	}
	dfs(root)
	return result[k-1]
}

// 优化方法 用递归实现中序再k--
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		stack, root = stack[:len(stack)-1], stack[len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
