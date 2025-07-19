package main

/*
将有序数组转换为二叉搜索树

给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

示例 1：
输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

示例 2：
输入：nums = [1,3]
输出：[3,1]
解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。

*/

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: mid} // !!! BUG: 应填 nums[mid]，而不是下标 mid

	// dfs 递归变量先占位再重新赋值——Go 闭包写法 OK
	dfs := func(nums []int, node *TreeNode) {}
	dfs = func(nums []int, node *TreeNode) {
		if len(nums) == 0 {
			return
		}

		mid = len(nums) / 2 // !!! BUG 1: 复用最外层同名变量，
		//           多层递归相互覆盖，结果错乱
		// !!! BUG 2: 下面所有地方都用“全局 root”，
		//           而没有使用当前节点 node，子树会被写进根节点
		//           并且 Left/Right 还没分配直接访问会 panic
		root.Left.Val = nums[mid-1]  // !!! BUG 3: root.Left 为空指针，直接 .Val 会崩溃
		root.Right.Val = nums[mid+1] // !!! BUG 4: root.Right 为空指针，同上

		// !!! BUG 5: 切片下标越界风险
		//           当 mid==0 时 nums[:mid-1] 负数下标；mid==len(nums)-1 时 mid+1 溢出
		dfs(nums[:mid-1], root.Left)
		dfs(nums[mid+1:], root.Right)
	}

	dfs(nums, root)
	return root
}

/*func sortedArrayToBST(nums []int) *TreeNode {
	left := 0
	right := len(nums) -1
	root := &TreeNode{}
	dfs := func(root *TreeNode,left int ,right int)  {

	}
	dfs = func(root *TreeNode, left int, right int)  {
		if left > right {
			return
		}
		// 存在一个隐患，不过这里是数组不需要考虑，通过使用 left + (right-left)/2 而不是 (left+right)/2 来避免整数溢出。
		mid := (left+right) /2
		root.Val = nums[mid]
		dfs(root.Left,left,mid-1)
		dfs(root.Right,mid+1,right)
	}
	dfs(root, left,right)
	return root
}
*/

/*
上述代码存在空指针问题：
未创建子节点的节点：在 sortedArrayToBST 函数中创建了根节点 root，但没有为它的左右子节点分配内存。你需要在需要创建子节点的时候分配内存。
*/

func sortedArrayToBST(nums []int) *TreeNode {
	// 确定初始左右边界
	left := 0
	right := len(nums) - 1

	// 定义递归函数dfs，初始情况下返回nil
	dfs := func(left, right int) *TreeNode {
		return nil
	}

	// 实际的递归函数定义
	dfs = func(left, right int) *TreeNode {
		if left > right {
			// 边界情况：左边界大于右边界，返回nil，表示没有节点
			return nil
		}
		// 存在一个隐患，不过这里是数组不需要考虑，通过使用 left + (right-left)/2 而不是 (left+right)/2 来避免整数溢出。
		// 计算当前子树的根节点索引
		mid := (left + right) / 2
		// 创建根节点，值为当前中间元素
		root := &TreeNode{Val: nums[mid]}
		// 递归构建左子树和右子树
		root.Left = dfs(left, mid-1)
		root.Right = dfs(mid+1, right)
		return root
	}

	// 调用dfs函数开始构建平衡二叉搜索树并返回根节点
	return dfs(left, right)
}

func main() {
	nums := []int{
		-10, -3, 0, 5, 9,
	}
	sortedArrayToBST(nums)
}
