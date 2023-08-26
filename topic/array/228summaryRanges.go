package main

import "strconv"

/*

给定一个  无重复元素 的 有序 整数数组 nums 。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

"a->b" ，如果 a != b
"a" ，如果 a == b


示例 1：

输入：nums = [0,1,2,4,5,7]
输出：["0->2","4->5","7"]
解释：区间范围是：
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
示例 2：

输入：nums = [0,2,3,4,6,8,9]
输出：["0","2->4","6","8->9"]
解释：区间范围是：
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"


提示：

0 <= nums.length <= 20
-231 <= nums[i] <= 231 - 1
nums 中的所有值都 互不相同
nums 按升序排列
*/

/*
解题思路：双指针
注意点：需要注意int转字符串的方法strconv.Itoa
*/
/*func summaryRanges(nums []int) (ans []string) {
	f := func(i, j int) string {
		if i == j {
			return string(nums[i])
		}
		return string(nums[i]) + "->" + string(nums[j])
	}
	for i, j, n := 0, 0, len(nums); i < n; i = j + 1 {
		j = i
		for j+1 < n && nums[j+1] == nums[j]+1 {
			j++
		}
		ans = append(ans, f(i, j))
	}
	return
}*/

//
/*
上述答案错误
上述解题不应该用string做类型转换应该用strconv.Itoa
在Go语言中，使用 string() 方法将整数直接转换为字符串并不会得到预期的结果，因为 string() 方法并没有专门用于将整数转换为字符串的实现。相反，它将整数的 Unicode 码点值解释为字符，并生成一个代表该字符的字符串。
这意味着当您使用 string() 方法将整数转换为字符串时，实际上会生成一个只包含一个字符的字符串，该字符对应于该整数的 Unicode 码点值。
*/

func summaryRanges(nums []int) (ans []string) {
	f := func(i, j int) string {
		if i == j {
			return  strconv.Itoa(nums[i])
		}
		//不应该用string做类型转换应该用strconv.Itoa
		//在Go语言中，使用 string() 方法将整数直接转换为字符串并不会得到预期的结果，因为 string() 方法并没有专门用于将整数转换为字符串的实现。相反，它将整数的 Unicode 码点值解释为字符，并生成一个代表该字符的字符串。
		//这意味着当您使用 string() 方法将整数转换为字符串时，实际上会生成一个只包含一个字符的字符串，该字符对应于该整数的 Unicode 码点值。
		//return string(nums[i]) + "->" + string(nums[j])
		return strconv.Itoa(nums[i]) + "->" + strconv.Itoa(nums[j])

	}
	for i, j, n := 0, 0, len(nums); i < n; i = j + 1 {
		j = i
		for j+1 < n && nums[j+1] == nums[j]+1 {
			j++
		}
		ans = append(ans, f(i, j))
	}
	return
}