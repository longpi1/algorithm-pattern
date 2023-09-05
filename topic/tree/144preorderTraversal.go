package main

/*
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。



示例 1：


输入：root = [1,null,2,3]
输出：[1,2,3]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
示例 4：


输入：root = [1,2]
输出：[1,2]
示例 5：


输入：root = [1,null,2]
输出：[1,2]

*/

//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int,0)

	return traversal(root,result)
}

// 递归
func traversal(cur *TreeNode, result []int) []int{
	if cur == nil {
		return result
	}
	result = append(result, cur.Val)
	result=traversal(cur.Left,result)
	result=traversal(cur.Right,result)
	return result
}



// 非递归遍历
func preorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	stack = append(stack, root)
	if root == nil{
		return []int{}
	}
	for len(stack) != 0 {
		// !!! 定义对应的接收栈
		node := stack[len(stack)-1]
		result = append(result, node.Val)
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}

func main(){
  root :=	TreeNode{Val: 1,Left: &TreeNode{Val: 2},Right: &TreeNode{Val: 3}}
  preorderTraversal(&root)
}