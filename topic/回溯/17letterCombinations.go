package main


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

var (
	digitToChars = map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	path []byte
	res  []string
)

func letterCombinations(digits string) []string {
	path, res = make([]byte, 0), make([]string, 0)

	if digits == "" {
		return res
	}

	dfs(digits, 0)
	return res
}

func dfs(digits string, start int) {
	if len(path) == len(digits) { // 终止条件，字符串长度等于 digits 的长度
		tmp := string(path)
		res = append(res, tmp)
		return
	}

	digit := digits[start] // 取下一个数字
	chars := digitToChars[rune(digit)] // 取数字对应的字符集

	for j := 0; j < len(chars); j++ {
		path = append(path, chars[j])
		dfs(digits, start+1)
		path = path[:len(path)-1]
	}

}
