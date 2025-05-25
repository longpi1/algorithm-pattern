package main

import "math"

/*
4. 寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5


实现思路：
https://leetcode.cn/problems/median-of-two-sorted-arrays/solutions/2950686/tu-jie-xun-xu-jian-jin-cong-shuang-zhi-z-p2gd/?envType=problem-list-v2&envId=array
*/

func findMedianSortedArrays(a, b []int) float64 {
	if len(a) > len(b) {
		a, b = b, a // 保证下面的 i 可以从 0 开始枚举， 确保 a 是较短的数组，以减少二分查找的复杂度
	}

	m, n := len(a), len(b)
	a = append([]int{math.MinInt}, append(a, math.MaxInt)...)
	b = append([]int{math.MinInt}, append(b, math.MaxInt)...)

	// 枚举 nums1 有 i 个数在第一组
	// 那么 nums2 有 j = (m+n+1)/2 - i 个数在第一组
	i, j := 0, (m+n+1)/2
	for {
		if a[i] <= b[j+1] && a[i+1] > b[j] { // 写 >= 也可以
			max1 := max(a[i], b[j])     // 第一组的最大值
			min2 := min(a[i+1], b[j+1]) // 第二组的最小值
			if (m+n)%2 > 0 {
				return float64(max1)
			}
			return float64(max1+min2) / 2
		}
		i++ // 继续枚举
		j--
	}
}

//作者：灵茶山艾府
//链接：https://leetcode.cn/problems/median-of-two-sorted-arrays/solutions/2950686/tu-jie-xun-xu-jian-jin-cong-shuang-zhi-z-p2gd/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
//func main() {
//	fmt.Println(5 / 2)
//	fmt.Println(5 % 2)
//}
