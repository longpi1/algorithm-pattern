package main

import (
	"strconv"
	"strings"
)

/*
394. 字符串解码
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例 1：
输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：
输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：
输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：
输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"
*/

// 需要梳理好什么时候入栈和出栈处理

// decodeString 函数用于解码包含数字和字符的字符串，按照指定规则进行重复和拼接。
func decodeString(s string) string {
	numStack := []int{}       // 用于存储数字的栈
	strStack := []string{}    // 用于存储字符串的栈
	num := 0                  // 当前数字
	result := ""              // 当前结果字符串

	// 遍历输入字符串中的每个字符
	for _, char := range s {
		if char >= '0' && char <= '9' {
			n, _ := strconv.Atoi(string(char))   // 将字符转换为数字
			num = num*10 + n                      // 构建多位数字
		} else if char == '[' {
			strStack = append(strStack, result)    // 将当前结果字符串压入字符串栈
			result = ""                            // 重置结果字符串
			numStack = append(numStack, num)       // 将当前数字压入数字栈
			num = 0                                // 重置当前数字
		} else if char == ']' {
			count := numStack[len(numStack)-1]     // 获取当前应该重复的次数
			numStack = numStack[:len(numStack)-1]   // 弹出数字栈的栈顶元素
			str := strStack[len(strStack)-1]       // 获取上一层的字符串
			strStack = strStack[:len(strStack)-1]   // 弹出字符串栈的栈顶元素
			result = string(str) + strings.Repeat(result, count)  // 将当前字符串重复 count 次并拼接到上一层的字符串上
		} else {
			result += string(char)   // 遇到字符时直接拼接到结果字符串中
		}
	}
	return result
}

/*
该函数通过使用两个栈，分别用于存储数字和字符串，以及一个变量用于存储当前结果字符串，实现了对给定字符串的解码。
在遍历输入字符串的过程中，根据遇到的字符进行相应的操作，最终得到解码后的字符串
*/

func main()  {
 	s := "3[a2[c]]"
 	print(decodeString(s))
}