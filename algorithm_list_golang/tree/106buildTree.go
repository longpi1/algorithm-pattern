package main

/*
106. 从中序与后序遍历序列构造二叉树

给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。

示例 1:
输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]

示例 2:
输入：inorder = [-1], postorder = [-1]
输出：[-1]

*/

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	//根节点
	node := &TreeNode{Val: postorder[n-1]}
	index := 0
	for i := index; i < n; index++ {
		if inorder[index] == postorder[n-1] {
			break
		}
	}
	// 分割后序
	leftPostorder := postorder[:index]
	rightPostorder := postorder[index:n]
	// 分割中序
	leftInorder := inorder[:index]
	rightInorder := inorder[index+1:]

	node.Left = buildTree(leftInorder,leftPostorder)
	node.Right = buildTree(rightInorder,rightPostorder)
	return node
}

func main(){
	postorder := []int{9,15,7,20,3}
	inorder := []int{9,3,15,20,7}
	buildTree(inorder,postorder)
}