package main

import "fmt"

/*
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。
判题标准:

系统会用下面的代码来测试你的题解:

int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案

int k = removeDuplicates(nums); // 调用

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
如果所有断言都通过，那么您的题解将被 通过。



示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

方法一：双指针
这道题目的要求是：对给定的有序数组 nums\textit{nums}nums 删除重复元素，在删除重复元素之后，每个元素只出现一次，并返回新的长度，上述操作必须通过原地修改数组的方法，使用 O(1)O(1)O(1) 的空间复杂度完成。

由于给定的数组 nums\textit{nums}nums 是有序的，因此对于任意 i<ji<ji<j，如果 nums[i]=nums[j]\textit{nums}[i]=\textit{nums}[j]nums[i]=nums[j]，则对任意 i≤k≤ji \le k \le ji≤k≤j，必有 nums[i]=nums[k]=nums[j]\textit{nums}[i]=\textit{nums}[k]=\textit{nums}[j]nums[i]=nums[k]=nums[j]，即相等的元素在数组中的下标一定是连续的。利用数组有序的特点，可以通过双指针的方法删除重复元素。

如果数组 nums\textit{nums}nums 的长度为 000，则数组不包含任何元素，因此返回 000。

当数组 nums\textit{nums}nums 的长度大于 000 时，数组中至少包含一个元素，在删除重复元素之后也至少剩下一个元素，因此 nums[0]\textit{nums}[0]nums[0] 保持原状即可，从下标 111 开始删除重复元素。

定义两个指针 fast\textit{fast}fast 和 slow\textit{slow}slow 分别为快指针和慢指针，快指针表示遍历数组到达的下标位置，慢指针表示下一个不同元素要填入的下标位置，初始时两个指针都指向下标 111。

假设数组 nums\textit{nums}nums 的长度为 nnn。将快指针 fast\textit{fast}fast 依次遍历从 111 到 n−1n-1n−1 的每个位置，对于每个位置，如果 nums[fast]≠nums[fast−1]\textit{nums}[\textit{fast}] \ne \textit{nums}[\textit{fast}-1]nums[fast]

=nums[fast−1]，说明 nums[fast]\textit{nums}[\textit{fast}]nums[fast] 和之前的元素都不同，因此将 nums[fast]\textit{nums}[\textit{fast}]nums[fast] 的值复制到 nums[slow]\textit{nums}[\textit{slow}]nums[slow]，然后将 slow\textit{slow}slow 的值加 111，即指向下一个位置。

遍历结束之后，从 nums[0]\textit{nums}[0]nums[0] 到 nums[slow−1]\textit{nums}[\textit{slow}-1]nums[slow−1] 的每个元素都不相同且包含原数组中的每个不同的元素，因此新的长度即为 slow\textit{slow}slow，返回 slow\textit{slow}slow 即可。

作者：力扣官方题解
链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array/solutions/728105/shan-chu-pai-xu-shu-zu-zhong-de-zhong-fu-tudo/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func removeDuplicates(nums []int) int {
	//用逆向遍历是因为，删除了末尾的元素，其余元素的下标位子并未改变；而若用正向遍历，若删除了前面的元素，则后面所有的元素的下标都发生了变化？
	for i := len(nums) -1; i > 0 ; i-- {
		if nums[i]==nums[i-1]{
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return  len(nums)
}


func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Printf("result: %v",removeDuplicates(nums))
}