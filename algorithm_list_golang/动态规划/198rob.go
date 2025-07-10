package main

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

示例 1：
输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。

	偷窃到的最高金额 = 1 + 3 = 4 。

示例 2：
输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。

	偷窃到的最高金额 = 2 + 9 + 1 = 12 。
*/
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := make([]int, len(nums))
	result[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if i > 1 {
			//fmt.Println(i, " before ", result[i])
			// 错误: result[i] 初始值为0，max(0, result[i-2]+nums[i]) 逻辑错误
			// 应该比较"不偷当前房子"和"偷当前房子"两种情况
			// 正确应该是: result[i] = max(result[i-1], result[i-2]+nums[i])
			result[i] = max(result[i], result[i-2]+nums[i])
		} else {
			// 错误3: 当 i == 1 时，result[1] 还未赋值，值为0
			// 错误4: 逻辑错误，应该比较 nums[0] 和 nums[1]，选择较大值
			// 正确应该是: result[1] = max(nums[0], nums[1]) 或 max(result[0], nums[1])
			result[i] = max(result[0], result[1])
		}
		//fmt.Println(i, " ", result[i])
	}
	return result[len(nums)-1]
}

//func rob(nums []int) int {
//	result := 0
//	dp := make([]int, 0)
//	n := len(nums)
//	start := 0
//	for  start < n{
//		if nums[start+1] != nil {
//			if nums[start+1]
//		}else{
//			result += nums[start]
//		}
//	}
//}

/*

上述思路错误


*/

/*
解题思路如下：
确定dp数组（dp table）以及下标的含义
dp[i]：考虑下标i（包括i）以内的房屋，最多可以偷窃的金额为dp[i]。
确定递推公式
决定dp[i]的因素就是第i房间偷还是不偷。
如果偷第i房间，那么dp[i] = dp[i - 2] + nums[i] ，即：第i-1房一定是不考虑的，找出 下标i-2（包括i-2）以内的房屋，最多可以偷窃的金额为dp[i-2] 加上第i房间偷到的钱。
如果不偷第i房间，那么dp[i] = dp[i - 1]，即考 虑i-1房，（注意这里是考虑，并不是一定要偷i-1房，这是很多同学容易混淆的点）
然后dp[i]取最大值，即dp[i] = max(dp[i - 2] + nums[i], dp[i - 1]);
*/
//func rob(nums []int) int {
//	n := len(nums)
//	dp := make([]int, n)
//	//dp数组如何初始化
//	//从递推公式dp[i] = max(dp[i - 2] + nums[i], dp[i - 1]);可以看出，递推公式的基础就是dp[0] 和 dp[1]
//	//从dp[i]的定义上来讲，dp[0] 一定是 nums[0]，dp[1]就是nums[0]和nums[1]的最大值即：dp[1] = max(nums[0], nums[1]);
//	dp[0] = nums[0]
//	if n == 1 {
//		return nums[0]
//	}
//	dp[1] = max(nums[0], nums[1])
//
//	for i := 2; i < n; i++ {
//		//4.确定遍历顺序
//		//dp[i] 是根据dp[i - 2] 和 dp[i - 1] 推导出来的，那么一定是从前到后遍历！
//		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
//	}
//	return dp[n-1]
//}

func main() {
	nums := []int{1, 2, 3, 1}
	print(rob(nums))
}
