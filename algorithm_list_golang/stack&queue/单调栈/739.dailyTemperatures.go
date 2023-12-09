package main

/*
739. 每日温度
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，
下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

示例 1:
输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
示例 2:
输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]
示例 3:
输入: temperatures = [30,60,90]
输出: [1,1,0]

*/
/*
暴力解法存在时间超时问题
*/
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				res[i] = j - i
				break
			}
		}
	}
	return res
}

/*
优化方法：单调栈
单调栈的使用： 这是一个单调递减栈，栈内保存的元素索引对应的温度值是单调递减的。

遍历温度数组： 遍历给定的温度数组。

破坏单调性时的处理： 如果栈非空且当前遍历元素 v 大于栈顶元素对应的温度值，则说明当前元素破坏了栈的单调性。此时，栈顶元素所对应的结果即为当前索引 i 减去栈顶元素的索引，表示在栈顶元素出现的位置之后，第一个比它大的温度所对应的天数。

入栈操作： 当前元素索引入栈，保持栈的单调性。
*/

func dailyTemperatures(num []int) []int {
	// 单调栈
	ans := make([]int, len(num)) // 用于保存结果的数组
	stack := []int{} // 单调栈，保存元素的索引
	for i, v := range num {
		// 栈非空且当前遍历元素v破坏了栈的单调性
		for len(stack) != 0 && v > num[stack[len(stack)-1]] {
			// 出栈
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = i - top // 计算出栈元素所对应的结果，即当前索引减去栈顶元素的索引
		}
		stack = append(stack, i) // 当前元素入栈
	}
	return ans
}

