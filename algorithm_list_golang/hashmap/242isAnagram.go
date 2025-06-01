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
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int, len(s))
	arr1 := []byte(s)
	arr2 := []byte(t)
	for _, str := range arr1 {
		m[str]++
	}
	for _, str := range arr2 {
		m[str]--
	}
	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	// 优化点，字符串切割换为byte(s)或者转换为rune更合适
	strings1 := strings.Split(s, "")
	strings2 := strings.Split(t, "")
	l1 := len(strings1)
	if l1 != len(strings2) {
		return false
	}
	m := make(map[string]int, len(strings1))
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

// 其他思路：排序
func isAnagram(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

// 对于进阶问题，Unicode 是为了解决传统字符编码的局限性而产生的方案，它为每个语言中的字符规定了一个唯一的二进制编码。
//而 Unicode 中可能存在一个字符对应多个字节的问题，为了让计算机知道多少字节表示一个字符，面向传输的编码方式的 UTF−8 和 UTF−16 也随之诞生逐渐广泛使用，
//具体相关的知识读者可以继续查阅相关资料拓展视野，这里不再展开。

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}
