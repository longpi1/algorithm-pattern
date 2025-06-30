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
// 第一版思路错误
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	strs := []byte(s)

	// tmp 用于记录当前无重复子串的长度，maxStrLen 用于记录全局最大长度。
	tmp, maxStrLen := 0, 0

	// m 用于存储字符及其在字符串中最后一次出现的索引。
	m := make(map[byte]int)

	// 遍历整个字符串
	for i := 0; i < len(strs); i++ {
		// 检查当前字符 strs[i] 是否已经存在于 map 中
		val, ok := m[strs[i]]

		// 情况一：当前字符是新字符（在当前窗口内没有出现过）
		if !ok {
			// 当前无重复子串的长度加 1
			tmp++
			// 更新全局最大长度
			maxStrLen = max(maxStrLen, tmp)

			// 情况二：当前字符是重复字符
		} else {
			// === 问题 1：不应该直接删除 map 中的元素 ===
			// 滑动窗口算法的核心是移动窗口的左边界。
			// 简单地删除重复的字符 `strs[i]` 是错误的。
			// 考虑 "tmmzuxt"：
			// 当 i=2 (第二个'm')时，你删除了 map 中的 'm'。
			// 当 i=5 ('x')时，ok=false，tmp++。
			// 当 i=6 ('t')时，ok=true，val 是 0 (第一个't'的索引)。
			// 这时，你的 map 状态是不完整的，无法正确判断窗口的有效性。
			// 正确的做法应该是移除所有在重复字符 `strs[val]` 左边的字符，
			// 而不是只删除 `strs[i]` 这一个。
			delete(m, strs[i])

			// === 问题 2：tmp 的计算逻辑完全错误 ===
			// `tmp = i - val - 1` 这个公式没有任何意义。
			// 它试图根据新旧索引的差来计算新的窗口长度，但这是不正确的。
			// 让我们来看一个例子: s = "abcabcbb"
			// i=0, 'a': tmp=1, max=1, m={'a':0}
			// i=1, 'b': tmp=2, max=2, m={'a':0, 'b':1}
			// i=2, 'c': tmp=3, max=3, m={'a':0, 'b':1, 'c':2}
			// i=3, 'a': 遇到重复。val=0。
			//   按你的公式: tmp = 3 - 0 - 1 = 2。这碰巧对了，当前窗口是 "bca"。
			//   但我们再看一个例子: s = "pwwkew"
			//   i=0, 'p': tmp=1, max=1, m={'p':0}
			//   i=1, 'w': tmp=2, max=2, m={'p':0, 'w':1}
			//   i=2, 'w': 遇到重复。val=1。
			//     按你的公式: tmp = 2 - 1 - 1 = 0。这是错误的！正确的当前窗口 "w" 长度应该是 1。
			tmp = i - val - 1

			// (注释掉的打印语句)
			//fmt.Printf("tmp:%v", tmp)
			//fmt.Printf("val:%v", val)
			//fmt.Printf("i:%v", i)
			//fmt.Println()
		}
		// 无论如何，都更新当前字符的最新位置。
		m[strs[i]] = i
	}
	return maxStrLen
}

