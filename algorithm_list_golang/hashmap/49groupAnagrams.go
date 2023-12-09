package main

import (
	"fmt"
	"sort"
)

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:
输入: strs = [""]
输出: [[""]]
示例 3:
输入: strs = ["a"]
输出: [["a"]]

*/


/*func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, s := range strs {
		sort.Slice(s, func(i, j int) bool {
			return  s[i] > s[j]
		})
		m[s] = append(m[s], s)

	}
	result := make([][]string, 0,len(m))
	for strs :=  range m {
		result = append(result, m[strs])
	}
	return result

}
*/

/*
上述代码存在以下几个问题：
sort.Slice 函数的使用不正确：sort.Slice 函数用于对切片进行排序，但您尝试对字符串 s 调用 sort.Slice，这是不允许的。应该将字符串转换为字符切片并对切片进行排序。

错误的结果存储：您尝试将已排序的字符串作为键来存储结果，这会导致您只能获得一个相同的字符串，而不是一组字母异位词。

结果提取时的错误：在提取结果时，您使用了 range 循环，但循环中的变量 strs 是一个字符串而不是切片。您应该使用切片的键来提取结果。
*/

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	// 遍历输入字符串数组
	for _, s := range strs {
		// 将字符串转换为字符切片，然后对切片进行按字母逆序排序
		str := []byte(s)
		sort.Slice(str, func(i, j int) bool {
			return str[i] > str[j]
		})
		// 将排序后的切片转换回字符串，作为键来存储字母异位词
		sortedStr := string(str)
		// 将原始字符串添加到对应键的切片中
		m[sortedStr] = append(m[sortedStr], s)
	}

	// 创建一个用于存储结果的二维字符串切片
	result := make([][]string, 0, len(m))
	// 遍历存储字母异位词的映射
	for _, group := range m {
		// 将每个字母异位词组的切片添加到结果中
		result = append(result, group)
	}
	return result
}


func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("%v", groupAnagrams(strs))
}