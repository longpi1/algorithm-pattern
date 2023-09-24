package main

/*
76. 最小覆盖子串

注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。

示例 2：
输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。

示例 3:
输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。
*/

func minWindow(s string, t string) string {
	sLen := len(s)
	tLen := len(t)
	if tLen > sLen {
		return ""
	}
	var result string
	m := make(map[byte]int)
	for i := 0; i < tLen; i++ {
		m[t[i]] ++
	}



	return result
}

/*
实现思路：
双map + 滑动窗口
*/
func minWindow(s string, t string) string {
	// 使用 map 来实现滑动窗口
	sMap, tMap := map[byte]int{}, map[byte]int{}
	sLen, tLen := len(s), len(t)

	for i := 0; i < tLen; i++ {
		tMap[t[i]]++
	}

	left, right, length := 0, -1, sLen+1 // 初始化左右指针和最小窗口长度
	for i, j := 0, 0; j < sLen; j++ {
		sMap[s[j]]++

		// 判断是否满足条件
		for isCovered(sMap, tMap) {
			// 更新最小窗口
			if j-i+1 < length {
				length = j - i + 1
				left = i
				right = j
			}

			// 收缩窗口
			sMap[s[i]]--
			i++
		}
	}

	// 返回最小窗口的子字符串
	return s[left : right+1]
}

// 判断 s 是否完全覆盖 t
func isCovered(sMap, tMap map[byte]int) bool {
	for k, tv := range tMap {
		if sMap[k] < tv {
			return false
		}
	}
	return true
}
