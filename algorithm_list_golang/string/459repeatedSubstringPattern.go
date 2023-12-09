package main

/*
给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。



示例 1:

输入: s = "abab"
输出: true
解释: 可由子串 "ab" 重复两次构成。
示例 2:

输入: s = "aba"
输出: false
示例 3:

输入: s = "abcabcabcabc"
输出: true
解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)


提示：

1 <= s.length <= 104
s 由小写英文字母组成
*/




/*
第一次没有做出来
思路：
https://leetcode.cn/problems/repeated-substring-pattern/solutions/386481/zhong-fu-de-zi-zi-fu-chuan-by-leetcode-solution/
*/
func repeatedSubstringPattern(s string) bool {
	n  := len(s)
	for i := 1; i <= n/2 ; i++ {
		if n % i == 0 {
			match := true
			for j := i; j < n; j++ {
				//关键步骤：这里遍历每个j和j-i的位置看是否相同
				if s[j] != s[j - i] {
					match = false
					break
				}
				return true
			}
			if match {
				return true
			}
		}
	}
	return false
}

//kmp做法
func repeatedSubstringPattern(s string) bool {
	return kmp(s + s, s)
}

func kmp(query, pattern string) bool {
	n, m := len(query), len(pattern)
	fail := make([]int, m)
	for i := 0; i < m; i++ {
		fail[i] = -1
	}
	for i := 1; i < m; i++ {
		j := fail[i - 1]
		for j != -1 && pattern[j + 1] != pattern[i] {
			j = fail[j]
		}
		if pattern[j + 1] == pattern[i] {
			fail[i] = j + 1
		}
	}
	match := -1
	for i := 1; i < n - 1; i++ {
		for match != -1 && pattern[match + 1] != query[i] {
			match = fail[match]
		}
		if pattern[match + 1] == query[i] {
			match++
			if match == m - 1 {
				return true
			}
		}
	}
	return false
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/repeated-substring-pattern/solutions/386481/zhong-fu-de-zi-zi-fu-chuan-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	repeatedSubstringPattern("aaaaaa")
}