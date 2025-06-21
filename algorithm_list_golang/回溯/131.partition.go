package main

import "fmt"

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
func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	result := make([][]string, 0)
	path := make([]string, 0)

	// dfs 函数定义，传入了字符串 s 和索引 index
	dfs := func(s string, index int) {}
	dfs = func(s string, index int) {
		// 错误点 1: 终止条件逻辑错误且 `len(path) == len(s)` 不对
		// `isRecall(path)` 应该是检查 path 中每个字符串是否都是回文，而不是检查 path 数组本身是否回文。
		// 此外，`len(path) == len(s)` 是错误的终止条件。当 `index == len(s)` 时，表示整个原始字符串 `s` 已经被成功分割。
		if isRecall(path) && len(path) == len(s) {
			tmp := make([]string, len(path))

			// 错误点 2: 结果集添加错误
			// 这里将 path 直接添加到 result，但 path 是一个共享的切片，后续还会被修改。
			// 应该添加 tmp (已经深拷贝的 path)。
			copy(tmp, path)
			result = append(result, path) // BUG: 应该 append(result, tmp) 而不是 path
			return
		}

		// 错误点 3: 循环条件 `i < len(s)-1` 导致遗漏最后一个字符
		// 循环应该遍历到 `len(s)`，才能正确处理到字符串的最后一个字符作为子串的结尾。
		for i := index; i < len(s)-1; i++ { // BUG: 循环条件应为 `i < len(s)`

			// 错误点 4: 递归时修改了传入的字符串 s
			// `s = s[i+1:]` 会修改当前 `dfs` 作用域内的 `s` 变量，但这不是期望的行为。
			// `dfs` 应该始终基于原始字符串 `s` 和当前的 `index` 来确定子串。
			// 每次递归应该基于原始的完整字符串，并通过 `index` 和 `i` 来确定截取的子串范围。
			path = append(path, s[index:i+1]) // OK: 获取子串
			fmt.Printf("path: %v", path)      // 调试输出，在最终代码中应移除

			s = s[i+1:] // BUG: 不应在这里修改字符串 `s`。`dfs` 的第一个参数 `s` 应该始终是原始的完整字符串。

			// 错误点 5: 递归调用参数错误
			// 递归调用 `dfs(s, i+1)` 中的 `s` 已经被修改，并且 `i+1` 应该是下一个子串的起始索引，
			// 但由于 `s` 被截断，`i+1` 的含义也变了。
			// 应该传递原始字符串 `s`，并让 `dfs` 的 `index` 参数负责跟踪当前处理到原始字符串的哪个位置。
			dfs(s, i+1) // BUG: 这里的 `s` 已经被截断，且 `i+1` 对应的是原始字符串中的索引。逻辑混淆。

			path = path[:len(path)-1] // OK: 回溯操作
		}
	}
	dfs(s, 0)
	return result
}

// 错误点 6: isRecall 函数逻辑错误且名称不符
// `isRecall` 的命名与功能不符。这个函数应该是检查一个字符串是否是回文串，
// 而不是检查一个 `[]string` 切片是否是回文串（这在“分割回文串”问题中通常不需要）。
// 并且，即使是检查 `[]string` 是否回文，`sArr[right]` 也应该改成 `sArr[len(sArr)-1-right]` 或 `sArr[len(sArr)-1-left]`。
func isRecall(sArr []string) bool {
	if len(sArr) == 1 {
		return true
	}
	left := 0
	right := len(sArr) // BUG: right 应该是 len(sArr)-1
	for left < right {
		// BUG: sArr[left] != sArr[right] 比较的是 `[]string` 中的元素，不是字符串内容是否回文
		// 假设 sArr 存储的是单个回文字符串，比如 {"aba"}, {"a", "b", "a"}。
		// 这个函数是检查 "aba" 是不是回文，还是检查 {"a", "b", "a"} 这个数组是回文？
		// 这个问题通常是检查子串 s[index:i+1] 是否回文。
		if sArr[left] != sArr[right] { // BUG: 索引越界或逻辑错误
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "aab"
	result := partition(s)
	fmt.Printf("result: %v", result)
}

//下面是修正后的代码

// isPalindrome 辅助函数：判断一个字符串是否是回文串
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// partition 主函数：分割回文串
func partition(s string) [][]string {
	result := make([][]string, 0) // 存储所有符合条件的分割方案
	path := make([]string, 0)     // 存储当前正在构建的分割方案

	// backtrack 是回溯函数
	// startIndex: 当前从原始字符串 s 的哪个索引位置开始尝试截取子串
	var backtrack func(startIndex int)
	backtrack = func(startIndex int) {
		// 1. 终止条件 / 找到解的条件
		// 如果 startIndex 已经到达字符串 s 的末尾，说明整个字符串已经被成功分割，
		// 此时 path 中存储的是一个有效的分割方案。
		if startIndex == len(s) {
			// 注意：需要对 path 进行深拷贝，因为它是共享的，后续还会被修改。
			tmp := make([]string, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return // 找到一个解后，当前分支结束
		}

		// 2. 遍历所有可能的选择
		// 从 startIndex 开始，遍历所有可能的结束位置 i 来截取子串 s[startIndex:i+1]
		for i := startIndex; i < len(s); i++ { // 修正: 循环条件应为 `i < len(s)`
			// 截取当前子串
			sub := s[startIndex : i+1]

			// 3. 剪枝条件
			// 只有当截取的子串 sub 是回文串时，才考虑将其添加到 path 并进行下一步递归。
			if isPalindrome(sub) { // 修正: 调用正确的 isPalindrome 函数

				// 4. 做选择 (将当前回文子串加入 path)
				path = append(path, sub)

				// 5. 递归调用，进入下一层决策
				// 下一次递归从 i+1 处开始，因为 s[startIndex:i+1] 已经处理完毕。
				backtrack(i + 1) // 修正: 递归调用时，传递正确的下一个起始索引

				// 6. 撤销选择 (回溯)
				// 将当前子串从 path 中移除，回到上一个决策点，尝试其他分割方式。
				path = path[:len(path)-1]
			}
		}
	}

	// 初始调用回溯函数，从字符串的起始位置 (索引 0) 开始
	backtrack(0)
	return result
}

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

//var (
//	path []string   // 存储回文子串的路径
//	res  [][]string // 存储所有回文分割方案的结果
//)
//
//func partition(s string) [][]string {
//	path, res = make([]string, 0), make([][]string, 0)
//	dfs(s, 0) // 深度优先搜索
//	return res
//}

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
