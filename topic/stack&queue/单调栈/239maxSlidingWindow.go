package main

import (
	"fmt"
)

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。



示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：

输入：nums = [1], k = 1
输出：[1]


提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104
1 <= k <= nums.length
*/

/*func maxSlidingWindow(nums []int, k int) []int {
	max, second_max := math.MinInt64,math.MinInt64
	n := len(nums)
	var result []int
	if n ==  0 {
		return []int{}
	}
	//先求得最开始的最大值与，第二大值

	for i := 0; i < k; i++ {
		if nums[i] > max {
			second_max = max
			max = nums[i]
		}else if nums[i] > second_max {
			second_max = nums[i]
		}
	}
	result = append(result, max)
	for i := 1; i < n -k +1  ; i++ {
		if nums[i-1] == max {
			if nums[i+k-1] > second_max {
				max = nums[i+k-1]
				second_max = math.MinInt64
				for j := i; j < i+k; j++ {
					if nums[j] > second_max{
						second_max = nums[j]
					}
				}
				result = append(result, nums[i+k-1])
				continue
			}else {
				max = second_max
				second_max = math.MinInt64
				for j := i; j < i+k; j++ {
					if nums[j] != max && nums[j] > second_max{
						second_max = nums[j]
					}
				}
				result = append(result, max)
				continue
			}
		}else if nums[i+k-1] > max {
			second_max = max
			max = nums[i+k-1]
			result = append(result, nums[i+k-1])
			continue
		}else if nums[i+k-1] > second_max {
			second_max = nums[i+k-1]
		}

		result = append(result, max)
		continue
	}
	return result
}*/
/*
上述代码关于复杂并且存在错误
*/

/*
解题思路：
思路与算法
单调递减队列
由于我们需要求出的是滑动窗口的最大值，如果当前的滑动窗口中有两个下标 i 和 j，其中 i 在 j 的左侧（i<j），并且 i 对应的元素不大于 j 对应的元素（nums[i]≤nums[j]]），那么会发生什么呢？
当滑动窗口向右移动时，只要 iii 还在窗口中，那么 jjj 一定也还在窗口中，这是 iii 在 jjj 的左侧所保证的。因此，由于 nums[j] 的存在，nums[i] 一定不会是滑动窗口中的最大值了，我们可以将 nums[i] 永久地移除。
因此我们可以使用一个队列存储所有还没有被移除的下标。在队列中，这些下标按照从小到大的顺序被存储，并且它们在数组 nums 中对应的值是严格单调递减的。因为如果队列中有两个相邻的下标，它们对应的值相等或者递增，那么令前者为 i，后者为 j，
就对应了上面所说的情况，即 nums[i] 会被移除，这就产生了矛盾。当滑动窗口向右移动时，我们需要把一个新的元素放入队列中。为了保持队列的性质，
我们会不断地将新的元素与队尾的元素相比较，如果前者大于等于后者，那么队尾的元素就可以被永久地移除，我们将其弹出队列。我们需要不断地进行此项操作，直到队列为空或者新的元素小于队尾的元素。
由于队列中下标对应的元素是严格单调递减的，因此此时队首下标对应的元素就是滑动窗口中的最大值。但与方法一中相同的是，此时的最大值可能在滑动窗口左边界的左侧，并且随着窗口向右移动，它永远不可能出现在滑动窗口中了。因此我们还需要不断从队首弹出元素，直到队首元素在窗口中为止。
为了可以同时弹出队首和队尾的元素，我们需要使用双端队列。满足这种单调性的双端队列一般称作「单调队列」。
*/

// 单调递减队列
// push 函数用于维护一个单调递减队列，新元素 val 会将队列中小于 val 的元素弹出
func push(queue []int, val int) []int {
	for len(queue) != 0 && val > queue[len(queue)-1] {
		queue = queue[:len(queue)-1]
	}
	queue = append(queue, val)
	// 返回更新后的队列
	return queue
}

// maxSlidingWindow 函数用于计算滑动窗口中的最大值
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	result := make([]int, 0)
	// 定义一个单调递减的队列
	queue := make([]int, 0)

	// 初始化滑动窗口中的第一个窗口
	for i := 0; i < k; i++ {
		queue = push(queue, nums[i])
	}
	// 将第一个窗口的最大值加入结果数组
	result = append(result, queue[0])

	// 遍历滑动窗口
	for i := 1; i < n-k+1; i++ {
		// 移除窗口的第一个元素，如果它等于队列的最大值
		if nums[i-1] == queue[0] {
			queue = queue[1:]
		}
		// 将新的元素加入窗口并维护单调递减队列
		queue = push(queue, nums[i+k-1])
		// 将队列中的最大值加入结果数组
		result = append(result, queue[0])
	}
	return result
}
func main(){
	//nums := []int{1,3,-1,-3,5,3,6,7}
	//k := 3
	nums := []int{9,10,9,-7,-4,-8,2,-6}
	k := 5
	fmt.Printf("result: %v",maxSlidingWindow(nums,k))
}
