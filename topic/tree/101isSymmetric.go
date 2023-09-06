package main

/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。

示例 1：
输入：root = [1,2,2,3,4,4,3]
输出：true

示例 2：
输入：root = [1,2,2,null,3,null,3]
输出：false
提示：

树中节点数目在范围 [1, 1000] 内
-100 <= Node.val <= 100
*/


/*func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelValues := []int{}
		if levelSize % 2 == 0 {
			i :=0
			j := levelSize -1
			for i < j {
				if queue[i].Val != queue[j].Val {
					return false
				}
				i ++
				j --
			}
		}

		for i := 0; i < levelSize; i++ {
			node := queue[i]
			levelValues = append(levelValues, node.Val)


				queue = append(queue, node.Left)


				queue = append(queue, node.Right)

		}
		result = append(result, levelValues)
		queue = queue[levelSize:]
	}
	return true
}*/

/*
上述队列思路错误
*/


/*

*/

// 解法1：基于队列迭代
func isSymmetric(root *TreeNode) bool {
	// 创建一个队列，用于存储需要比较的节点
	var queue []*TreeNode

	// 如果根节点不为空，将左子树和右子树的根节点加入队列
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}

	// 循环处理队列中的节点
	for len(queue) > 0 {
		// 从队列中取出左子树和右子树的根节点
		left := queue[0]
		right := queue[1]

		// 将已处理的节点从队列中移除
		queue = queue[2:]

		// 如果左右节点都为空，说明对称，继续下一轮比较
		if left == nil && right == nil {
			continue
		}

		// 如果左右节点其中一个为空，或者它们的值不相等，说明不对称，返回false
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		// 将左子树的左节点、右子树的右节点、右子树的左节点、左子树的右节点加入队列
		queue = append(queue, left.Left, right.Right, right.Left, left.Right)
	}

	// 遍历完所有节点后仍未返回false，说明二叉树是对称的，返回true
	return true
}



/*
如果一个树的左子树与右子树镜像对称，那么这个树是对称的。

因此，该问题可以转化为：两个树在什么情况下互为镜像？

如果同时满足下面的条件，两个树互为镜像：

它们的两个根结点具有相同的值
每个树的右子树都与另一个树的左子树镜像对称


我们可以实现这样一个递归函数，通过「同步移动」两个指针的方法来遍历这棵树，p 指针和 q 指针一开始都指向这棵树的根，随后 p 右移时，q 左移，p 左移时，q 右移。每次检查当前 p 和 q 节点的值是否相等，如果相等再判断左右子树是否对称。


*/

// 解法2：后序的思路进行递归
// 定义递归函数defs，用于判断两个节点是否对称
func defs(left *TreeNode, right *TreeNode) bool {
	// 如果两个节点都为空，说明对称
	if left == nil && right == nil {
		return true
	}
	// 如果其中一个节点为空，不对称
	if left == nil || right == nil {
		return false
	}
	// 如果两个节点的值不相等，不对称
	if left.Val != right.Val {
		return false
	}
	// 递归比较左子树的左节点与右子树的右节点，以及右子树的左节点与左子树的右节点
	// 外层的与外层的对比，内层的与内层对比，依次递归比较
	return defs(left.Left, right.Right) && defs(right.Left, left.Right)
}

// 定义主函数isSymmetric，用于检查二叉树是否对称
func isSymmetric(root *TreeNode) bool {
	// 调用递归函数defs，从根节点的左子树和右子树开始比较
	return defs(root.Left, root.Right)
}