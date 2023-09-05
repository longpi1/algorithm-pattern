package main

/*
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
示例 1：
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]

示例 2：
输入：root = [2,1,3]
输出：[2,3,1]
示例 3：

输入：root = []
输出：[]


提示：

树中节点数目范围在 [0, 100] 内
-100 <= Node.val <= 100
*/



type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//前序遍历实现
func invertTree(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	root.Left , root.Right=  swap(root.Left,root.Right)
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}


func swap(left *TreeNode, right *TreeNode) (*TreeNode,*TreeNode){
	tmp := left
	left = right
	right = tmp
	return left, right
}

//层序遍历实现：
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := []*TreeNode{root}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		cur.Left, cur.Right = cur.Right, cur.Left

		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
	return root
}

func main(){
	root :=	TreeNode{Val: 1,Left: &TreeNode{Val: 2},Right: &TreeNode{Val: 3,Left: &TreeNode{Val: 4}}}
	invertTree(&root)
}