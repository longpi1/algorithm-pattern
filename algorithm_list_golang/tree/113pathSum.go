package main
/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。

示例 1：
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]
示例 2：
输入：root = [1,2,3], targetSum = 5
输出：[]
示例 3：

输入：root = [1,2], targetSum = 0
输出：[]
*/

/*func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	valQueue := []int{root.Val}
	path := []int{root.Val}
	for len(queue) != 0 {
		n := len(queue)
		for i := 0; i <n ; i++ {
			node := queue[i]
			val := valQueue[i]
			if node.Left == nil && node.Right == nil && val == targetSum{
				result = append(result, path)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
				valQueue = append(valQueue,node.Left.Val+val)
				path = append(path,node.Left.Val)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				valQueue = append(valQueue,node.Right.Val+val)
				path = append(path,node.Right.Val)
			}
		}
		path = path[n:]
		valQueue = valQueue[n:]
		queue = queue[n:]
	}
	return result
}*/

/*
上述代码错误：
函数的主要目标是查找二叉树中从根节点到叶子节点的路径，使得路径上节点值之和等于给定的 targetSum。尽管代码的整体逻辑基本正确，但它有一个重要的问题，会导致结果不正确：
问题在于 result 切片中存储的 path 切片是一个引用，而不是一个复制。这意味着当你将 path 添加到 result 时，实际上只是将一个引用添加到了 result 中，而不是 path 的副本。因此，后续修改 path 的操作会影响已经添加到 result 中的路径。
为了解决这个问题，你应该将 path 的复制添加到 result 中，而不是直接添加引用。你可以使用内置的 append 函数来创建 path 的一个副本，然后将副本添加到 result。以下是修改后的代码：
*/

//基于广度遍历
func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	pathQueue := [][]int{{root.Val}}

	for len(queue) > 0 {
		node := queue[0]
		path := pathQueue[0]

		queue = queue[1:]
		pathQueue = pathQueue[1:]

		if node.Left == nil && node.Right == nil && sum(path) == targetSum {
			result = append(result, path)
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
			newPath := make([]int, len(path))
			copy(newPath, path)
			newPath = append(newPath, node.Left.Val)
			pathQueue = append(pathQueue, newPath)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
			newPath := make([]int, len(path))
			copy(newPath, path)
			newPath = append(newPath, node.Right.Val)
			pathQueue = append(pathQueue, newPath)
		}
	}

	return result
}
func sum(path []int) int {
	s := 0
	for _, val := range path {
		s += val
	}
	return s
}

// 基于回溯
func pathSum(root *TreeNode, target int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)

	if root == nil {
		return ans
	}

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, nowSum int) {
		nowSum += root.Val
		path = append(path, root.Val)

		if nowSum == target && root.Left == nil && root.Right == nil {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			ans = append(ans, pathCopy)
			// ans = append(ans, append([]int{}, path...))
		}

		if root.Left != nil {
			dfs(root.Left, nowSum)
			path = path[0:len(path) - 1]
		}
		if root.Right != nil {
			dfs(root.Right, nowSum)
			path = path[0:len(path) - 1]
		}
		if nowSum >= target {
			return
		}
	}
	dfs(root, 0)
	return ans
}