// 第二版思路错误 left移动存在错误
func lengthOfLongestSubstring(s string) int {
	strs := []byte(s)
	maxLength := 0

	// m 用于存储字符及其最后一次出现的索引
	m := make(map[byte]int)

	// left 是滑动窗口的左边界
	left := 0

	// i 是滑动窗口的右边界
	for i := 0; i < len(strs); i++ {
		// 检查当前字符 strs[i] 是否已经存在于 map 中
		if val, ok := m[strs[i]]; !ok {
			// 情况一：当前字符是新字符（在 map 中不存在）
			// 将其索引存入 map
			m[strs[i]] = i
			// 更新最大长度，i - left + 1 是当前窗口的长度
			maxLength = max(maxLength, i-left+1)
		} else {
			// 情况二：当前字符是重复字符

			// === 问题点：未能正确处理 left 指针的移动 ===
			// 这里的逻辑是：一旦发现重复，就将左边界 `left` 直接移动到重复字符上一次出现的位置 `val`。
			// 这存在一个严重的问题：`left` 指针可能会向左移动（回退）！
			//
			// 让我们用一个例子来说明: s = "abba"
			//
			// 1. i=0, s[0]='a': ok=false. m={'a':0}, max=1, left=0
			// 2. i=1, s[1]='b': ok=false. m={'a':0, 'b':1}, max=2, left=0
			// 3. i=2, s[2]='b': ok=true. val=1.
			//    执行 `left = val`，于是 `left` 变为 1。这是正确的，窗口现在是 "ba"。
			// 4. i=3, s[3]='a': ok=true. val=0 (这是 'a' 上一次出现的位置)。
			//    执行 `left = val`，于是 `left` 变为 0。
			//    出问题了！`left` 指针从 1 回退到了 0！
			//    这导致窗口 `[left, i]` 变成了 `[0, 3]`，即 "abba"，这是一个包含重复字符的无效窗口。
			//    但你的代码没有在 `else` 分支里更新 `maxLength`，所以 `maxLength` 的值没有被错误地更新。
			//    然而，这种 `left` 指针的回退破坏了滑动窗口“只向前不后退”的基本原则，导致了后续计算的混乱。
			//
			// 正确的逻辑应该是：`left` 只能向右移动，或者保持不动。
			// 当发现重复字符时，应该将 `left` 更新为 `max(left, val + 1)`，
			// 确保它不会移动到比当前 `left` 更靠左的位置。
			left = val
		}
	}
	return maxLength
}

// 上一版错误修复后
func lengthOfLongestSubstring(s string) int {
	// map 用于存储字符及其最后一次出现的索引
	charIndexMap := make(map[byte]int)

	maxLen := 0

	// `left` 是滑动窗口的左边界
	left := 0

	// `right` 是滑动窗口的右边界
	for right := 0; right < len(s); right++ {
		currentChar := s[right]

		// 检查当前字符是否在 map 中存在
		if lastIndex, ok := charIndexMap[currentChar]; ok {
			// 如果存在，我们需要决定新的左边界。
			// 新的左边界必须是 "当前 left" 和 "重复字符的下一个位置" 中更靠右的那一个。
			// 这可以防止 left 指针回退。
			// 使用 max 函数可以简化这个逻辑。
			if lastIndex+1 > left {
				left = lastIndex + 1
			}
			// 等价于: left = max(left, lastIndex + 1)
		}

		// 【修复点2】无论如何，都要更新当前字符的最新索引
		charIndexMap[currentChar] = right

		// 【修复点3】无论如何，都要计算当前窗口长度并更新最大值
		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}

//func lengthOfLongestSubstring(s string) int {
//	n := len(s)
//	if n == 0 {
//		return 0
//	}
//	max := 1
//	for first := 0; first < n; first++ {
//		m := make(map[byte]bool)
//		m[s[first]] = true
//		for second := first + 1; second < n; second++ {
//			if !m[s[second]] {
//				m[s[second]] = true
//				tmp := second - first + 1
//				if tmp > max {
//					max = tmp
//				}
//			} else {
//				break
//			}
//		}
//	}
//	return max
//}
//
//// 优化方法
//func lengthOfLongestSubstring(s string) (ans int) {
//	left := 0         // 左指针，表示当前无重复字符子串的起始位置
//	cnt := [128]int{} // 用来记录字符出现的次数，ASCII字符集有128个字符
//	for right, c := range s {
//		cnt[c]++         // 更新字符c的出现次数
//		for cnt[c] > 1 { // 如果字符c的出现次数大于1，表示有重复字符
//			cnt[s[left]]-- // 移动左指针，并减少左边字符的出现次数，直到没有重复字符
//			left++         // 移动左指针
//		}
//		ans = max(ans, right-left+1) // 计算当前无重复字符子串的长度并更新最大值
//	}
//	return
//}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	s := "pwwkew"
	print(lengthOfLongestSubstring(s))
}
