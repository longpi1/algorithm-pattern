package main

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。



示例 1:

输入: nums = [1,3,5,6], target = 5
输出: 2
示例 2:

输入: nums = [1,3,5,6], target = 2
输出: 1
示例 3:

输入: nums = [1,3,5,6], target = 7
输出: 4


提示:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 为 无重复元素 的 升序 排列数组
-104 <= target <= 104

方法一：二分查找
思路与算法

假设题意是叫你在排序数组中寻找是否存在一个目标值，那么训练有素的读者肯定立马就能想到利用二分法在 O(log⁡n)O(\log n)O(logn) 的时间内找到是否存在目标值。但这题还多了个额外的条件，即如果不存在数组中的时候需要返回按顺序插入的位置，那我们还能用二分法么？答案是可以的，我们只需要稍作修改即可。

考虑这个插入的位置 pos\textit{pos}pos，它成立的条件为：

nums[pos−1]<target≤nums[pos] \textit{nums}[pos-1]<\textit{target}\le \textit{nums}[pos]
nums[pos−1]<target≤nums[pos]
其中 nums\textit{nums}nums 代表排序数组。由于如果存在这个目标值，我们返回的索引也是 pos\textit{pos}pos，因此我们可以将两个条件合并得出最后的目标：「在一个有序数组中找第一个大于等于 target\textit{target}target 的下标」。

问题转化到这里，直接套用二分法即可，即不断用二分法逼近查找第一个大于等于 target\textit{target}target 的下标 。下文给出的代码是笔者习惯的二分写法，ans\textit{ans}ans 初值设置为数组长度可以省略边界条件的判断，因为存在一种情况是 target\textit{target}target 大于数组中的所有数，此时需要插入到数组长度的位置。

作者：力扣官方题解
链接：https://leetcode.cn/problems/search-insert-position/solutions/333632/sou-suo-cha-ru-wei-zhi-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/


//### 解题思路：二分查找low_bound 【时间复杂度O(lgn),n是数组长度】
//
//- 核心要素
//    - 注意区间开闭，三种都可以
//    - 循环结束条件：当前区间内没有元素
//    - 下一次二分查找区间：不能再查找(区间不包含)mid，防止死循环
//    - 返回值：大于等于target的第一个下标（注意循环不变量）
//
//- 有序数组中二分查找的四种类型（下面的转换仅适用于数组中都是整数）
//    1. 第一个大于等于x的下标： low_bound(x)
//    2. 第一个大于x的下标：可以转换为`第一个大于等于 x+1 的下标` ，low_bound(x+1)
//    3. 最后一个一个小于x的下标：可以转换为`第一个大于等于 x 的下标` 的`左边位置`, low_bound(x) - 1;
//    4. 最后一个小于等于x的下标：可以转换为`第一个大于等于 x+1 的下标` 的 `左边位置`, low_bound(x+1) - 1;
func searchInsert(nums []int, target int) int {
	left ,right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target{
			left = mid +1
		}else{
			right = mid -1
		}
	}
	return left
}
func main()  {
	nums := []int{1,3,5,6}
	target := 2
	println(searchInsert(nums,target))
}