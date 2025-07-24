package main

import (
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

/*
下述代码逻辑错误。总结一下主要问题：
无法处理嵌套：这是最根本的问题。你的线性扫描和简单的 queue 无法处理 k[...k[...]] 这样的结构。
状态保存缺失：当遇到 [ 时，你没有机制来保存当前的重复次数和已经构建的字符串。
数字解析错误：
将字符的 ASCII 码错误地当作数字值。
无法处理多位数。
字符串拼接逻辑错误：
queue 混淆了不同层级的字符串。
结果直接拼接到全局 result，破坏了层级关系。
不完整的处理：忽略了括号外的、以及循环结束时仍在 queue 中的字符。
*/
func decodeString(s string) string {
	// --- 问题 1: 数据结构选择不当 ---
	// 你用 `queue` 来存储括号内的字符串，用 `stack` 来存储原始输入。
	// 但整个算法的核心是处理嵌套结构，这需要一个真正的“栈”来保存
	// 遇到 `[` 前的字符串和数字。你的 `queue` 和 `stack` 并未实现这个功能。
	queue := make([]byte, 0)
	stack := []byte(s)

	// `tmpNum` 用来存数字，但它只能存一个一位数，这是个大问题。
	tmpNum := 0
	result := ""

	// --- 问题 2: 循环条件和遍历方式错误 ---
	// for i := 0; i < len(stack)-1; i++ {
	// 1. `len(stack)-1` 会导致最后一个字符被忽略。
	// 2. 这种简单的线性 for 循环无法处理嵌套结构。它只能从左到右扫描，
	//    没有递归或真正的栈操作来处理内部的解码。
	for i := 0; i < len(stack)-1; i++ {
		tmp := stack[i]
		if tmp == '[' {
			// 你只是简单地跳过了 `[`，但没有做任何处理，比如保存之前的状态。
			continue
		}
		if tmp == ']' {
			// 当遇到 `]` 时，你用全局的 `result` 去拼接，这完全是错误的。
			// 解码应该发生在当前层级。`result` 应该是最终的结果。
			// 例如，对于 "a[b]c"，当处理到 `]` 时，`result` 应该是 "ab"，
			// 而不是直接把 `b` 加到最终结果里。
			//
			// 此外，如果遇到嵌套 `3[a2[c]]`，当遇到第一个 `]` 时，
			// 你如何知道要重复的是 `c` 而不是 `ac`？你的 `queue` 无法区分。
			result = result + strings.Repeat(string(queue), tmpNum)
			// 清空 `queue`，准备下一个片段，但这对于嵌套是无效的。
			queue = make([]byte, 0)
			continue
		}

		// --- 问题 3: 数字处理逻辑严重错误 ---
		if tmp >= '0' && tmp <= '9' { // ASCII 比较
			// `tmpNum = int(tmp)` 是错误的！
			// `tmp` 是一个 byte，它的值是字符的 ASCII 码。
			// 例如，字符 '3' 的 ASCII 码是 51。所以 `tmpNum` 会变成 51，而不是 3。
			// 正确的转换应该是 `int(tmp - '0')`。
			//
			// 更严重的是，这个逻辑无法处理多位数，比如 "12[a]"。
			// 当遇到 '1'，tmpNum 变成 1。当遇到 '2'，tmpNum 就被覆盖成了 2。
			// 你需要一个循环来解析完整的数字。
			tmpNum = int(tmp)
			continue
		}

		// --- 问题 4: 字母处理逻辑错误 ---
		// `queue` 只是简单地追加所有遇到的字母。
		// 对于 "3[a]2[b]"，当处理到 'b' 时，`queue` 会变成 `['a', 'b']`，
		// 这显然是错误的。`queue` 应该在遇到 `[` 时被清空或压栈。
		queue = append(queue, tmp)
	}

	// --- 问题 5: 最终结果不完整 ---
	// 循环结束后，如果 `queue` 中还有未处理的字符（比如 "a2[b]c" 中的 'c'），
	// 或者整个字符串就没有括号（比如 "abc"），它们都被忽略了。
	// 最终的 `result` 只包含了被 `]` 触发拼接的部分。
	return result
}

// 需要梳理好什么时候入栈和出栈处理

// decodeString 函数用于解码包含数字和字符的字符串，按照指定规则进行重复和拼接。
func decodeString(s string) string {
	// numStack: 存重复次数
	// strStack: 存遇到 '[' 前的字符串
	numStack := []int{}
	strStack := []string{}

	currentNum := 0
	currentRes := ""

	for _, char := range s {
		if char >= '0' && char <= '9' {
			// 1. 如果是数字，更新 currentNum (处理多位数)
			currentNum = currentNum*10 + int(char-'0')
		} else if char == '[' {
			// 2. 如果是 '['
			// a. 将当前数字和字符串分别入栈
			numStack = append(numStack, currentNum)
			strStack = append(strStack, currentRes)
			// b. 重置，为括号内的内容做准备
			currentNum = 0
			currentRes = ""
		} else if char == ']' {
			// 3. 如果是 ']'
			// a. 从栈中弹出数字和之前的字符串
			repeatTimes := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevRes := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			// b. 核心操作：构建新的字符串
			currentRes = prevRes + strings.Repeat(currentRes, repeatTimes)
		} else {
			// 4. 如果是字母，直接拼接到当前结果
			currentRes += string(char)
		}
	}

	return currentRes
}

/*
该函数通过使用两个栈，分别用于存储数字和字符串，以及一个变量用于存储当前结果字符串，实现了对给定字符串的解码。
在遍历输入字符串的过程中，根据遇到的字符进行相应的操作，最终得到解码后的字符串
*/

func main() {
	s := "3[a2[c]]"
	print(decodeString(s))
}
