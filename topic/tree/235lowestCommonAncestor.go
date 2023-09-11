package main

/*
235. 二叉搜索树的最近公共祖先

给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]

示例 1:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。
示例 2:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。

*/
/*
解题思路：
二叉搜索树的最近公共祖先不需要判断当前节点是否为空:
因为是有序的,因此可以知道p/q在左子树里或者右子树里,找到了就返回,因此没有遍历到空节点上.
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	x := root.Val
	if p.Val < x && q.Val < x {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > x && q.Val > x {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
