package main

/*
给你一棵二叉树的根节点，返回该树的 直径 。

二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。

两节点之间路径的 长度 由它们之间边数表示。

示例 1：
输入：root = [1,2,3,4,5]
输出：3
解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
示例 2：

输入：root = [1,2]
输出：1
*/


/*func diameterOfBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}
	// 分别求左右节点最大深度然后再+1
	leftDepth := nodeDepth(root.Left)
	rightDepth := nodeDepth(root.Right)

	return leftDepth + rightDepth +1
}

// 基于队列实现此题：
func nodeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0

	for len(queue) != 0 {
		levelSize := len(queue)
		for i := 0; i< levelSize; i++ {

			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)

			}
		}
		depth ++
		queue = queue[levelSize:]
	}
	return depth
}
*/

/*
上述思路存在一些问题：
代码看起来基本上是正确的，但在计算二叉树的直径时，存在一个问题。直径不仅仅是左子树深度和右子树深度的和加1，因为直径是通过根节点的最长路径，而这条路径可能经过根节点。
*/


func diameterOfBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}

	// 分别求左右节点最大深度然后再+1
	leftDepth := nodeDepth(root.Left)
	rightDepth := nodeDepth(root.Right)
	// 分别计算左子树和右子树的直径
	leftDiameter := diameterOfBinaryTree(root.Left)
	rightDiameter := diameterOfBinaryTree(root.Right)
	// 取三者中的最大值作为结果
	return max(leftDepth+rightDepth, max(leftDiameter, rightDiameter))
}

// 基于队列实现此题：
func nodeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0

	for len(queue) != 0 {
		levelSize := len(queue)
		for i := 0; i< levelSize; i++ {

			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)

			}
		}
		depth ++
		queue = queue[levelSize:]
	}
	return depth
}

// 辅助函数，计算两个整数的最大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 其他思路：
/*
dfs深度优先遍历递归解决 最大直径maxDia初始值为0 计算以指定节点为根的子树(二叉树)的最大直径maxDia,其实就是该二叉树左子树的最大高度+右子树的最大高度,
因此我们对二叉树从根节点root开始进行深度优先遍历，每遍历一个节点，迭代maxDia的值(将maxDia与以该节点为根的二叉树的最大直径即lh+rh之和比较，取较大值)，
遍历结束后的maxDia即为该二叉树的最大直径。dfs深度优先遍历时返回以该节点为根的二叉树的最大高度，也就是1+max(lh+rh),时间复杂度度O(n),空间复杂度O(h),n为该二叉树节点个数，h为该二叉树高度。

*/

func diameterOfBinaryTree(root *TreeNode) int {
	maxDia := 0
	if root == nil{
		return maxDia
	}
	var dfs func(*TreeNode)int
	dfs = func(node *TreeNode)int{
		if node == nil{
			return 0
		}
		lh := dfs(node.Left)
		rh := dfs(node.Right)
		maxDia = max(maxDia, lh+rh)
		return 1 + max(lh, rh)
	}
	dfs(root)
	return maxDia
}
