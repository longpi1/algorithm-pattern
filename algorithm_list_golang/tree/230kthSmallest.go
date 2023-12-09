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
	var result []int
	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {return}
		dfs(n.Left)
		result = append(result, n.Val)
		dfs(n.Right)
	}
	dfs(root)
	return result[k-1]
}


//优化方法 用递归实现中序再k--
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