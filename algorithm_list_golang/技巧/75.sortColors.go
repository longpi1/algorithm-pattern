package main

/*
75. 颜色分类
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。


示例 1：
输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
示例 2：
输入：nums = [2,0,1]
输出：[0,1,2]
*/

func swapColors(colors []int, target int) (countTarget int) {
	for i, c := range colors {
		if c == target {
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return
}

func sortColors(nums []int) {
	count0 := swapColors(nums, 0) // 把 0 排到前面
	swapColors(nums[count0:], 1)  // nums[:count0] 全部是 0 了，对剩下的 nums[count0:] 把 1 排到前面
}

func sortColors(nums []int) {
	p0, p1 := 0, 0

	// 遍历数组
	for i, c := range nums {
		if c == 0 {
			// 如果当前元素为0，将其交换到p0位置
			nums[i], nums[p0] = nums[p0], nums[i]

			// 如果p0小于p1，说明已经有1存在，将当前元素与p1交换
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}

			// 移动p0和p1指针
			p0++
			p1++
		} else if c == 1 {
			// 如果当前元素为1，将其交换到p1位置
			nums[i], nums[p1] = nums[p1], nums[i]

			// 移动p1指针
			p1++
		}
	}
}

func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i, x := range nums {
		nums[i] = 2
		if x <= 1 {
			nums[p1] = 1
			p1++
		}
		if x == 0 {
			nums[p0] = 0
			p0++
		}
	}
}
