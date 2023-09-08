package main


/*
给定一个二叉树 root ，返回其最大深度。

二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。



示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：3
示例 2：

输入：root = [1,null,2]
输出：2

*/

/*func maxDepth(root *TreeNode) int {
	depth := 0
	dfs := func(root  *TreeNode) {}
	dfs = func(root *TreeNode) {
		if root == nil{
			return
		}
		depth += 1
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return depth
}
*/
/*
上述思路错误：
1.首先，深度计算的方式是不正确的。在 dfs 函数中，你只是累加了一个变量 depth，但没有正确地计算深度。深度应该表示从根节点到当前节点的深度，但你的代码只是简单地累加，没有考虑深度的正确计算方式。
2.另外，深度优先搜索的方式也不正确。你在 dfs 函数中定义了一个匿名函数，但这个匿名函数没有正确地递归调用左子树和右子树。你需要传递深度信息，并在递归调用中更新深度。
*/
// 基于后序遍历思想做此题
/*
1. 如何思考二叉树相关问题？
    - 不要一开始就陷入细节，而是思考整棵树与其左右子树的关系。
2. 为什么需要使用递归？
    - 子问题和原问题是相似的，他们执行的代码也是相同的（类比循环），但是子问题需要把计算结果返回给上一级，这更适合用递归实现。
3. 为什么这样写就一定能算出正确答案？
    - 由于子问题的规模比原问题小，不断“递”下去，总会有个尽头，即递归的边界条件 ( base case )，直接返回它的答案“归”；
    - 类似于数学归纳法（多米诺骨牌），n=1时类似边界条件；n=m时类似往后任意一个节点
4. 计算机是怎么执行递归的？
    - 当程序执行“递”动作时，计算机使用栈保存这个发出“递”动作的对象，程序不断“递”，计算机不断压栈，直到边界时，程序发生“归”动作，正好将执行的答案“归”给栈顶元素，随后程序不断“归”，计算机不断出栈，直到返回原问题的答案，栈空。
5. 另一种递归思路
    - 维护全局变量，使用二叉树遍历函数，不断更新全局变量最大值。
递归相关讲解视频：https://www.bilibili.com/video/BV1UD4y1Y769/?vd_source=34718180774b041b23050c8689cdbaf2
*/
func maxDepth(root *TreeNode) int {
	// 辅助函数，用于计算深度
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, depth int) int {
		// 边界条件
		if node == nil {
			return depth // 当节点为空时返回深度
		}
		// 递归计算左子树和右子树的深度，取较大值
		leftDepth := dfs(node.Left, depth+1)
		rightDepth := dfs(node.Right, depth+1)
		// 中
		max := max(leftDepth, rightDepth)
		return max
	}

	// 调用深度计算函数，从根节点开始计算深度
	return dfs(root, 0)
}

// 辅助函数，用于比较两个整数的最大值
func max(a int, b int) int{
	if a > b {
		return a
	}
	return b
}


// 基于队列实现此题：
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0

	for len(queue) != 0 {
		levelSize := len(queue)
		for i := 0; i< levelSize; i++ {

			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)

			}
		}
		depth ++
		queue = queue[levelSize:]
	}
	return depth
}

func main()  {
	root :=	&TreeNode{Val: 1,Left: &TreeNode{Val: 2},Right: &TreeNode{Val: 3,Left: &TreeNode{Val: 4}}}
	print(maxDepth(root))
}

