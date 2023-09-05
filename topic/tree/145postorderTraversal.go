package main

/*
给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。



示例 1：


输入：root = [1,null,2,3]
输出：[3,2,1]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 实现1：递归
func postorderTraversal(root *TreeNode) []int {

	result := make([]int,0)

	return traversal(root,result)
}

func traversal(cur *TreeNode, result []int) []int{
	if cur == nil {
		return result
	}

	result=traversal(cur.Left,result)

	result=traversal(cur.Right,result)
	result = append(result, cur.Val)
	return result
}


// 非递归遍历
/*func postorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	stack = append(stack, root)

	if root == nil{
		return []int{}
	}
	m := make(map[*TreeNode]bool)
	for len(stack) != 0 {
		// !!! 定义对应的接收栈
		node := stack[len(stack)-1]
		m[root] = true
		if !m[node.Right] && node.Right != nil {
			stack = append(stack, node.Right)
			m[node.Right] = true
		}
		if !m[node.Left] && node.Left != nil {
			stack = append(stack, node.Left)
			m[node.Left] = true
		}
		node = stack[len(stack)-1]
		result = append(result, node.Val)
		stack = stack[:len(stack)-1]
	}
	return result
}*/
/*
上述迭代法实现后序遍历错误：
1.使用了一个 map（m）来记录节点是否已经访问过。这本身不是错误，但是在遍历的过程中，你应该将节点 node 添加到 map 中，而不是 root，因为你想要跟踪每个节点是否被访问过。
2.在入栈右子节点和左子节点时，检查了 !m[node.Right] 和 !m[node.Left]，这样会导致在访问某个节点时，如果它的子节点已经被访问过，那么它的子节点将不会再次入栈，这可能会导致遗漏一些节点。
*/
// 实现2：迭代
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{}
	result := []int{}
	visited := map[*TreeNode]bool{}

	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		// 检查左子节点是否已访问，如果未访问，将左子节点入栈
		if node.Left != nil && !visited[node.Left] {
			stack = append(stack, node.Left)
		} else if node.Right != nil && !visited[node.Right] {
			// 检查右子节点是否已访问，如果未访问，将右子节点入栈
			stack = append(stack, node.Right)
		} else {
			// 左右子节点都已访问或为空，将当前节点出栈并记录到结果中
			result = append(result, node.Val)
			// 在入栈右子节点和左子节点时，检查了 !m[node.Right] 和 !m[node.Left]，这样会导致在访问某个节点时，如果它的子节点已经被访问过，那么它的子节点将不会再次入栈，这可能会导致遗漏一些节点。
			visited[node] = true
			stack = stack[:len(stack)-1]
		}
	}

	return result
}


// 实现3，使用反转实现的代码如下：
func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	stack1 := []*TreeNode{} // 用于节点的遍历
	stack2 := []*TreeNode{} // 用于逆序保存结果

	stack1 = append(stack1, root)

	for len(stack1) > 0 {
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]

		stack2 = append(stack2, node)

		if node.Left != nil {
			stack1 = append(stack1, node.Left)
		}
		if node.Right != nil {
			stack1 = append(stack1, node.Right)
		}
	}

	for len(stack2) > 0 {
		node := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		result = append(result, node.Val)
	}

	return result
}


func main(){
	root :=	TreeNode{Val: 1,Right: &TreeNode{Val: 3,Left: &TreeNode{Val: 2}}}
	postorderTraversal(&root)
}