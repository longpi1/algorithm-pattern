package main


/*
3. 无重复字符的最长子串

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	max := 1
	for first := 0; first < n; first++ {
		m := make(map[byte]bool)
		m[s[first]] = true
		for second := first+1; second < n; second++ {
			if !m[s[second]] {
				m[s[second]] = true
				tmp := second - first + 1
				if tmp > max {
					max = tmp
				}
			}else{
				break
			}
		}
	}
	return max
}

// 优化方法
func lengthOfLongestSubstring(s string) (ans int) {
	left := 0             // 左指针，表示当前无重复字符子串的起始位置
	cnt := [128]int{}     // 用来记录字符出现的次数，ASCII字符集有128个字符
	for right, c := range s {
		cnt[c]++           // 更新字符c的出现次数
		for cnt[c] > 1 {   // 如果字符c的出现次数大于1，表示有重复字符
			cnt[s[left]]-- // 移动左指针，并减少左边字符的出现次数，直到没有重复字符
			left++          // 移动左指针
		}
		ans = max(ans, right-left+1) // 计算当前无重复字符子串的长度并更新最大值
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main(){
	s := "abcabcbb"
	print(lengthOfLongestSubstring(s))
}