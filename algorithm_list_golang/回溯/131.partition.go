package main

/*
131. 分割回文串
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。

示例 1：
输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：
输入：s = "a"
输出：[["a"]]

*/

/*func partition(s string) [][]string {
	n := len(s)
	if n == 0 {
		return nil
	}
	result := make([][]string,0)
	str := []byte(s)
	tmp := make([]string, 0)

	dfs := func() {
		if isPartition(string(tmp[:])) {
			result = append(result, string(str[:]))
		}
		for i := 0; i < n; i++ {

		}
	}
	dfs()
	return result
}

func isPartition(s string) bool {
	if len(s) == 1 {
		return true
	}
	n := len(s)
	for i := 0; i <= 2/n; i++ {
		if s[i] != s[n-i] {
			return false
		}
	}
	return true
}
*/

/*
上述思路错误
*/

/*
代码注解：

partition 函数： 接受一个字符串 s，用于分割回文子串。函数中初始化了 path 存储回文子串的路径，res 存储所有回文分割方案的结果。
然后调用 dfs 函数进行深度优先搜索，最终返回结果。
dfs 函数： 接受两个参数 s（输入字符串）和 start（起始位置），用于查找回文子串。在函数中，通过遍历字符串，
截取可能的回文子串，并判断是否为回文。如果是回文，则将该子串加入 path，递归调用 dfs 继续寻找下一个回文子串，
最后回溯，将当前子串从 path 中移除，继续查找其他可能的回文子串。
isPalindrome 函数： 判断输入字符串 s 是否为回文。
main 函数： 提供一个示例，演示如何使用 partition 函数。
*/

var (
	path []string   // 存储回文子串的路径
	res  [][]string // 存储所有回文分割方案的结果
)

func partition(s string) [][]string {
	path, res = make([]string, 0), make([][]string, 0)
	dfs(s, 0) // 深度优先搜索
	return res
}

//func dfs(s string, start int) {
//	if start == len(s) {  // 如果起始位置等于s的长度，说明已经找到了一组分割方案
//		tmp := make([]string, len(path))
//		copy(tmp, path)  // 将当前的回文子串路径复制到临时变量tmp中
//		res = append(res, tmp)  // 将临时变量tmp添加到结果中
//		return
//	}
//	for i := start; i < len(s); i++ {
//		str := s[start : i+1]  // 取出当前子串
//		if isPalindrome(str) {  // 如果当前子串是回文子串
//			path = append(path, str)  // 将当前子串添加到路径中
//			dfs(s, i+1)  // 递归查找下一个子串
//			path = path[:len(path)-1]  // 回溯，将当前子串从路径中移除
//		}
//	}
//}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false // 如果不是回文，返回false
		}
	}
	return true // 是回文，返回true
}
