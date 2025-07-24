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
//    从下标为 0 跳到下标为 1 的位置，跳1步，然后跳3步到达数组的最后一个位置。
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

package main

// 思路错误
func jump(nums []int) int {
	// 边界条件处理正确，如果数组长度小于等于1，不需要跳跃。
	if len(nums) <= 1 {
		return 0
	}

	result := 0
	// `cover` 用来记录能跳到的最远距离，这个变量定义是正确的。
	cover := 0

	// --- 问题 1: 循环的边界和目的不清晰 ---
	// for i := 0; i < len(nums)-1; i++ {
	// 这个循环的目的是什么？是遍历每一步，还是遍历每一次跳跃？
	// 贪心算法的核心是：在当前跳跃范围内，找到下一次跳跃能达到的最远位置。
	// 你的循环是遍历数组的每一个位置，这偏离了“按步跳跃”的核心。
	for i := 0; i < len(nums)-1; i++ {

		// --- 问题 2: `result++` 的时机完全错误 ---
		// if cover < i+nums[i] {
		//	   cover = i + nums[i]
		//	   result++
		// }
		// 这是整个代码最核心的错误。
		// 它的意思是：“只要我发现了一个能跳得更远的位置，我就增加一次跳跃次数”。
		//
		// 让我们用一个例子来说明：nums = [2, 3, 1, 1, 4]
		//
		// - i = 0: `nums[0]=2`。`cover` 从 0 更新为 `0+2=2`。`result` 变成 1。
		//          (这里已经错了，我们还没决定跳，只是在探索)
		// - i = 1: `nums[1]=3`。`i+nums[i]` 是 `1+3=4`。`4 > cover(2)`，
		//          `cover` 更新为 4。`result` 变成 2。
		//          (我们明明还在第一步的覆盖范围内探索，怎么能又算一次跳跃呢？)
		// - i = 2: `nums[2]=1`。`i+nums[i]` 是 `2+1=3`。`3 < cover(4)`，什么都不做。
		// - i = 3: `nums[3]=1`。`i+nums[i]` 是 `3+1=4`。`4 == cover(4)`，什么都不做。
		//
		// 最终 `result` 会是一个很大的、错误的值。
		//
		// 正确的逻辑是：只有在**当前这一步的覆盖范围已经走完，必须进行下一次跳跃时**，才增加 `result`。
		if cover < i+nums[i] {
			cover = i + nums[i]
			result++
		}

		// 这个判断是正确的，如果当前能覆盖到终点，就返回结果。
		if cover >= len(nums)-1 {
			return result
		}
	}

	// 根据题目保证“总能到达最后一个位置”，所以这个返回值理论上不会被触发。
	// 但写上它作为防御性编程是好的。
	return -1
}

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
			count++
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
			if i+nums[i] >= position {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}

func main() {
	nums := []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}
	jump(nums)
}
