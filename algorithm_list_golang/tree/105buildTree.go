package main

/*
105. 从前序与中序遍历序列构造二叉树

给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

示例 1:
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]

示例 2:
输入: preorder = [-1], inorder = [-1]
输出: [-1]

*/

/*
解题思路：
第一步：如果数组大小为零的话，说明是空节点了。

第二步：如果不为空，那么取前序数组第一个元素作为节点元素。

第三步：找到前序数组第一个元素在中序数组的位置，作为切割点

第四步：切割中序数组，切成中序左数组和中序右数组 （顺序别搞反了，一定是先切中序数组）

第五步：切割前序数组，切成前序左数组和前序右数组

第六步：递归处理左区间和右区间
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	//根节点
	node := &TreeNode{Val: preorder[0]}
	index := 0
	for i := index; i < len(preorder); index++ {
		if inorder[index] == preorder[0] {
			break
		}
	}
	// 分割中序
	leftInorder := inorder[:index]
	rightInorder := inorder[index+1:]
	// 分割前序
	leftPreorder := preorder[1 : index+1]
	rightPreorder := preorder[index+1:]
	node.Left = buildTree(leftPreorder, leftInorder)
	node.Right = buildTree(rightPreorder, rightInorder)
	return node
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	buildTree(preorder, inorder)
}
