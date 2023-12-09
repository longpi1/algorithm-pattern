package main

import "fmt"

/*
139. 单词拆分
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

示例 1：
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
示例 2：
输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
     注意，你可以重复使用字典中的单词。
示例 3：
输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
*/

// wordBreak 检查字符串s是否能被拆分为wordDict中的单词
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true // 将单词字典转换为集合，方便检查单词是否存在
	}
	dp := make([]bool, len(s)+1) // dp[i]表示s的前i个字符是否可以被拆分为wordDict中的单词
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] { // 如果前j个字符可以被拆分，且s[j:i]在单词字典中
				dp[i] = true // 则前i个字符也可以被拆分
				break
			}
		}
	}
	return dp[len(s)] // 返回s是否可以被拆分为wordDict中的单词
}

// wordBreak 动态规划求解背包问题版本
func wordBreak(s string, wordDict []string) bool {
	dp := make([]int, len(s)+1) // dp[i]表示装满背包s的前i位字符的方式数量
	dp[0] = 1
	for i := 0; i <= len(s); i++ { // 背包
		for j := 0; j < len(wordDict); j++ { // 物品
			if i >= len(wordDict[j]) && wordDict[j] == s[i-len(wordDict[j]):i] {
				dp[i] += dp[i-len(wordDict[j])] // 更新背包
			}
		}
	}
	return dp[len(s)] > 0 // 如果装满背包的方式数量大于0，返回true，否则返回false
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict)) // Output: true
}
