package main

import "strings"

/*
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。



示例 1：

输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
示例 2：

输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。


提示：

1 <= haystack.length, needle.length <= 104
haystack 和 needle 仅由小写英文字符组成
*/

//用工具方法
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}


//暴力解决
/*func strStr(haystack string, needle string) int {
	str1 := []byte(haystack)
	str2 := []byte(needle)
	len1 := len(str1)
	len2 := len(str2)
	for i := 0; i < len1 -len2; i++ {
		for j := 0; j < len2; j++ {
			if str2[j] == str1[i]
		}
	}
	return -1
}*/
/*
上述代码存在以下几个问题：
1.不需要转换为byte数组
2.判断语句if str2[j] == str1[i]错误 应该为haystack[i+j] != needle[j]即可；判断是否相同
*/

//暴力解决
func strStr(haystack string, needle string) int {
	var i, j int
	for i = 0; i < len(haystack) - len(needle)+1; i++ {
		for j = 0; j <len(needle) ; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if len(needle) == j{
			return i
		}
	}

	return -1
}


// KMP

// 方法一:前缀表使用减1实现

// getNext 构造前缀表next
// params:
//		  next 前缀表数组
//		  s 模式串
func getNext(next []int, s string) {
	j := -1 // j表示 最长相等前后缀长度
	next[0] = j

	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j] // 回退前一位
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j // next[i]是i（包括i）之前的最长相等前后缀长度
	}
}
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := make([]int, len(needle))
	getNext(next, needle)
	j := -1 // 模式串的起始位置 next为-1 因此也为-1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j] // 寻找下一个匹配点
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 { // j指向了模式串的末尾
			return i - len(needle) + 1
		}
	}
	return -1
}

作者：代码随想录
链接：https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/solutions/732461/dai-ma-sui-xiang-lu-kmpsuan-fa-xiang-jie-mfbs/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
