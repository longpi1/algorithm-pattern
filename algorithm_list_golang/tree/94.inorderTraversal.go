package main

/*
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。



示例 1：


输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]


提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	visit := make(map[*TreeNode]bool)
	result := make([]int, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		for node.Left != nil && !visit[node.Left] {
			visit[node.Left] = true
			stack = append(stack, node.Left)
			// 第一次写忘记切节点位置，导致一直循环超时
			node = node.Left
		}
		n := len(stack) - 1

		result = append(result, node.Val)
		stack = stack[:n]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return result
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	return traversal(root, result)
}

// 递归
func traversal(cur *TreeNode, result []int) []int {
	if cur == nil {
		return result
	}

	result = traversal(cur.Left, result)
	result = append(result, cur.Val)
	result = traversal(cur.Right, result)
	return result
}

/*func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	if root == nil{
		return []int{}
	}
	// !!! 首先将左子树遍历添加入栈
	stack = append(stack, root)
	for root.Left != nil {
		stack = append(stack, root.Left)
		root = root.Left
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
}*/

// 上述思路错误

// 实现2非递归遍历
/*
解题技巧：
1.创建一个空切片 res 用于存储中序遍历的结果。
2.创建一个空栈 stack 用于辅助遍历。
3.使用一个循环，将当前节点 root 的左子树上的所有节点都入栈，并不断更新 root 到左子节点，直到 root 为空或没有左子节点为止。这个过程将把左子树上的所有节点都依次入栈。
4.当左子树上的节点都入栈后，开始出栈，取出栈顶节点 node，将其值添加到结果切片 res 中。
5.然后，将 node 的右子节点（如果存在）赋值给 root，并进入一个新的循环。这个新循环将当前节点 root 的右子树上的所有节点都入栈，直到没有右子节点为止。
6此时，栈顶节点的左子树和右子树都已经被处理，继续出栈，添加到结果切片中。
7.重复上述过程，直到栈为空且所有节点都被处理完。
8.最终，返回结果切片 res，其中包含了二叉树的中序遍历结果。
*/

// 迭代法刚开始思路错误
func inorderTraversal(root *TreeNode) []int {
	// 1.创建一个空切片 res 用于存储中序遍历的结果。
	res := []int{}
	// 2.创建一个空栈 stack 用于辅助遍历。
	stack := []*TreeNode{}
	// 3.使用一个循环，将当前节点 root 的左子树上的所有节点都入栈，并不断更新 root 到左子节点，直到 root 为空或没有左子节点为止。这个过程将把左子树上的所有节点都依次入栈。
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for len(stack) != 0 {
		// 4.当左子树上的节点都入栈后，开始出栈，取出栈顶节点 node，将其值添加到结果切片 res 中。
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		// 5.然后，将 node 的右子节点（如果存在）赋值给 root，并进入一个新的循环。这个新循环将当前节点 root 的右子树上的所有节点都入栈，直到没有右子节点为止。
		node = node.Right

		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		// 6此时，栈顶节点的左子树和右子树都已经被处理，继续出栈，添加到结果切片中。
	}
	return res
}

// 迭代法其他方式

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for 0 < len(stack) || root != nil { //root != nil 只为了第一次root判断，必须放最后
		for root != nil {
			stack = append(stack, root) //入栈
			root = root.Left            //移至最左
		}
		n := len(stack) - 1             //栈顶
		res = append(res, stack[n].Val) //中序输出
		root = stack[n].Right           //右节点会进入下次循环，如果 =nil，继续出栈
		stack = stack[:n]               //出栈
	}
	return res
}
