package main

/*
283. 移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。


示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:
输入: nums = [0]
输出: [0]

*/
/*
func moveZeroes(nums []int)  {
	n := len(nums)
	left := 0
	right := n -1
	for left < right {
		if nums[left] == 0 {
			nums = append(nums, nums[left])
			nums1 := nums[0:left]
			nums2 := nums[left+1:]
			nums = append(nums1, nums2...)
			right --
		}else{
			left ++
		}
	}
}*/

/*
上述代码存在以下问题：
对切片进行修改：在循环中，您尝试对切片 nums 进行修改，但这样的修改可能会导致问题，因为 nums 的底层数组可能会因为重新分配而发生变化。这会导致不可预测的结果。
复杂的切片操作：您的代码尝试通过创建新的切片来移除零元素，这样的操作可能会导致性能问题，因为它需要创建新的切片并复制数据。
不必要的右指针：您的循环中使用了一个右指针 right，但实际上不需要它，因为您的目标是将所有的零元素移动到切片的末尾，而不需要明确的右指针。
*/

func moveZeroes(nums []int) {
	n := len(nums)
	left := 0 // 左指针指向当前非零元素的位置
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[left], nums[i] = nums[i], nums[left] // 交换非零元素到左边
			left++
		}
	}
	//// 在末尾填充零元素
	//for left < n {
	//	nums[left] = 0
	//	left++
	//}
}

/*
双指针法（快慢指针法）： 通过一个快指针和慢指针在一个for循环下完成两个for循环的工作。

定义快慢指针：

快指针：寻找新数组的元素 ，新数组就是不含有目标元素的数组
慢指针：指向更新 新数组下标的位置
*/
func moveZeroes(nums []int) {
	slowIndex := 0 // 慢指针，用于记录当前非零元素应该存放的位置

	// 遍历整个数组
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != 0 { // 如果当前元素不是零
			// 将当前非零元素与慢指针指向的位置的元素交换位置，
			// 这样可以确保非零元素被依次移到数组的前面
			nums[slowIndex], nums[fastIndex] = nums[fastIndex], nums[slowIndex]
			slowIndex++ // 慢指针向前移动
		}
	}
}



func main(){
	nums := []int{0,1,0,1}
	moveZeroes(nums)
}
