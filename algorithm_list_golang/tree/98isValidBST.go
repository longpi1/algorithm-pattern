package main

import "math"

/*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。


示例 1：
输入：root = [2,1,3]
输出：true

示例 2：
输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。
*/

/*func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}


	return isVaild(root)
}

func isVaild(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Val <= root.Left.Val {
			return false
		}
	}
	if root.Right != nil {
		if root.Val >= root.Right.Val {
			return false
		}
	}
	return isVaild(root.Left) && isVaild(root.Right)
}*/

/*
上述思路错误：
不能保证是否满足之前的元素
*/


/*
解题思路：
基于中序遍历解题，因为二叉搜索树中序遍历结果为从小到大
*/
//基于中序遍历解题，因为二叉搜索树中序遍历结果为从小到大
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//这里不应该初始化
	//pre := &TreeNode{}
	var pre *TreeNode
	// 调用辅助函数 isVaild 来验证二叉搜索树的有效性
	return isVaild(root,&pre)
}

func isVaild(root *TreeNode,pre **TreeNode) bool {
	if root == nil {
		return true
	}
	// 递归验证左子树
	flag1 := isVaild(root.Left,pre)
	if *pre != nil && root.Val <= (*pre).Val {
		return false
	}
	//在 Go 中，函数参数是按值传递的，这意味着在 isVaild 函数中的 pre 变量实际上是 root 的一个副本，而不是指向同一个内存位置的指针。因此，你的 pre 副本不会在函数内部正确地跟踪前一个节点。
	//pre = root

	// 将 pre 参数改为指向指针的指针 **TreeNode，这样可以在递归调用中正确地更新前一个节点的值。这样的修改可以确保你正确地检查二叉搜索树的有效性。
	*pre = root
	// 递归验证右子树
	flag2 := isVaild(root.Right, pre)

	// 返回左子树和右子树的验证结果的逻辑与
	return flag1 && flag2
}


// 其他思路
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}


func main()  {
	//root :=	&TreeNode{Val: 5,Left: &TreeNode{Val: 4},Right: &TreeNode{Val: 6,Left: &TreeNode{Val: 3},Right: &TreeNode{Val: 7}}}
	root :=	&TreeNode{Val: 2,Left: &TreeNode{Val: 2}}
	print(isValidBST(root))
}
