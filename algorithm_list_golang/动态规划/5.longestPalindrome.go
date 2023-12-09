package main

/*
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：
输入：s = "cbbd"
输出："bb"
*/

// 动态规划
/*
这段代码使用了动态规划来查找最长回文子串。它创建了一个二维动态规划数组 dp，其中 dp[i][j] 表示从索引 i 到 j 的子串是否是回文子串。然后，通过遍历字符串 s 和填充 dp 数组，找到最长的回文子串的起始索引和长度，
最后返回最长回文子串。这个方法的时间复杂度为 O(n^2)，适用于较长的字符串。
*/
func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s // 如果字符串为空或只有一个字符，直接返回它本身
	}

	// 创建一个二维动态规划数组，dp[i][j]表示从索引i到j的子串是否是回文子串
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	start, maxLength := 0, 1 // 初始化最长回文子串的起始索引和长度

	// 初始化长度为1的子串都是回文子串
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// 遍历字符串，检查长度为2和更长的子串是否是回文子串
	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1 // 子串的结束索引
			// 如果子串的两端字符相同，并且中间的子串也是回文子串
			// 在 length == 2 的条件下，我们直接检查 s[i] == s[j] 即可，因为长度为2的子串的中间部分为空，不需要进一步检查
			if s[i] == s[j] && (length == 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				// 更新最长回文子串的起始索引和长度
				if length > maxLength {
					maxLength = length
					start = i
				}
			}
		}
	}

	// 通过起始索引和最大长度截取最长回文子串并返回
	return s[start : start+maxLength]
}


// 暴力破解
func longestPalindrome(s string) string {
	// 如果字符串为空或长度为1，直接返回该字符串作为回文子串
	if len(s) <= 1 {
		return s
	}

	// 初始化最长回文子串的起始和结束位置
	start, end := 0, 0

	// 遍历字符串中的每个字符
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			// 检查从i到j的子串是否是回文子串
			if isPalindrome(s, i, j) {
				// 如果是回文子串，并且长度大于已知最长回文子串的长度
				if j-i > end-start {
					// 更新最长回文子串的起始和结束位置
					start = i
					end = j
				}
			}
		}
	}

	// 通过起始和结束位置截取最长回文子串并返回
	return s[start : end+1]
}

// 辅助函数，检查字符串s的子串是否是回文子串
func isPalindrome(s string, left, right int) bool {
	for left < right {
		// 如果左右字符不相等，则不是回文子串
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	// 如果左右字符相等，是回文子串
	return true
}


/*
解题思路：
使用了中心扩展法来查找最长回文子串。
它首先遍历字符串中的每个字符，以该字符为中心向两边扩展
，同时计算奇数长度和偶数长度的回文子串。然后，它更新最长回文子串的起始和结束位置，并最终返回该回文子串。
*/
func longestPalindrome(s string) string {
	// 如果字符串为空或长度为1，直接返回该字符串作为回文子串
	if len(s) <= 1 {
		return s
	}

	// 初始化最长回文子串的起始和结束位置
	start, end := 0, 0

	// 遍历字符串中的每个字符
	for i := 0; i < len(s); i++ {
		// 计算以当前字符为中心的回文子串长度（奇数长度）
		len1 := expandAroundCenter(s, i, i)
		// 计算以当前字符和下一个字符之间的空隙为中心的回文子串长度（偶数长度）
		len2 := expandAroundCenter(s, i, i+1)

		// 取两种情况中的较长回文子串长度
		maxLen := max(len1, len2)

		// 如果当前回文子串的长度大于已知最长回文子串的长度
		if maxLen > end-start {
			// 更新最长回文子串的起始和结束位置
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	// 通过起始和结束位置截取最长回文子串并返回
	return s[start : end+1]
}

// 辅助函数，计算以left和right为中心的回文子串长度
func expandAroundCenter(s string, left, right int) int {
	// 防止越界，同时判断左右字符是否相等
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// 返回回文子串的长度
	return right - left - 1
}

