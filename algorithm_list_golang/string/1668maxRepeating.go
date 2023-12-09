package main

/*
给你一个字符串 sequence ，如果字符串 word 连续重复 k 次形成的字符串是 sequence 的一个子字符串，那么单词 word 的 重复值为 k 。单词 word 的 最大重复值 是单词 word 在 sequence 中最大的重复值。如果 word 不是 sequence 的子串，那么重复值 k 为 0 。

给你一个字符串 sequence 和 word ，请你返回 最大重复值 k 。



示例 1：

输入：sequence = "ababc", word = "ab"
输出：2
解释："abab" 是 "ababc" 的子字符串。
示例 2：
输入：sequence = "ababc", word = "ba"
输出：1
解释："ba" 是 "ababc" 的子字符串，但 "baba" 不是 "ababc" 的子字符串。
示例 3：
输入：sequence = "ababc", word = "ac"
输出：0
解释："ac" 不是 "ababc" 的子字符串。

提示：
1 <= sequence.length <= 100
1 <= word.length <= 100
sequence 和 word 都只包含小写英文字母。
*/

//标准库实现
//func maxRepeating(sequence string, word string) int {
//	for k := len(sequence) / len(word); k > 0; k-- {
//		if strings.Contains(sequence, strings.Repeat(word, k)) {
//			return k
//		}
//	}
//	return 0
//}



func maxRepeating(sequence string, word string) int {
	count := 0
	start := 0
	for  start < len(sequence)-len(word) +1{
		if sequence[start] == word[0]{
			if len(word) == 1 {
				return 1
			}
			for j := 1; j < len(word) ; j++ {
				if sequence[start+j] != word[j] {
					start ++
					break

				}
				if j == len(word) -1 {
					start += len(word)
					count ++
				}
			}

		}else{
			start ++
		}

	}
	return count
}

func main() {
	sequence := "aaabaaaabaaabaaaabaaaabaaaabaaaaba"
	word := "aaaba"
	println(maxRepeating(sequence,word))
}