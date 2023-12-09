package main

import "fmt"

/*
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为 target 的不同组合数少于 150 个。



示例 1：

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。
示例 2：

输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]
示例 3：

输入: candidates = [2], target = 1
输出: []


提示：

1 <= candidates.length <= 30
2 <= candidates[i] <= 40
candidates 的所有元素 互不相同
1 <= target <= 40

解决思路：
回溯算法：
# 回溯算法

回溯 <---->递归1.递归的下面就是回溯的过程

2.回溯法是一个 纯暴力的 搜索

3.回溯法解决的问题：

	3.1组合 如：1234  两两组合

	3.2切割问题 如：一个字符串有多少个切割方式 ，或者切割出来是回文

	3.3子集 ： 1 2 3 4  的子集

	3.4排列问题（顺序）

	3.5棋盘问题：n皇后  解数独

4.回溯可抽象成树形结构

5.模板
```go
result = []
func backtrack(选择列表,路径):
    if 满足结束条件:
        result.add(路径)
        return
    for 选择 in 选择列表:
        做选择
        backtrack(选择列表,路径)
        撤销选择
```
*/

/*func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	n := len(candidates)
	dfs := func(int) {}
	var result []int
	temp := target
	dfs = func(temp int) {
		if temp < 0 {
			return
		}
		if temp ==0 {
			tmp := make([]int, len(result))
			copy(tmp ,result)
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < n ; i++ {
			if candidates[i] == 0 {
				continue
			}
			temp = temp - candidates[i]

			result = append(result, candidates[i])
			dfs(temp)
			temp = temp + candidates[i]
			result = result[:n-1]
		}
	}
	dfs(temp)

	return ans
}*/

/*
上述思路此外：原因如下：
1.要避免重复寻找，for循环i应该通过变量传递，也就是startIndex,for i := 也就是startIndex; i < n ; i++ {错误；
2.result = result[:n-1],应该为result = result[:len(result)-1] ，不然数组长度小于n时并不会成功回溯；

树形结构应该如下：
https://img-blog.csdnimg.cn/20201123202227835.png
相关讲解视频：
https://www.bilibili.com/video/BV1KT4y1M7HJ/?vd_source=34718180774b041b23050c8689cdbaf2
*/

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	n := len(candidates)

	// 定义深度优先搜索函数，用于生成组合
	var dfs func(temp int, startIndex int, result []int)

	dfs = func(temp int, startIndex int, result []int) {
		if temp < 0 {
			return
		}
		if temp == 0 {
			// 如果目标值为0，表示找到了一个组合，将结果复制到临时切片并添加到答案中
			tmp := make([]int, len(result))
			copy(tmp, result)
			ans = append(ans, tmp)
			return
		}
		for i := startIndex; i < n; i++ {
			if candidates[i] == 0 {
				continue // 跳过候选列表中的0
			}
			temp = temp - candidates[i] // 尝试将当前候选值添加到组合中
			result = append(result, candidates[i]) // 记录当前候选值
			// 递归搜索下一个组合
			dfs(temp, i, result)
			temp = temp + candidates[i] // 回溯，将当前候选值从组合中移除
			result = result[:len(result)-1] // 回溯，从结果中移除最后一个元素
		}
	}

	// 初始化深度优先搜索
	dfs(target, 0, []int{})

	return ans
}




func main()  {
	candidates := []int{2,3,6,7}
	target := 7
	fmt.Printf("result:%v",combinationSum(candidates,target))
}
