package main




/*
287. 寻找重复数

给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。

你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。


示例 1：
输入：nums = [1,3,4,2,2]
输出：2
示例 2：
输入：nums = [3,1,3,4,2]
输出：3
*/






/*
思路分析：

这个问题可以转化为链表中环的检测问题。可以把数组中的每个元素看作是链表中的一个节点，节点的值是指向下一个节点的索引。
由于数组中必然存在重复的元素，所以链表中必然存在环。为了找到环的起始点（即重复元素），使用快慢指针法。

第一步：设置两个指针，一个慢指针每次移动一个步长，一个快指针每次移动两个步长。
因为数组中的元素值都在[1,n]的范围内，所以这样移动不会越界。

第二步：快慢指针相遇后，将其中一个指针重新指向数组的起始位置，
然后两个指针每次都移动一个步长，当它们再次相遇时，就是环的起始点，也就是重复的元素。
这是因为在环中，两个指针相遇时，再次相遇的点必然是环的起点。

这个算法的时间复杂度是 O(n)，空间复杂度是 O(1)。
*/


func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]

	// 寻找环的起始点
	for slow != fast {
		slow = nums[slow]
		// 抽象理解为走两步
		fast = nums[nums[fast]]
	}

	// 重置一个指针，继续寻找环的起始点，即为重复元素
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
/*
相遇时，慢指针走的距离：D+S1
假设相遇时快指针已经绕环 n 次，它走的距离：D+n(S1+S2)+S1
因为快指针的速度是 2 倍，所以相同时间走的距离也是 2 倍：
D+n(S1+S2)+S1=2(D+S1)
即 (n−1)S1+nS2=D(n-1)
我们不关心绕了几次环，取 n = 1 这种特定情况，消掉 S1：
D=S2

*/