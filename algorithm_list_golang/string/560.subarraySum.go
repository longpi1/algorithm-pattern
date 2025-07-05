package main

/*
560. 和为 K 的子数组

给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。

子数组是数组中元素的连续非空序列。

示例 1：
输入：nums = [1,1,1], k = 2
输出：2
示例 2：
输入：nums = [1,2,3], k = 3
输出：2

提示：
1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/

/*
逻辑：

外层循环 for left < right 固定了子数组的起始点 left。
内层循环 for i := left; i < right; i++ 遍历了从 left 开始的所有可能的结束点 i。
对于每一个由 (left, i) 确定的子数组，代码计算其和 tmp，并与 k 比较。
时间复杂度：O(N²)，其中 N 是数组 nums 的长度。因为存在两层嵌套循环，在最坏情况下，内层循环和外层循环都接近 N 次。对于大数据量的输入，这会非常慢。

空间复杂度：O(1)，因为只使用了几个额外的变量。
*/
func subarraySum(nums []int, k int) int {
	result := 0
	left := 0
	right := len(nums)
	for left < right {
		tmp := 0
		for i := left; i < right; i++ {
			tmp += nums[i]
			if tmp == k {
				result++
			}
		}
		left++
	}
	return result
}

/*
优化代码：
优化思路：前缀和 + 哈希表
这是一个解决此类子数组和问题的经典高效方法。

核心思想：

定义 preSum[i] 为 nums[0...i] 的和。
那么，任意一个子数组 nums[j...i] 的和就可以表示为 preSum[i] - preSum[j-1]。
我们要求解的是 sum(nums[j...i]) == k，这等价于 preSum[i] - preSum[j-1] == k。
将这个等式变形，得到 preSum[j-1] == preSum[i] - k。
算法流程：

创建一个哈希表（在 Go 中是 map），用来存储 前缀和 -> 该前缀和出现的次数。
初始化一个变量 preSum（当前的前缀和）为 0，结果 count 为 0。
关键一步：在哈希表中存入 mp[0] = 1。这代表前缀和为 0 的情况出现了 1 次（即一个元素都不选的时候）。
这一步是为了正确处理那些从索引 0 开始的、和恰好为 k 的子数组。例如，nums = [3, 4], k = 3。当遍历到 nums[0] 时，
preSum 为 3，我们需要找 preSum - k = 3 - 3 = 0 的历史前缀和，mp[0] 的存在就让这种情况被正确计数。
遍历数组 nums 中的每一个元素 num：
a. 更新当前前缀和：preSum += num。
b. 在哈希表中查找 preSum - k。
如果存在，说明我们找到了若干个满足条件的子数组的起点。将这些子数组的数量（即 mp[preSum - k] 的值）累加到 count 上。
c. 将当前的前缀和 preSum 存入哈希表（或更新其出现次数）：mp[preSum]++。
遍历结束后，返回 count。
*/
func subarraySum(nums []int, k int) int {
	// 结果计数器
	count := 0
	// 当前的前缀和
	preSum := 0
	// 哈希表，用于存储 {前缀和 -> 出现次数}
	// key: 前缀和
	// value: 该前缀和出现的次数
	mp := make(map[int]int)

	// 初始化：前缀和为0的情况出现1次（空数组）
	// 这是为了处理从索引0开始的子数组，其和恰好为k的情况。
	// 例如 nums=[1,2,3], k=3。当遍历到索引1时, preSum=3。
	// 我们需要找 preSum-k = 3-3=0 的历史前缀和。
	mp[0] = 1

	// 遍历数组
	for _, num := range nums {
		// 1. 更新当前的前缀和
		preSum += num

		// 2. 查找是否存在 preSum - k 的历史前缀和
		// 如果 mp[preSum - k] 存在，说明从某个位置到当前位置的子数组和为 k
		if times, ok := mp[preSum-k]; ok {
			count += times
		}

		// 3. 将当前的前缀和存入哈希表，并更新其出现次数
		mp[preSum]++
	}

	return count
}

func subarraySum(nums []int, k int) int {
	count := 0 // 用于记录符合条件的子数组数量

	for start := 0; start < len(nums); start++ { // 外层循环遍历所有可能的子数组起始位置
		sum := 0 // 用于记录当前子数组的累计和

		// 内层循环逆向遍历子数组
		for end := start; end >= 0; end-- { // 从起始位置往前遍历子数组
			sum += nums[end] // 累计当前子数组的元素和

			if sum == k { // 如果当前子数组的元素和等于目标值 k
				count++ // 增加符合条件的子数组数量
			}
		}
	}

	return count // 返回符合条件的子数组数量
}

func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}
