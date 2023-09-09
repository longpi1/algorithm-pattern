package main


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
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
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

*/
func isValidBST(root *TreeNode) bool {
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
}

func main()  {
	root :=	&TreeNode{Val: 5,Left: &TreeNode{Val: 4},Right: &TreeNode{Val: 6,Left: &TreeNode{Val: 3},Right: &TreeNode{Val: 7}}}
	print(isValidBST(root))
}
