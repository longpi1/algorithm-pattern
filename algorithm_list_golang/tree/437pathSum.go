package main

/*
437. 路径总和 III

给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

示例 1：
输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
输出：3
解释：和等于 8 的路径有 3 条，如图所示。
示例 2：
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：3

*/

// 这个题不会，解题思路如下
/*

*/
// 计算以当前节点为根节点的子树中，路径和等于目标和的路径数量
func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return 0
	}
	val := root.Val

	// 如果当前节点的值等于目标和，增加路径计数器
	if val == targetSum {
		res++
	}

	// 递归计算左子树中满足条件的路径数量
	res += rootSum(root.Left, targetSum-val)

	// 递归计算右子树中满足条件的路径数量
	res += rootSum(root.Right, targetSum-val)

	return res
}

// 计算二叉树中，路径和等于目标和的路径数量
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 以当前节点为根节点的子树中满足条件的路径数量
	res := rootSum(root, targetSum)

	// 递归计算左子树中满足条件的路径数量
	res += pathSum(root.Left, targetSum)

	// 递归计算右子树中满足条件的路径数量
	res += pathSum(root.Right, targetSum)

	return res
}

