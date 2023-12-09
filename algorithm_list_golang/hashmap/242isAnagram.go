package main

import (
	"sort"
	"strings"
)

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。



示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false


提示:

1 <= s.length, t.length <= 5 * 104
s 和 t 仅包含小写字母


进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
*/

func isAnagram(s string, t string) bool {
	strings1 := strings.Split(s,"")
	strings2 := strings.Split(t,"")\
	l1 := len(strings1)
	if l1 != len(strings2) {
		return false
	}
	m := make(map[string]int,len(strings1))
	for i := 0; i < l1; i++ {
		m[strings1[i]] += 1
	}
	for i := 0; i < l1; i++ {
		m[strings2[i]] -= 1
	}
	for _, value := range m {
		if value != 0 {
			return false
		}
	}
	return true
}


//其他思路：排序
func isAnagram(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

