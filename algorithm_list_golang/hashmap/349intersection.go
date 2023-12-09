package main

import (
	"fmt"
	"sort"
)

/*
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的


提示：

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000
*/

/*func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	m := make(map[int]bool,0)
	for i :=0; i<len(nums1); i++{
		m[nums1[i]] = true
	}

	for i := 0; i <len(nums2) ; i++ {
		if _ , ok := m[nums2[i]]; ok{
			result = append(result, nums2[i])
		}
	}
	//去重结果
	sort.Ints(result)
	for i := 1; i < len(result); i++ {
		if result[i] == result[i-1] {
			result = append(result[:i-1],result[i:]...)
		}
	}

	return result
}*/
/*
上述代码存在一个错误，数组去重方式从头遍历会导致索引错乱问题： append(result[:i-1],result[i:]...) 应该从最后的元素开始遍历

*/
func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	m := make(map[int]bool,0)
	for i :=0; i<len(nums1); i++{
		m[nums1[i]] = true
	}

	for i := 0; i <len(nums2) ; i++ {
		if _ , ok := m[nums2[i]]; ok{
			result = append(result, nums2[i])
		}
	}
	//去重结果
	sort.Ints(result)
	//这里应该从最后开始遍历，避免数组混乱
	//for i := 1; i < len(result); i++ {
	for i :=  len(result) -1; i > 0; i-- {
		if result[i] == result[i-1] {
			result = append(result[:i-1],result[i:]...)
		}
	}

	return result
}

//优化版
func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	m := make(map[int]bool, len(nums1))

	// 将 nums1 中的元素添加到 map 中
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = true
	}

	// 遍历 nums2，如果在 map 中找到元素，则添加到结果中，并从 map 中删除
	for i := 0; i < len(nums2); i++ {
		if m[nums2[i]] {
			result = append(result, nums2[i])
			delete(m, nums2[i]) // 删除 map 中的元素，避免重复添加
		}
	}

	return result
}


// 排序+双指针
func intersection(nums1 []int, nums2 []int) (res []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if res == nil || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return
}



func main()  {
	num1 := []int{4,9,5}
	num2 := []int{9,4,9,8,4}
	fmt.Printf("result:%v",intersection(num1,num2))
}