package main
/*
//给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
// 示例 1:
//输入: nums = [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
// 示例 2:
//输入: nums = [2,3,0,1,4]
//输出: 2
// 提示:
// 1 <= nums.length <= 104
// 0 <= nums[i] <= 1000
// 题目保证可以到达 nums[n-1]
//
// Related Topics 贪心 数组 动态规划
*/
/*
解题思路：
这道题是典型的贪心算法，通过局部最优解得到全局最优解。以下两种方法都是使用贪心算法实现，只是贪心的策略不同。
*/

// 解法1正向
func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0 // 如果数组只有一个元素，直接返回 true，无需跳跃
	}

	// 初始化覆盖范围
	cover := 0
	maxCover := 0
	// 初始化跳跃次数
	count := 0
	for i := 0; i <= maxCover; i++ {
		// 更新覆盖范围，选择能够跳到更远位置的情况
		if cover < i+nums[i] {
			cover = i + nums[i]
		}
		if i == maxCover {
			maxCover = cover
			count ++
		}

		// 如果已经可以覆盖数组的最后一个位置，返回 count
		if maxCover >= n-1 {
			return count
		}
	}

	return count // 无法跳跃到数组的最后一个位置，返回 count
}


// 解法2反向
func jump(nums []int) int {
	position := len(nums) - 1
	steps := 0
	for position > 0 {
		for i := 0; i < position; i++ {
			if i + nums[i] >= position {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}



func main(){
	nums := []int{7,0,9,6,9,6,1,7,9,0,1,2,9,0,3}
	jump(nums)
}