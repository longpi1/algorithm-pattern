package main

import "slices"

/*
189. 轮转数组

给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例 2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]
*/
func rotate(nums []int, k int) {
	// 创建一个新的切片用于存储旋转后的元素
	newNums := make([]int, len(nums))
	k %= len(nums) // 轮转 k 次等于轮转 k % n 次
	n := len(nums)

	n1 := nums[n-k : n]
	n2 := nums[:n-k]
	newNums = append(n1, n2...)
	copy(nums, newNums)
}

/*
上述思路错误
*/

// 方法一：使用额外的数组
// 我们可以使用额外的数组来将每个元素放至正确的位置。用 nnn 表示数组的长度，我们遍历原数组，将原数组下标为 iii 的元素放至新数组下标为 (i+k) mod n(i+k)\bmod n(i+k)modn 的位置，最后将新数组拷贝至原数组即可。
func rotate(nums []int, k int) {
	// 创建一个新的切片用于存储旋转后的元素
	newNums := make([]int, len(nums))

	// 遍历原始切片 nums
	for i, v := range nums {
		// 计算新的位置，使用取模操作确保循环移动
		newIdx := (i + k) % len(nums)

		// 将原始切片 nums 中的元素放入新位置 newIdx
		newNums[newIdx] = v
	}

	// 将新的切片内容复制回原始切片 nums
	copy(nums, newNums)
}

/*
方法三：数组翻转
该方法基于如下的事实：当我们将数组的元素向右移动 kkk 次后，尾部 k mod nk\bmod nkmodn 个元素会移动至数组头部，其余元素向后移动 k mod nk\bmod nkmodn 个位置。

该方法为数组的翻转：我们可以先将所有元素翻转，这样尾部的 k mod nk\bmod nkmodn 个元素就被移至数组头部，然后我们再翻转 [0,k mod n−1][0, k\bmod n-1][0,kmodn−1] 区间的元素和 [k mod n,n−1][k\bmod n, n-1][kmodn,n−1] 区间的元素即能得到最后的答案。

我们以 n=7n=7n=7，k=3k=3k=3 为例进行如下展示：
*/

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
func rotate(nums []int, k int) {
	k %= len(nums) // 轮转 k 次等于轮转 k % n 次
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}
