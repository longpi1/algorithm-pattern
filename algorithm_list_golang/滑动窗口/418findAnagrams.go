package main

import "fmt"

/*
438. 找到字符串中所有字母异位词

给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。


*/

/*
下述代码只能对一部分
1. 效率极低：时间复杂度过高
这是最主要的问题。你的算法本质上是一个暴力的、带有一些小优化的嵌套循环。

外层循环: for left < len(s)-1 遍历了 s 的几乎所有位置作为起点。
内层循环: for i := left; i < len(s); i++ 从 left 开始向右扫描。
在最坏的情况下（例如 s = "aaaaaaaa", p = "aa"），外层循环和内层循环会组合成一个接近 O(N*M) 的复杂度，其中 N 是 s 的长度，M 是 p 的长度。对于这道题，这通常会导致超时 (Time Limit Exceeded)。

正确的解法——滑动窗口 (Sliding Window)——只需要 O(N) 的时间复杂度。

2. 大量的重复计算
效率低下的根源在于重复计算。当你的外层循环 left++ 后，内层循环会从 left 重新开始构建和检查 tmpMap。

思考一下：

当 left = 0 时，你检查了窗口 s[0:len(p)]。
当 left = 1 时，你又从头开始检查窗口 s[1:len(p)+1]。
这两个窗口有大量的重叠部分（s[1:len(p)-1]），而你的算法完全抛弃了上一步的计算结果，每次都重新建立 tmpMap，这是巨大的浪费。

3. 逻辑漏洞与不健壮的指针移动
你的 break 和 left 指针更新逻辑存在问题：

left = i 的问题: 当遇到一个不在 p 中的字符 s[i] 时，你试图将 left 直接跳到 i。但外层循环还有一个 left++，所以下一次的起始点会是 i+1。这看起来是一个优化，但它嵌在一个低效的框架里，效果有限。
break 的问题: 当 tmpMap[s[i]] > m[s[i]] 时，你 break 了内层循环。这意味着当前以 left 为起点的窗口无效了。然后外层循环 left++，从 left+1 重新开始。这个逻辑太慢了。比如 s = "abacaba", p = "aab", 当 left=0 时，窗口检查到 s[2]='a' 时发现有两个'a'了，而 p 中也需要两个'a'，但下一个s[3]='c'不在p中，你的逻辑会怎么处理？这会变得很复杂。
外层循环条件: for left < len(s)-1 可能导致错过以倒数第二个字符为起点的情况，如果 len(p) 恰好为1。更准确的条件应该是 for left <= len(s) - len(p)。
*/
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	result := make([]int, 0)
	m := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		m[p[i]] += 1
	}
	left := 0

	for left < len(s)-1 {
		tmpMap := make(map[byte]int)
		for i := left; i < len(s); i++ {
			_, ok := m[s[i]]
			if !ok {
				left = i
				break
			}
			tmpMap[s[i]] += 1
			if tmpMap[s[i]] > m[s[i]] {
				break
			}
			if i-left+1 == len(p) {
				result = append(result, left)
				break
			}
		}
		left++

	}
	return result
}

/*func findanagrams(s string, p string) []int {
	result := make([]int,0)
	slen := len(s)
	plen := len(p)
	if slen == 0 || plen == 0 || slen < plen {
		return result
	}
	m := make(map[byte]int)
	for i := 0; i < plen ; i++ {
		m[p[i]]++
	}
	for i := 0; i < slen; i++ {

	}
	return result
}*/
/*
实现思路
这种方法通过维护两个字符频率的哈希表，分别记录字符串 p 中字符的频率和当前窗口中字符的频率。在遍历字符串 s 时，不断调整窗口大小，并更新窗口哈希表。
当窗口的字符频率与 p 中的字符频率一致时，表示找到一个字母异位词的起始索引。这种方法避免了字符计数的重复计算，提高了效率。
*/

func findanagrams(s string, p string) []int {
	result := make([]int, 0)
	slen := len(s)
	plen := len(p)
	if slen == 0 || plen == 0 || slen < plen {
		return result
	}

	// 创建字符频率的哈希表
	pfreq := make(map[byte]int)
	for i := 0; i < plen; i++ {
		pfreq[p[i]]++
	}

	// 初始化窗口左右边界和字符频率哈希表
	left, right := 0, 0
	windowfreq := make(map[byte]int)

	for right < slen {
		// 右边界字符进入窗口，更新窗口哈希表
		windowfreq[s[right]]++

		// 如果窗口中某字符的频率超过了 p 中的频率，则需要左边界右移
		for windowfreq[s[right]] > pfreq[s[right]] {
			// 左边界字符移出窗口，更新窗口哈希表
			windowfreq[s[left]]--
			// 左边界右移
			left++
		}

		// 如果窗口的长度等于 p 的长度，表示找到了一个字母异位词
		if right-left+1 == plen {
			result = append(result, left)
		}

		// 右边界右移
		right++
	}

	return result
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	result := findAnagrams(s, p)

	fmt.Printf("result:%v", result)
}
