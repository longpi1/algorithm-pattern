package sort


// mergeSort 函数接受一个整数切片 arr 作为参数，用于对切片进行归并排序。
func mergeSort(arr []int) []int {
	// 如果切片长度小于等于 1，表示已经排序好了，直接返回原切片。
	if len(arr) <= 1 {
		return arr
	}

	// 计算中间位置，用于分割切片。
	middle := len(arr) / 2
	// 递归排序左半部分切片。
	left := mergeSort(arr[:middle])
	// 递归排序右半部分切片。
	right := mergeSort(arr[middle:])

	// 调用 merge 函数将左右两部分切片合并成一个有序切片。
	return merge(left, right)
}

// merge 函数用于将两个有序整数切片 left 和 right 合并成一个有序整数切片。
func merge(left, right []int) []int {
	// 创建一个用于存储合并结果的切片，预分配足够的容量以避免多次分配内存。
	result := make([]int, 0, len(left)+len(right))
	// 定义两个指针 i 和 j 分别指向左右两部分切片。
	i, j := 0, 0

	// 循环比较左右两部分切片的元素，将较小的元素添加到 result 中。
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 处理剩余的元素，如果有的话。
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	// 返回合并后的有序切片。
	return result
}
