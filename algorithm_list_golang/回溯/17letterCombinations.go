package main

import "fmt"

/*
17. 电话号码的字母组合

给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：
输入：digits = ""
输出：[]
示例 3：
输入：digits = "2"
输出：["a","b","c"]
*/

// 错误代码，留下记录用于后续再写的时候发现错误以及加深印象
func letterCombinations(digits string) []string {
	digitToChars := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	path := make([]rune, 0)
	res := make([]string, 0)
	digitsBytes := []rune(digits)
	dfs := func(index int) {}
	dfs = func(index int) {
		tmp := make([]rune, 0)
		copy(tmp, path)
		fmt.Println(string(path))
		res = append(res, string(tmp))
		for i := index; i < len(digitsBytes); i++ {
			s := digitToChars[digitsBytes[i]]
			sArr := []rune(s)
			for j := 0; j < len(sArr); j++ {
				//fmt.Println(string(sArr[j]))
				path = append(path, (sArr[j]))
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

// letterCombinations 返回电话号码字母组合
func letterCombinations(digits string) []string {
	// 如果输入为空字符串，直接返回空结果
	if digits == "" {
		return []string{}
	}

	digitToChars := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := make([]string, 0) // 存储所有找到的有效组合
	// path 用来构建当前组合，通常用 []byte 或 []rune，或者 strings.Builder
	// 这里使用 []byte 更常见，因为最终转 string
	path := make([]byte, 0, len(digits)) // 预分配容量，减少 append 时的重新分配

	// 将 digits 字符串转换为 []byte，方便索引访问
	// 假设 digits 仅包含 ASCII 数字字符
	digitsBytes := []byte(digits)

	// dfs 是回溯函数
	// index: 当前正在处理 digits 字符串的哪个字符（数字）
	var dfs func(index int)
	dfs = func(index int) {
		// 1. 终止条件 / 找到解的条件
		// 如果 index 等于 digits 的长度，说明已经处理完所有数字，path 构成一个完整的组合
		if index == len(digitsBytes) {
			res = append(res, string(path)) // 将当前完整组合添加到结果集
			return
		}
		//
		// 获取当前数字对应的字符集合
		digit := digitsBytes[index]
		chars, ok := digitToChars[rune(digit)] // map 的键是 rune，所以需要转换
		if !ok {
			// 如果有非法数字（如 '0' 或 '1'），可以跳过或处理错误
			// 题目通常假设输入是有效的 '2'-'9'
			// 对于 '0'/'1' 的情况，它们没有对应字母，所以这里应该直接返回或者跳过
			// 简单起见，这里假设输入合法
			return
		}

		// 2. 遍历所有可能的选择 (当前数字对应的所有字符)
		for _, char := range chars { // 遍历字符串 chars 中的每个 rune
			// 3. 做选择 (将当前字符加入 path)
			path = append(path, byte(char)) // 将 rune 转换为 byte 添加，如果 char 是 ASCII 字符

			// 4. 递归调用，进入下一层决策 (处理下一个数字)
			dfs(index + 1)

			// 5. 撤销选择 (回溯)
			// 将当前字符从 path 中移除，回到上一个决策点
			path = path[:len(path)-1]
		}
	}

	// 初始调用回溯函数，从处理 digits 的第一个字符 (索引 0) 开始
	dfs(0)
	return res
}

func main() {
	digits := "23"
	combinations := letterCombinations(digits)
	fmt.Printf("result: %v", combinations)

}

//
//var (
//	digitToChars = map[rune]string{
//		'2': "abc",
//		'3': "def",
//		'4': "ghi",
//		'5': "jkl",
//		'6': "mno",
//		'7': "pqrs",
//		'8': "tuv",
//		'9': "wxyz",
//	}
//	path []byte
//	res  []string
//)
//
//func letterCombinations(digits string) []string {
//	path, res = make([]byte, 0), make([]string, 0)
//
//	if digits == "" {
//		return res
//	}
//
//	dfs(digits, 0)
//	return res
//}
//
//func dfs(digits string, start int) {
//	if len(path) == len(digits) { // 终止条件，字符串长度等于 digits 的长度
//		tmp := string(path)
//		res = append(res, tmp)
//		return
//	}
//
//	digit := digits[start]             // 取下一个数字
//	chars := digitToChars[rune(digit)] // 取数字对应的字符集
//
//	for j := 0; j < len(chars); j++ {
//		path = append(path, chars[j])
//		dfs(digits, start+1)
//		path = path[:len(path)-1]
//	}
//
//}
