package main

/*
112. 路径总和

给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
叶子节点 是指没有子节点的节点。

示例 1：
输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true
解释：等于目标和的根节点到叶节点路径如上图所示。
示例 2：
输入：root = [1,2,3], targetSum = 5
输出：false
解释：树中存在两条根节点到叶子节点的路径：
(1 --> 2): 和为 3
(1 --> 3): 和为 4
不存在 sum = 5 的根节点到叶子节点的路径。
示例 3：
输入：root = [], targetSum = 0
输出：false
解释：由于树是空的，所以不存在根节点到叶子节点的路径。
*/
/*
解题思路：
sum ： 从根节点到叶子节点的路径上的节点值相加的目标和
对 root 递归。转为判断：root 的左、右子树中能否找出和为 sum-root.val 的路径
就变成一个规模小一点的相同问题
即，每遍历一个节点，sum 就减去当前节点值，当遍历到叶子节点时，因为没有子节点了，如果 sum - 当前叶子节点值 == 0 ，即找到了从根节点到叶子节点的和为 sum 的路径
时间复杂度：O(n)，每个节点被遍历一次

*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 基于递归思路
/*func hasPathSum(root *TreeNode, targetSum int) bool {
	// 如果根节点为空，返回 false
	if root == nil {
		return false
	}

	// 如果当前节点是叶子节点，且路径和等于 targetSum，则返回 true
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}

	// 递归检查左子树和右子树
	leftPath := hasPathSum(root.Left, targetSum-root.Val)
	rightPath := hasPathSum(root.Right, targetSum-root.Val)

	// 返回左子树或右子树的结果
	return leftPath || rightPath
}*/

// 基于广度遍历
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 如果根节点为空，返回 false
	if root == nil {
		return false
	}

	// 初始化两个队列，一个用于存储节点，一个用于存储从根节点到当前节点的路径和
	queue := []*TreeNode{root}
	valQueue := []int{root.Val}

	// 使用广度优先搜索（BFS）遍历二叉树
	for len(queue) != 0 {
		n := len(queue) // 当前层的节点数量

		// 遍历当前层的节点
		for i := 0; i < n; i++ {
			node := queue[i]   // 获取当前节点
			val := valQueue[i] // 获取从根节点到当前节点的路径和

			// 如果当前节点是叶子节点且路径和等于目标值，返回 true
			if node.Left == nil && node.Right == nil && val == targetSum {
				return true
			}

			// 将左子节点和对应的路径和加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
				valQueue = append(valQueue, node.Left.Val+val)
			}

			// 将右子节点和对应的路径和加入队列
			if node.Right != nil {
				queue = append(queue, node.Right)
				valQueue = append(valQueue, node.Right.Val+val)
			}
		}

		// 移除已经处理的节点和路径和
		valQueue = valQueue[n:]
		queue = queue[n:]
	}

	// 遍历完整棵树后仍未找到路径和等于目标值的情况，返回 false
	return false
}



func main(){
	root :=	&TreeNode{Val: 5,Left: &TreeNode{Val: 4,Left: &TreeNode{Val: 11,Left: &TreeNode{Val: 7},Right: &TreeNode{Val: 2}}},Right: &TreeNode{Val: 8,Left: &TreeNode{Val: 13},Right: &TreeNode{Val: 4,Right: &TreeNode{Val: 1}}}}
	print(hasPathSum(root, 22))
}



