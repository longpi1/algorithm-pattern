package main

/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。

示例 1：
输入：root = [1,2,2,3,4,4,3]
输出：true

示例 2：
输入：root = [1,2,2,null,3,null,3]
输出：false
提示：

树中节点数目在范围 [1, 1000] 内
-100 <= Node.val <= 100
*/


func isSymmetric(root *TreeNode) bool {
	queue := make([]*TreeNode,0)
	queue = append(queue, root)

	for len(queue) !=0 {
		root := queue[0]
		if root.Left

	}
	return true
}