package main

import "fmt"

/*
103. 二叉树的锯齿形层序遍历
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[20,9],[15,7]]

示例 2：
输入：root = [1]
输出：[[1]]

示例 3：
输入：root = []
输出：[]
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
下述代码犯了两个关键错误：

子节点入队顺序与结果顺序的矛盾：

之字形遍历要求的是输出结果在奇数/偶数层是反向的。
而遍历过程本身，即下一层节点的处理顺序，永远应该是从左到右的。我们只是在收集当前层的值时，决定是从左到右取还是从右到左取。
你的代码试图通过改变下一层节点的入队顺序（先右后左）来影响当前层的输出，这是行不通的。队列的 FIFO 特性决定了，你先放 node.Right 进队，它就会在下一轮中先被处理。这会彻底打乱整个树的层级结构。
在遍历当前层时修改队列：

这是一个非常危险的操作。你的 for 循环是 for i := 0; i < size; i++，其中 size 是当前层的节点数。
在循环内部，你又执行了 queue = append(queue, ...)。这会不断地向 queue 的末尾添加新元素。
虽然循环条件 i < size 保证了你只会处理当前层的 size 个节点，但这种边遍历边修改同一个数据结构的做法非常容易出错，并且不符合队列处理的清晰逻辑。正确的做法是，处理完当前层的 size 个节点后，队列中剩下的就自然是下一层的全部节点。

	    1
	   / \
	  2   3
	 / \   \
	4   5   6

正确结果：[[1], [3, 2], [4, 5, 6]]

追踪你的代码：

第 0 层 (index=0, flag=true, 从左到右)

处理 1，tmp = [1]。
flag为true，先入队 1.Right(3)，后入队 1.Left(2)。
queue 变为 [3, 2]。
result 添加 [1]。
第 1 层 (index=1, flag=false, 从右到左)

queue 是 [3, 2]，size = 2。
flag 为 false。
处理节点 3:
tmp 变为 [3]。
flag为false，先入队 3.Left(nil)，后入队 3.Right(6)。queue 变为 [3, 2, 6]。
处理节点 2:
tmp 变为 [3, 2]。
flag为false，先入队 2.Left(4)，后入队 2.Right(5)。queue 变为 [3, 2, 6, 4, 5]。
queue = queue[size:]：queue 变为 [6, 4, 5]。
result 添加 [3, 2]。
第 2 层 (index=2, flag=true, 从左到右)

queue 是 [6, 4, 5]。
...
tmp 将会是 [6, 4, 5]。
最终结果：[[1], [3, 2], [6, 4, 5]] 你的结果是错误的！ 因为下一层的处理顺序被打乱了。正确的第三层应该是 [4, 5, 6]。
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := []*TreeNode{root}
	index := 0
	for len(queue) != 0 {
		size := len(queue)
		tmp := make([]int, 0, size)
		flag := (index % 2) == 0
		fmt.Println(flag)
		for i := 0; i < size; i++ {
			node := queue[i]
			tmp = append(tmp, node.Val)
			// 主要错误点
			if !flag {
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			} else {
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
			}
		}
		queue = queue[size:]
		index++
		result = append(result, tmp)
	}
	return result
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	queue := []*TreeNode{root}
	leftToRight := true // true 表示当前层从左往右输出

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			// 根据方向填充 level
			if leftToRight {
				level[i] = node.Val
			} else {
				level[size-1-i] = node.Val // 关键：反着写
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level)
		leftToRight = !leftToRight // 方向翻转
	}
	return res
}

func main() {
	/*
			root =
			[1,2,3,4,null,null,5]
		[[1],[3,2],[4,5]]
	*/
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}
	order := zigzagLevelOrder(root)
	fmt.Printf("result:%v", order)
}
