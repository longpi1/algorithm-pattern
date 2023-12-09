package sort

// quickSort 函数接受一个整数切片 arr 作为参数，用于对切片进行快速排序。
func quickSort(arr []int) []int {
	// 如果切片长度小于等于 1，表示已经排序好了，直接返回原切片。
	if len(arr) <= 1 {
		return arr
	}

	// 选择第一个元素作为基准值（pivot）。
	pivot := arr[0]
	// 创建两个空切片，用于存放比基准值小的元素和比基准值大的元素。
	var left, right []int

	// 遍历切片中除基准值外的元素。
	for _, v := range arr[1:] {
		// 如果元素小于等于基准值，将其加入 left 切片。
		if v <= pivot {
			left = append(left, v)
		} else {
			// 否则，将其加入 right 切片。
			right = append(right, v)
		}
	}

	// 递归对 left 和 right 切片进行快速排序。
	left = quickSort(left)
	right = quickSort(right)

	// 将 left 切片、基准值和 right 切片合并成一个有序切片，并返回。
	return append(append(left, pivot), right...)
}

