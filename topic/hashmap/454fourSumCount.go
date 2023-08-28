package main


/*

给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0


示例 1：

输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
输出：2
解释：
两个元组如下：
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
示例 2：

输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
输出：1


  提示：

n == nums1.length
n == nums2.length
n == nums3.length
n == nums4.length
1 <= n <= 200
-228 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 228
*/

/*
我们可以将四个数组分成两部分，A 和 B 为一组，C 和 D 为另外一组。

对于 A 和 B，我们使用二重循环对它们进行遍历，得到所有 A[i]+B[j] 的值并存入哈希映射中。对于哈希映射中的每个键值对，每个键表示一种 A[i]+B[j]，对应的值为 A[i]+B[j] 出现的次数。

对于 C 和 D，我们同样使用二重循环对它们进行遍历。当遍历到 C[k]+D[l] 时，如果 −(C[k]+D[l]) 出现在哈希映射中，那么将 −(C[k]+D[l])对应的值累加进答案中。

最终即可得到满足 A[i]+B[j]+C[k]+D[l]=0的四元组数目。

*/
/*
*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int, 0)
	result := 0
	// 1.先求得前两个数组之和，通过哈希表保存
	for i :=0; i<len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			m[nums1[i]+nums2[j]] += 1
		}
	}
	// 2.然后求数组3和4之和判断是否等于0；
	for i :=0; i<len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			require := 0 - nums3[i] - nums4[j]
			if  _, ok:=m[require] ; ok{
				// 注意事项 ，这里不是+1 ，而是直接等于对应的value，以为容许重复
				result += m[require]
			}
		}
	}
	return  result
}


func main(){

}