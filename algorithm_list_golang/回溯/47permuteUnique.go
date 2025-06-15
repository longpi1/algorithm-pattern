package main

import (
	"fmt"
	"sort"
)

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。



示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]


提示：

1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/

// 下述代码存在index 参数的错误使用和缺少 used 数组：
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	path := make([]int, 0)
	dfs := func(index int) {}
	dfs = func(index int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return result
}

// permuteUnique 返回包含重复数字的数组的所有不重复全排列。
func permuteUnique(nums []int) [][]int {
	// 1. 对数组进行排序。这是处理重复元素并避免生成重复排列的关键步骤。
	sort.Ints(nums)

	result := make([][]int, 0)        // 存储所有找到的有效排列
	path := make([]int, 0, len(nums)) // path 存储当前正在构建的排列，预分配容量

	// used 数组：标记 nums 中每个元素是否已被当前 path 使用。
	// 长度与 nums 相同，初始全部为 false。
	used := make([]bool, len(nums))

	// dfs 是回溯函数
	// currentPath: 当前构建的排列 (通常不需要作为参数传递，直接使用外层 path 变量)
	// used: 标记哪些元素已被使用
	var dfs func()

	dfs = func() {
		// 1. 终止条件 / 找到解的条件
		// 当当前排列的长度等于原始数组的长度时，找到一个有效排列
		if len(path) == len(nums) {
			// 注意：必须对 path 进行深拷贝，否则 result 中的所有排列都将指向同一个底层数组，
			// 随着 path 的修改而变化，导致结果不正确。
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return // 找到一个解后，当前分支结束
		}

		// 2. 遍历所有可能的选择 (从 nums 的起始位置开始遍历)
		for i := 0; i < len(nums); i++ {
			// 3. 剪枝条件 (最重要的修正点之一)
			// a. 如果 nums[i] 已经被使用过，则跳过
			if used[i] {
				continue
			}
			// b. 处理重复元素：
			// 如果当前元素 nums[i] 与前一个元素 nums[i-1] 相同，
			// 并且 nums[i-1] 在当前层递归中**没有被使用过** (即 used[i-1] 为 false)。
			// 那么说明 nums[i-1] 已经被跳过了，如果再选择 nums[i]，
			// 就会生成与通过 nums[i-1] 能够生成的重复排列。
			// 举例：nums = [1, 1, 2]
			// - 当选择第一个 1 (nums[0]) 时：没问题。
			// - 当考虑选择第二个 1 (nums[1]) 时：
			//   如果 nums[0] 已经被使用 (used[0]为true)，那么 nums[1] 就可以被选 (因为已经开始了新的一层)。
			//   如果 nums[0] 没有被使用 (used[0]为false)，说明 nums[0] 在此层被跳过了。
			//   此时如果选择 nums[1]，将会导致和选择 nums[0] 产生重复的排列。
			//   所以我们跳过 nums[1]。
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// 4. 做选择
			path = append(path, nums[i])
			used[i] = true // 标记 nums[i] 已被使用

			// 5. 递归调用，进入下一层决策
			// 注意：这里不需要传递 index 参数，因为我们每次都从 0 开始遍历所有未使用的数字。
			dfs()

			// 6. 撤销选择 (回溯)
			used[i] = false           // 恢复 nums[i] 的使用状态
			path = path[:len(path)-1] // 从排列中移除
		}
	}

	// 初始调用回溯函数
	dfs()
	return result
}

func permuteUnique(nums []int) [][]int {
	var result [][]int
	var path []int
	// 题目中给的不一定有序，首先进行排序，否则nums[i-1] == nums[i]判断不成立，无法去重
	sort.Ints(nums)
	n := len(nums)
	used := make(map[int]bool, n)
	dfs := func() {}
	dfs = func() {
		if len(path) == n {
			tmp := make([]int, n)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if !used[i] {
				// 去重操作
				// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
				// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
				// 一个是树枝一个是树层！！！
				if i > 0 && nums[i-1] == nums[i] && used[i-1] == false {
					continue
				}
				used[i] = true
				path = append(path, nums[i])
				dfs()
				used[i] = false
				path = path[:len(path)-1]
			}
		}

	}
	dfs()
	return result
}

func main() {
	nums := []int{3, 3, 0, 3}
	fmt.Printf("result:%v", permuteUnique(nums))
}
