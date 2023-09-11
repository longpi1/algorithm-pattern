package main

/*
114. 二叉树展开为链表

给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。


示例 1：
输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [0]
输出：[0]
*/

/*
在代码中，传递了 result 切片的指针 *[]int，因为你要在 preTraversal 函数中修改 result 切片的内容。虽然数组是引用类型，但在 Go 中，切片是动态数组，它的长度可以动态增长，而数组的长度是固定的。
因此，如果你希望在函数内部添加元素到切片中，你需要传递切片的指针。
当你传递 result 切片的指针 *[]int 给 preTraversal 函数时，你可以在函数内部通过 *result 来访问和修改原始切片。如果你只传递 result 切片本身，那么在函数内部对切片的任何修改都会在函数退出后被丢弃，因为切片是按值传递的。
总之，传递切片的指针允许你在函数内部修改原始切片的内容，而不仅仅是操作切片的副本。这对于在 preTraversal 函数中构建一个前序遍历的结果列表非常重要，因此需要传递切片的指针。
Go 中的函数参数都是按值传递的，这意味着在函数内部对参数的修改不会影响原始变量的值，这是一种保护性的特性，可以防止函数无意中更改了原始数据。当你传递一个切片、数组、map 或结构体等引用类型时，仍然是按值传递的，但在函数内部操作引用类型时，实际上是在操作原始数据的副本，而不是原始数据本身。
这就是为什么在递归函数中，如果你想要修改原始引用类型的内容（比如在二叉树的遍历中修改节点的值），你需要传递指向引用类型的指针，而不仅仅是引用类型本身。通过传递指针，你可以在函数内部直接访问和修改原始数据，而不是操作它的副本。
在你的代码中，你传递了 result 切片的指针 *[]int，这样在 preTraversal 函数中可以通过 *result 直接访问和修改原始切片，而不仅仅是对切片的副本进行操作。这确保了递归函数可以正确地构建前序遍历的结果列表，因为它可以在函数内部修改原始切片的内容。
*/
func flatten(root *TreeNode)  {

	result := make([]int, 0)
	preTraversal(root,&result)
	newTree := root
	for i := 0; i < len(result); i++ {
		newTree.Val = result[i]
		newTree.Left = nil
		if i != len(result) -1 && newTree.Right == nil {
			newTree.Right = &TreeNode{}
		}
		newTree = newTree.Right
	}
}

func preTraversal(root *TreeNode,result *[]int)  {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	preTraversal(root.Left,result)
	preTraversal(root.Right,result)
}

func main() {
	root :=	&TreeNode{Val: 5,Left: &TreeNode{Val: 4},Right: &TreeNode{Val: 6,Left: &TreeNode{Val: 3},Right: &TreeNode{Val: 7}}}
	flatten(root)
}

