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

// QuickSortIterative 是非递归（迭代）版本的快速排序
func QuickSortIterative(nums []int) {
	if len(nums) < 2 {
		return // 数组为空或只有一个元素，无需排序
	}

	// 1. 创建我们的“任务清单”（用切片模拟栈）
	// 栈里存放的是一个个任务，每个任务是一个 [2]int 数组，代表 [low, high]
	taskStack := make([][2]int, 0)

	// 2. 将第一个、也是最大的任务（排序整个数组）放入清单
	taskStack = append(taskStack, [2]int{0, len(nums) - 1})

	// 3. 只要清单不为空，就持续处理任务
	for len(taskStack) > 0 {
		// 3.1 取出清单顶部的任务
		n := len(taskStack) - 1   // 获取栈顶索引
		task := taskStack[n]      // 获取任务
		taskStack = taskStack[:n] // 从清单中移除该任务 (Pop)

		low, high := task[0], task[1]

		// 3.2 对当前任务范围进行分区
		// partition 函数和递归版本完全一样
		pivotIndex := partition(nums, low, high)

		// 3.3 检查是否产生了新的子任务，并加入清单
		// 如果基准点左边至少还有两个元素，说明它是一个有效的新任务
		if pivotIndex-1 > low {
			taskStack = append(taskStack, [2]int{low, pivotIndex - 1})
		}

		// 如果基准点右边至少还有两个元素，也是一个有效的新任务
		if pivotIndex+1 < high {
			taskStack = append(taskStack, [2]int{pivotIndex + 1, high})
		}
	}
}

// partition 函数和递归版本完全一样，无需任何改动
// 它的职责就是：选一个基准，把小的放左边，大的放右边，然后返回基准的最终位置
func partition(nums []int, low, high int) int {
	pivotValue := nums[high]
	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] <= pivotValue {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}
