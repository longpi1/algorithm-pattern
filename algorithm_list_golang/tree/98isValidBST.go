package main

import "math"

/*
98. 验证二叉搜索树
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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路错误
/*
	错误示例
     5
    / \
   4   6
      / \
     3   7 <-- 这里 3 < 5，不满足 BST 定义
*/
func isValidBST(root *TreeNode) bool {
	// --- 问题 1: 对空树的定义错误 ---
	// if root == nil {
	// 	return false
	// }
	// 根据二叉搜索树的定义，一棵空树 (nil) 也是一棵有效的二叉搜索树。
	// 所以这里应该返回 true。
	if root == nil {
		return false // 错误，应为 return true
	}

	// 使用一个闭包内的 flag 变量来记录状态，这是一种可行的模式。
	flag := true

	// 声明和定义 dfs 函数
	var dfs func(nodeVal int, node *TreeNode)
	dfs = func(nodeVal int, node *TreeNode) {
		// --- 问题 2: 核心逻辑错误 - 只检查了局部性质 ---
		// BST 的定义是：
		// 1. 对于任意节点 N，其左子树中【所有】节点的值都必须小于 N.Val。
		// 2. 其右子树中【所有】节点的值都必须大于 N.Val。
		// 3. 其左右子树也必须分别是二叉搜索树。
		//
		// 你的代码只检查了 node.Left.Val < node.Val 和 node.Right.Val > node.Val。
		// 这只是一个必要条件，但不是充分条件。
		// 它没有检查到左子树中【所有】节点的情况。
		if node.Left != nil {
			// 你只比较了直接子节点和父节点的值。
			if node.Left.Val < nodeVal { // nodeVal 在这里就是 node.Val
				// --- 问题 3: 递归参数传递错误 ---
				// dfs(node.Left.Val, node.Left)
				// 当你向下递归到左子树时，你传递了 `node.Left.Val` 作为新的“父节点值”。
				// 这就丢失了来自更上层祖先节点的约束！
				// 例如，对于 [5, 1, 7, nil, nil, 6, 8] 这棵树：
				//       5
				//      / \
				//     1   7
				//        / \
				//       6   8
				//
				// 当你从 5 递归到 7 时，是合法的 (7 > 5)。
				// 然后你从 7 递归到 6，调用 `dfs(7, node for 6)`。
				// 在这个调用里，你检查 `6 < 7`，这是对的。
				// 但你忽略了一个至关重要的约束：节点 6 位于节点 5 的右子树中，
				// 所以它【必须】大于 5！你的代码没有这个约束检查。
				dfs(node.Left.Val, node.Left)
			} else {
				flag = false
				return
			}
		}

		// 对右子树的检查也存在同样的问题。
		if node.Right != nil {
			if node.Right.Val > nodeVal {
				// 同样，这里也丢失了来自祖先节点的约束。
				// 节点 6 虽然大于它的直接父节点 7 是错的，但假设是另一个例子：
				// [10, 5, 15, nil, nil, 6, 20]
				// 节点 6 虽然大于 5，但它小于其祖先 10，这是不合法的。
				// 你的代码无法发现这种情况。
				dfs(node.Right.Val, node.Right)
			} else {
				flag = false
				return
			}
		}
	}

	dfs(root.Val, root)
	return flag
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
1. 递归法 (带上下界) - 最直观
在递归时，不仅传递当前节点，还要传递一个 (min, max) 范围，表示当前节点的值必须在这个开区间内。
*/
func isValidBST(root *TreeNode) bool {
	// 使用 math.MinInt64 和 math.MaxInt64 作为初始的无穷小和无穷大边界
	return check(root, math.MinInt64, math.MaxInt64)
}

// check 函数验证以 node 为根的子树是否满足 BST 定义，
// 且所有节点的值必须在 (lower, upper) 这个开区间内。
func check(node *TreeNode, lower, upper int) bool {
	// 空树是有效的 BST
	if node == nil {
		return true
	}

	// 检查当前节点的值是否在有效范围内
	if node.Val <= lower || node.Val >= upper {
		return false
	}

	// 递归检查左右子树：
	// - 对左子树，它的所有节点必须小于当前节点的值，所以更新上界为 node.Val
	// - 对右子树，它的所有节点必须大于当前节点的值，所以更新下界为 node.Val
	return check(node.Left, lower, node.Val) && check(node.Right, node.Val, upper)
}

/*
2. 中序遍历法 - 最巧妙
利用 BST 的一个重要特性：一个有效的 BST，其中序遍历的结果必然是一个严格递增的序列。
*/
func isValidBST(root *TreeNode) bool {
	// prev 用于记录中序遍历中前一个节点的值
	// 使用一个指针，以便在递归中传递引用
	prev := math.MinInt64
	isBST := true

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || !isBST { // 如果已经发现不是 BST，可以提前终止
			return
		}

		// 1. 遍历左子树
		inorder(node.Left)

		// 2. 访问当前节点
		// 如果当前节点的值不大于前一个节点的值，则不是递增序列，不是 BST
		if node.Val <= prev {
			isBST = false
			return
		}
		// 更新 prev 为当前节点的值，为下一次比较做准备
		prev = node.Val

		// 3. 遍历右子树
		inorder(node.Right)
	}

	inorder(root)
	return isBST
}

func main() {
	//root :=	&TreeNode{Val: 5,Left: &TreeNode{Val: 4},Right: &TreeNode{Val: 6,Left: &TreeNode{Val: 3},Right: &TreeNode{Val: 7}}}
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}
	print(isValidBST(root))
}
