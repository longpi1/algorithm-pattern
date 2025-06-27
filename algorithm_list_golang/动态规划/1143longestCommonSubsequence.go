package main

/*
1143. 最长公共子序列

给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

示例 1：
输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。
示例 2：
输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc" ，它的长度为 3 。
示例 3：
输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0 。
*/

// 下述为错误代码：
/*
错误原因
逻辑错误：该算法没有实现LCS的逻辑。它只是简单地统计了在两个字符串中，字符两两相等的次数。
不考虑顺序：LCS的核心在于保持字符的相对顺序。例如，对于text1="banana"和text2="atana"，LCS是"aana"（长度4）。当前代码可能会错误地计算匹配。
不考虑“最长”：LCS是一个优化问题，需要找到长度最长的那个公共子序列。当前代码只是计数。
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	s1 := []byte(text1)
	s2 := []byte(text2)
	result := 0
	// 外层循环遍历 text1 中的每个字符
	for i := 0; i < len(s1); i++ {
		// 内层循环遍历 text2 中的每个字符
		for j := 0; j < len(s2); j++ {
			// 如果 text1[i] 和 text2[j] 相等
			if s1[i] == s2[j] {
				// 错误点：仅仅因为两个字符相等就增加 result。
				// 这没有考虑字符的顺序，也没有试图构建一个“子序列”。
				// 例如：text1 = "aa", text2 = "a"
				// i=0, s1[0]='a':
				//   j=0, s2[0]='a': s1[0]==s2[0], result = 1
				// i=1, s1[1]='a':
				//   j=0, s2[0]='a': s1[1]==s2[0], result = 2
				// 返回 2，但LCS是 "a"，长度为1。

				// 再例如：text1 = "abc", text2 = "axbyc"
				// 'a' vs 'a': result = 1
				// 'b' vs 'b': result = 2
				// 'c' vs 'c': result = 3
				// 返回 3，这是正确的。

				// 但如果 text1 = "aba", text2 = "bab"
				// i=0, s1[0]='a':
				//   j=1, s2[1]='a': result = 1
				// i=1, s1[1]='b':
				//   j=0, s2[0]='b': result = 2
				//   j=2, s2[2]='b': result = 3
				// i=2, s1[2]='a':
				//   j=1, s2[1]='a': result = 4
				// 返回 4。但LCS可以是 "ab" 或 "ba"，长度都是2。
				// 这个算法计算的是 s1 中的每个字符在 s2 中出现的次数之和 (如果字符相同)。
				result++
			}
		}
	}
	return result
}

/*
func longestCommonSubsequence(text1 string, text2 string) int {
	s1 := len(text1)
	s2 := len(text2)

}
*/

/*

继续动规五部曲分析如下：
确定dp数组（dp table）以及下标的含义
dp[i][j]：长度为[0, i - 1]的字符串text1与长度为[0, j - 1]的字符串text2的最长公共子序列为dp[i][j]
有同学会问：为什么要定义长度为[0, i - 1]的字符串text1，定义为长度为[0, i]的字符串text1不香么？
这样定义是为了后面代码实现方便，如果非要定义为为长度为[0, i]的字符串text1也可以，大家可以试一试！
确定递推公式
主要就是两大情况： text1[i - 1] 与 text2[j - 1]相同，text1[i - 1] 与 text2[j - 1]不相同

如果text1[i - 1] 与 text2[j - 1]相同，那么找到了一个公共元素，所以dp[i][j] = dp[i - 1][j - 1] + 1;

如果text1[i - 1] 与 text2[j - 1]不相同，那就看看text1[0, i - 2]与text2[0, j - 1]的最长公共子序列 和 text1[0, i - 1]与text2[0, j - 2]的最长公共子序列，取最大的。

即：dp[i][j] = max(dp[i - 1][j], dp[i][j - 1]);

*/

// longestCommonSubsequence 计算两个字符串的最长公共子序列的长度
func longestCommonSubsequence(text1 string, text2 string) int {
	t1 := len(text1)
	t2 := len(text2)
	dp := make([][]int, t1+1) // 创建动态规划表
	for i := range dp {
		dp[i] = make([]int, t2+1)
	}

	for i := 1; i <= t1; i++ {
		for j := 1; j <= t2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1 // 如果字符相等，当前位置的最长公共子序列长度为前一个位置的长度加1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) // 如果字符不相等，取左方和上方的较大值作为当前位置的最长公共子序列长度
			}
		}
	}
	return dp[t1][t2] // 返回动态规划表右下角的值，即最长公共子序列的长度
}

// max 返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
