package main

import (
	"fmt"
	"sort"
)

/*
90. 子集 II
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。



示例 1：

输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
示例 2：

输入：nums = [0]
输出：[[],[0]]


提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10
*/

// 错误代码 下述代码剪枝逻辑错误
func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	path := make([]int, 0)

	dfs := func(index int) {}
	dfs = func(index int) {
		//if len(path) <= len(nums) {
		fmt.Printf("path: %v", path)
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		//return
		//}
		for i := index; i < len(nums); i++ {
			if i == 0 {
				path = append(path, nums[i])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
			if i > 0 && nums[i] != nums[i-1] {
				path = append(path, nums[i])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
			// 正确写法如下
			/*
				// 剪枝/去重逻辑 (最关键的修正点)
					// 如果当前元素 nums[i] 和前一个元素 nums[i-1] 相同，
					// 并且 i > index（表示 nums[i-1] 是在当前递归层级中被跳过的），
					// 那么跳过当前元素 nums[i]，以避免生成重复的子集。
					// 举例：nums = [1, 2, 2]，index = 1
					// i=1，nums[1]=2。i > index (1 > 1) 不成立，不会触发去重。选择 2。path=[...,2]
					// i=2，nums[2]=2。i > index (2 > 1) 成立，nums[2] == nums[1] (2 == 2) 成立。
					// 此时跳过 nums[2]，因为如果选择它，会生成与通过 nums[1] 得到的重复组合。
					// 这个条件确保了对于重复元素，我们只在第一个位置选择它作为子集的起始元素。
					if i > index && nums[i] == nums[i-1] {
						continue // 跳过这个重复元素
					}
			*/
		}
	}
	dfs(0)
	return result
}

// subsetsWithDup 找出所有可能的子集，其中 nums 可能包含重复元素
func subsetsWithDup(nums []int) [][]int {
	// 1. 特殊情况处理：如果输入为空数组，返回包含一个空切片的二维切片 (即空集的子集是空集本身)
	if len(nums) == 0 {
		return [][]int{{}} // 返回 [][]int{} 也可以，但根据题目要求，空集的子集是 []
	}

	// 2. 核心：首先对输入数组进行排序。这是处理重复元素并避免生成重复组合的关键步骤。
	sort.Ints(nums)

	result := make([][]int, 0) // 存储所有找到的有效子集
	path := make([]int, 0)     // path 存储当前正在构建的子集

	// dfs 是回溯函数
	// index: 当前从 nums 数组的哪个位置开始考虑选择元素
	var dfs func(index int)

	dfs = func(index int) {
		// 1. 每次进入 dfs 函数时，path 都代表一个有效的子集（包括空集），将其加入结果集。
		// 注意：必须进行深拷贝，否则 result 中的所有子集都将指向同一个底层数组，
		// 随着 path 的修改而变化，导致结果不正确。
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)

		// 2. 遍历所有可能的选择 (从当前 index 开始，到 nums 的末尾)
		for i := index; i < len(nums); i++ {
			// 3. 剪枝/去重逻辑 (最关键的修正点)
			// 如果当前元素 nums[i] 和前一个元素 nums[i-1] 相同，
			// 并且 i > index（表示 nums[i-1] 是在当前递归层级中被跳过的），
			// 那么跳过当前元素 nums[i]，以避免生成重复的子集。
			// 举例：nums = [1, 2, 2]，index = 1
			// i=1，nums[1]=2。i > index (1 > 1) 不成立，不会触发去重。选择 2。path=[...,2]
			// i=2，nums[2]=2。i > index (2 > 1) 成立，nums[2] == nums[1] (2 == 2) 成立。
			// 此时跳过 nums[2]，因为如果选择它，会生成与通过 nums[1] 得到的重复组合。
			// 这个条件确保了对于重复元素，我们只在第一个位置选择它作为子集的起始元素。
			// 在同一层递归中，如果当前元素和前一个元素相同，并且前一个元素已经被跳过（不是作为当前组合的一部分），那么当前元素也要跳过。
			// 正确处理重复元素的关键在于 if i > index && nums[i] == nums[i-1] { continue }。这里的 i > index 是为了区分“重复使用”和“跳过重复元素”。当 i == index 时，说明是当前递归层级中，
			// 从 index 开始的第一个选择，即使 nums[index] == nums[index-1] (如果 index > 0) 也应该处理，因为 nums[index-1] 是在上一层递归中被考虑的。
			//但是，如果 i > index 且 nums[i] == nums[i-1]，那么 nums[i-1] 已经在当前层级被 for 循环迭代过了，如果 nums[i-1] 没被选择，那 nums[i] 作为重复元素就也不应该被选择，否则会生成重复组合。
			if i > index && nums[i] == nums[i-1] {
				continue // 跳过这个重复元素
			}

			// 4. 做选择 (将 nums[i] 加入当前子集)
			path = append(path, nums[i])

			// 5. 递归调用，进入下一层决策
			// 下一次从 i+1 开始考虑，确保每个元素最多被使用一次，且维持组合的递增顺序
			dfs(i + 1)

			// 6. 撤销选择 (回溯)
			// 将 nums[i] 从 path 中移除，回到上一个决策点
			path = path[:len(path)-1]
		}
	}

	// 初始调用回溯函数，从 nums 的第一个元素 (索引 0) 开始
	dfs(0)
	return result
}

func main() {
	nums := []int{1, 2, 2}
	result := subsetsWithDup(nums)
	fmt.Printf("result: %v", result)
}
