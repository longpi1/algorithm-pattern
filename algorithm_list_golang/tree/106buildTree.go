package main

/*
106. 从中序与后序遍历序列构造二叉树

给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。

示例 1:
输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]

示例 2:
输入：inorder = [-1], postorder = [-1]
输出：[-1]

*/

/*
下述代码错误点：
index 变量命名混淆和作用域问题：
在找到 rootVal 在 inorder 中的位置后，index 应该就是它在该数组中的索引。
但是你又引入了一个 index2 变量，并在第二个循环中使用了 postorder[index]。这里的 index 仍然是 inorder 中 rootVal 的索引，而不是 postorder 数组中子树的正确分割点。这是一个严重的逻辑错误。

postorder 数组的分割逻辑错误：
后序遍历序列的最后一个元素是根节点。
在中序遍历序列中，根节点左边的所有元素属于左子树，右边的所有元素属于右子树。
关键在于，postorder 序列的分割点不是简单的对应 inorder 中的索引。左子树的后序遍历序列是 postorder 中除去根节点之外的前 len(left_subtree_inorder) 个元素。右子树的后序遍历序列是剩下的元素（在 postorder 中）。
你尝试用 index2 来找 postorder 中的分割点，但 if postorder[index] != inorder[index] 这个条件完全没有意义，而且 postorder[index] 可能会越界（因为 index 是 inorder 的索引，可能比 postorder 长度大）。

子数组的切片范围错误：
inorder[:index-1] 错误：如果 index 是 0 (根节点是中序遍历的第一个元素)，index-1 就是 -1，会导致 panic: slice bounds out of range。正确的左子树中序遍历应该是 inorder[:index]。
postorder[:index2] 错误：index2 的计算是错的，即使对了，这里的分割点也应该是由左子树的长度决定。
postorder[index2+1:n-1] 错误：同样，index2 错误导致这里也错。

边界条件处理不完善：
if n == 1 { return root }：这个条件应该放在找到 rootVal 并构建 root 之后。如果 n 是 0 并且已经处理了，它不会走到这里。但这个判断的意义是，如果只有一个节点，它就是根节点，直接返回。这个本身没有太大问题，但更规范的解法是在递归的基准情况中统一处理。
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	n := len(postorder)

	rootVal := postorder[n-1]
	root := &TreeNode{Val: rootVal}
	if n == 1 {
		return root
	}
	index := 0
	for index = 0; index < n; index++ {
		if inorder[index] != rootVal {
			continue
		}
		break
	}
	index2 := 0
	for index2 = 0; index2 < n; index2++ {
		if postorder[index] != inorder[index] {
			continue
		}
		break
	}
	root.Left = buildTree(inorder[:index-1], postorder[:index2])
	root.Right = buildTree(inorder[index+1:], postorder[index2+1:n-1])
	return root

}

// buildTree 根据中序遍历和后序遍历序列构建二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 1. 基准情况：如果序列为空，则无法构建节点，返回nil
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 2. 后序遍历的最后一个元素是当前子树的根节点
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// 3. 在中序遍历序列中找到根节点的值，以区分左右子树
	//    此处的 `idx` 是根节点在中序序列中的索引
	var idx int
	for i, val := range inorder {
		if val == rootVal {
			idx = i
			break
		}
	}

	// 4. 递归构建左右子树
	// 左子树的中序遍历序列：inorder[:idx]
	// 左子树的后序遍历序列：postorder[:idx] （长度与中序左子树相同）
	root.Left = buildTree(inorder[:idx], postorder[:idx])

	// 右子树的中序遍历序列：inorder[idx+1:]
	// 右子树的后序遍历序列：postorder[idx : len(postorder)-1] （从idx开始到倒数第二个元素结束）
	root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])

	return root
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	//根节点
	node := &TreeNode{Val: postorder[n-1]}
	index := 0
	for i := index; i < n; index++ {
		if inorder[index] == postorder[n-1] {
			break
		}
	}
	// 分割后序
	leftPostorder := postorder[:index]
	rightPostorder := postorder[index:n]
	// 分割中序
	leftInorder := inorder[:index]
	rightInorder := inorder[index+1:]

	node.Left = buildTree(leftInorder, leftPostorder)
	node.Right = buildTree(rightInorder, rightPostorder)
	return node
}

func main() {
	postorder := []int{9, 15, 7, 20, 3}
	inorder := []int{9, 3, 15, 20, 7}
	buildTree(inorder, postorder)
}
