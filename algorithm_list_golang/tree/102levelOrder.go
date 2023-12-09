package main


/*
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。



示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]
*/

type TreeNode1 struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue1 := make([]*TreeNode, 0)
	queue = append(queue, root)
	tmp := make([]int, 0)
	if root == nil {
		return [][]int{}
	}
	for len(queue) != 0 {
		node := queue[0]
		if node.Left != nil {
			queue1 = append(queue1, node.Left)
		}
		if node.Right != nil {
			queue1 = append(queue1, node.Right)
		}
		tmp = append(tmp, node.Val)
		queue = queue[1:]

		if len(queue) == 0 {
			result = append(result, tmp)
			tmp = []int{}
			queue = queue1
			queue1 = []*TreeNode{}
		}
	}
	return result
}

//优化后的代码
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelValues := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[i]
			levelValues = append(levelValues, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelValues)
		queue = queue[levelSize:]
	}

	return result
}


func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}



func main(){
	root :=	TreeNode{Val: 1,Left: &TreeNode{Val: 2},Right: &TreeNode{Val: 3,Left: &TreeNode{Val: 4}}}
	levelOrder(&root)
}