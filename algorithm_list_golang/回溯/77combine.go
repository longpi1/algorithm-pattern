package main

import "fmt"

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。



示例 1：

输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
示例 2：

输入：n = 1, k = 1
输出：[[1]]


提示：

1 <= n <= 20
1 <= k <= n
*/



func combine(n int, k int) [][]int {
	var result [][]int
	var path []int
	used := make(map[int]bool)

	dfs := func(int) {}
	dfs = func(index int) {
		if len(path) == k {
			tmp := make([]int,k)
			copy(tmp,path)
			result =append(result, tmp)
			return
		}
		for i := index; i< n+1; i ++ {
			if !used[i] {
				used[i] = true
				path = append(path, i)
				dfs(i+1)
				used[i] = false
				path = path[:len(path)-1]
			}

		}
	}
	dfs(1)
	return result

}


func main(){
	n := 4
	k := 2
	fmt.Printf("result: %v", combine(n,k))
}