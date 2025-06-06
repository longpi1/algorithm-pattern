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
/*
下述代码可能存在的问题：
1.处理 nil 节点的方式错误：

问题描述： 在 isSymmetric 函数中，你只将 node.Left 和 node.Right 不为 nil 的节点添加到 tmp 队列中：
if node.Left != nil {
    tmp = append(tmp, node.Left)
}
if node.Right != nil {
    tmp = append(tmp, node.Right)
}

这意味着当一个节点有 nil 子节点时，这些 nil 节点不会被记录到下一层。然而，对于判断对称性，nil 节点的位置和数量是至关重要的。例如，如果左子树的某个位置是 nil，那么右子树对应位置也必须是 nil 才能对称。你的代码会过滤掉 nil，导致 judge 函数拿到的 tmp 数组不包含 nil 元素，从而无法判断结构上的对称性。
例子：
    1
   / \
  2   2
 /     \
3       3
这个树不是对称的（左 2 只有左孩子，右 2 只有右孩子）。 但你的代码在处理完 2 和 2 这一层后，tmp 可能会变成 [3, 3] (忽略了中间的 nil 节点)。judge([3, 3]) 会返回 true，导致误判。
judge 函数对 nil 节点不安全：

2.问题描述： 紧接着第一个问题，如果 isSymmetric 修正为将 nil 节点也加入 tmp 队列，那么 judge 函数在访问 arr[i].Val 时，可能会遇到 nil 指针，从而导致 panic (nil pointer dereference)。
例子： 如果 tmp 变为 [3, nil, nil, 3]，当 judge 访问 arr[1].Val 时就会崩溃。
levelSize 的计算和循环范围错误：

3.问题描述： levelSize := len(queue) - 1 这个计算是错误的。如果 len(queue) 是 1，那么 levelSize 变成 0，内层循环 for i := 0; i < levelSize; i++ 就不会执行，这意味着 root 的子节点永远不会被加入 tmp，代码只会处理根节点（如果根节点有子节点的话），然后就结束了。正确的 levelSize 应该是 len(queue)。
*/
func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		tmp := make([]*TreeNode, 0)
		levelSize := len(queue) - 1
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		if !judge(tmp) {
			return false
		}
		queue = tmp

	}
	return true
}

func judge(arr []*TreeNode) bool {
	j := len(arr) - 1
	for i := 0; i < (len(arr)-1)/2; i++ {
		if arr[i].Val != arr[j].Val {
			return false
		}
		j--
	}
	return true
}

// 修正后的代码
// isSymmetric 函数用于判断二叉树是否对称
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true // 空树被认为是对称的
	}

	// 使用队列进行广度优先搜索 (BFS)
	// 队列中存储的是成对的节点，用于比较它们是否互为镜像
	// 初始时，将根节点的左子节点和右子节点入队
	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		// 每次从队列中取出两个节点进行比较
		// 注意：每次迭代都必须取出两个节点，因为它们是成对入队的
		node1 := queue[0]
		node2 := queue[1]
		queue = queue[2:] // 从队列中移除已取出的两个节点

		// 1. 如果两个节点都是 nil，它们是镜像的，继续检查下一对
		if node1 == nil && node2 == nil {
			continue
		}

		// 2. 如果其中一个为 nil，另一个不为 nil，则不对称
		// 或如果两个节点的值不相等，则不对称
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}

		// 3. 如果两个节点都非空且值相等，则将它们的子节点按镜像顺序入队
		// 左子节点的左孩子应该和右子节点的右孩子比较
		queue = append(queue, node1.Left)
		queue = append(queue, node2.Right)

		// 左子节点的右孩子应该和右子节点的左孩子比较
		queue = append(queue, node1.Right)
		queue = append(queue, node2.Left)
	}

	// 如果所有节点都通过了检查，则树是对称的
	return true
}

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
