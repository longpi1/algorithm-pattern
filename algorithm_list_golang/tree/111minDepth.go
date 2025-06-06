package main

import "math"

/*
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：2


示例 2：
输入：root = [2,null,3,null,4,null,5,null,6]
输出：5


提示：
树中节点数的范围在 [0, 105] 内
-1000 <= Node.val <= 1000
*/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	min := 1
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				return min
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[levelSize:]
		min++
	}
	return min

}

// 基于队列层序遍历，也就是广度优先遍历
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0

	for len(queue) != 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				return depth + 1
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
		queue = queue[levelSize:]
	}
	return depth
}

// 深度优先遍历没有想出来，这里的参照官方，基于类似后序的思路进行递归遍历
func minDepth(root *TreeNode) int {
	// 如果根节点为空，返回深度0
	if root == nil {
		return 0
	}

	// 如果根节点没有左子树和右子树，说明根节点为叶子节点，返回深度1
	if root.Left == nil && root.Right == nil {
		return 1
	}

	// 初始化最小深度为一个较大的整数
	minD := math.MaxInt32

	// 递归计算左子树的最小深度，并与当前的最小深度比较取较小值
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}

	// 递归计算右子树的最小深度，并与当前的最小深度比较取较小值
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}

	// 返回最小深度加1，表示当前节点的深度
	return minD + 1
}

// 定义辅助函数 min，用于比较两个整数的较小值
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// minDepth 优化后的递归版本
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0 // 空树的最小深度为0
	}

	// 如果是叶子节点，其深度为1
	if root.Left == nil && root.Right == nil {
		return 1
	}

	// 如果只有左子节点，则最小深度是左子树的最小深度 + 1
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	// 如果只有右子节点，则最小深度是右子树的最小深度 + 1
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	// 如果左右子节点都存在，则最小深度是左右子树最小深度中的较小值 + 1
	leftMinDepth := minDepth(root.Left)
	rightMinDepth := minDepth(root.Right)

	return int(math.Min(float64(leftMinDepth), float64(rightMinDepth))) + 1
}
func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}
	print(minDepth(root))
}
