package main
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
	findanagrams(s,p)
}
