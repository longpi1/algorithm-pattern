package main

/*
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。



示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1


提示：

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
nums 中的每个值都 独一无二
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-104 <= target <= 104
*/

// 初始思路:
func search(nums []int, target int) int {
	n :=len(nums)
	if target > nums[n-1] {
		for i := 0; i < n ; i++ {
			if target == nums[i]{
				return i
			}
			if i+1 < n && nums[i] > nums[i+1]{
				return -1
			}
		}
	}else if  target < nums[n-1]{
		for i := n-1; i >= 0; i-- {

			if target == nums[i]{
				return i
			}
			if n >2 && i >0 && nums[i] < nums[i-1]{
				return -1
			}
		}
	}else {
		return n-1
	}

	return -1
}

// 其他思路：两次二分
// 两次二分
func findMin(nums []int) int {
	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
	for left+1 < right { // 开区间不为空
		mid := left + (right-left)/2
		if nums[mid] < nums[len(nums)-1] { // 蓝色
			right = mid
		} else { // 红色
			left = mid
		}
	}
	return right
}

func lowerBound(nums []int, left, right, target int) int {
	r0 := right
	for left+1 < right { // 开区间不为空
		// 循环不变量：
		// nums[left] < target
		// nums[right] >= target
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid // 范围缩小到 (mid, right)
		} else {
			right = mid // 范围缩小到 (left, mid)
		}
	}
	if right == r0 || nums[right] != target {
		return -1
	}
	return right
}

func searchV2(nums []int, target int) int {
	// 第一次二分 首先找到最小值
	i := findMin(nums)
	// 判断在哪一层，然后再进行二分查找
	if target > nums[len(nums)-1] {
		return lowerBound(nums, -1, i, target) // 左段
	}
	return lowerBound(nums, i-1, len(nums), target) // 右段
}




func main()  {
	nums := []int{1,3}
	target := 1
	print(search(nums,target))
}