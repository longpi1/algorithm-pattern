package main

import "fmt"

/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。



示例 1：

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]
示例 2：

输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]


提示：

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 已按 非递减顺序 排序


进阶：

请你设计时间复杂度为 O(n) 的算法解决本问题
*/

/*
暴力做法：
func sortedSquares(nums []int) []int {
    ans := make([]int, len(nums))
    for i, v := range nums {
        ans[i] = v * v
    }
    sort.Ints(ans)
    return ans
}


*/

/*
o(n),通过双指针


*/
func sortedSquares(nums []int) []int {
	n := len(nums)
	start := 0
	last := n-1
	result := make([]int,n)
	for pos := n - 1; pos >= 0; pos-- {
		head :=nums[start] * nums[start]
		tail :=nums[last] * nums[last]
		if head >= tail {
			result[pos] = head
			start ++
		} else{
			result[pos] = tail
			last --
		}

		}
		return  result
}



func main(){
	nums := []int{-4,-1,0,3,10}
	fmt.Printf("result: %v", sortedSquares(nums))
}