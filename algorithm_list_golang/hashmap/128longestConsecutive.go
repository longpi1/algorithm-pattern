package main

/*
128. 最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9

*/


func longestConsecutive(nums []int) int {
	numSet := map[int]bool{} // 创建一个哈希集合来存储输入数组中的数字
	for _, num := range nums {
		numSet[num] = true // 将每个数字添加到哈希集合中，以便快速查找
	}
	longestStreak := 0 // 用于记录最长连续序列的长度

	// 遍历哈希集合中的每个数字
	for num := range numSet {
		if !numSet[num-1] { // 如果当前数字的前一个数字不在哈希集合中
			currentNum := num // 从当前数字开始
			currentStreak := 1 // 初始化当前连续序列的长度为1

			// 继续查找下一个连续的数字
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if longestStreak < currentStreak { // 更新最长连续序列的长度
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak // 返回最长连续序列的长度
}
