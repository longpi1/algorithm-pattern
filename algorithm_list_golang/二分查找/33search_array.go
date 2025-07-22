package main

/*
33.搜索旋转排序数组

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

### 解题思路：二分查找low_bound 【时间复杂度O(lgn),n是数组长度】

- 核心要素
    - 注意区间开闭，三种都可以
    - 循环结束条件：当前区间内没有元素
    - 下一次二分查找区间：不能再查找(区间不包含)mid，防止死循环
    - 返回值：大于等于target的第一个下标（注意循环不变量）

- 有序数组中二分查找的四种类型（下面的转换仅适用于数组中都是整数）
    1. 第一个大于等于x的下标： low_bound(x)
    2. 第一个大于x的下标：可以转换为`第一个大于等于 x+1 的下标` ，low_bound(x+1)
    3. 最后一个一个小于x的下标：可以转换为`第一个大于等于 x 的下标` 的`左边位置`, low_bound(x) - 1;
    4. 最后一个小于等于x的下标：可以转换为`第一个大于等于 x+1 的下标` 的 `左边位置`, low_bound(x+1) - 1;
*/

// 初始思路:
func search(nums []int, target int) int {
	n := len(nums)
	if target > nums[n-1] {
		for i := 0; i < n; i++ {
			if target == nums[i] {
				return i
			}
			if i+1 < n && nums[i] > nums[i+1] {
				return -1
			}
		}
	} else if target < nums[n-1] {
		for i := n - 1; i >= 0; i-- {

			if target == nums[i] {
				return i
			}
			if n > 2 && i > 0 && nums[i] < nums[i-1] {
				return -1
			}
		}
	} else {
		return n - 1
	}

	return -1
}

// 二分
/*
在每一步中，首先判断哪一部分是网格的（根据旋转的性质），然后检查目标值是否在网格的部分中。
根据比较结果，调整左右指针的位置，从而实现在旋转排序阵列中的快速搜索。
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// 判断哪一部分是有序的
		if nums[left] <= nums[mid] { // 左半部分有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右半部分有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1 // 没找到目标值
}

// 其他思路：两次二分
// 两次二分
func findMin(nums []int) int {
	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
	for left+1 < right {           // 开区间不为空
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

// search 函数用于在旋转排序数组中搜索目标值的索引
// nums 是旋转排序数组，target 是要搜索的目标值
// 返回目标值在数组中的索引，如果未找到则返回 -1
func search(nums []int, target int) int {
	// 初始化左指针，指向数组的起始位置
	left, right := 0, len(nums)-1
	// 当左指针小于等于右指针时，继续循环搜索
	for left <= right {
		// 计算中间位置，使用 left + (right - left) / 2 避免整数溢出
		mid := left + (right-left)/2
		// 如果中间位置的元素等于目标值，直接返回中间位置的索引
		if nums[mid] == target {
			return mid
		}
		// 判断左半部分是否有序
		if nums[mid] >= nums[left] {
			// 如果目标值在左半部分的有序区间内
			if nums[mid] > target && target >= nums[left] {
				// 缩小右边界，继续在左半部分搜索
				right = mid - 1
			} else {
				// 目标值不在左半部分的有序区间内，缩小左边界，在右半部分搜索
				left = mid + 1
			}
		} else {
			// 左半部分无序，说明右半部分有序
			// 如果目标值在右半部分的有序区间内
			if nums[mid] < target && target <= nums[right] {
				// 缩小左边界，继续在右半部分搜索
				left = mid + 1
			} else {
				// 目标值不在右半部分的有序区间内，缩小右边界，在左半部分搜索
				right = mid - 1
			}
		}
	}
	// 循环结束后仍未找到目标值，返回 -1
	return -1
}

func main() {
	nums := []int{1, 3}
	target := 1
	print(search(nums, target))
}
