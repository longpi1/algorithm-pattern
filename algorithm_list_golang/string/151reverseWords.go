package main

import "strings"

/*
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。



示例 1：

输入：s = "the sky is blue"
输出："blue is sky the"
示例 2：

输入：s = "  hello world  "
输出："world hello"
解释：反转后的字符串中不能存在前导空格和尾随空格。
示例 3：

输入：s = "a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。


提示：

1 <= s.length <= 104
s 包含英文大小写字母、数字和空格 ' '
s 中 至少存在一个 单词
*/
// 思路错误。数组切割有问题
//func reverseWords(s string) string {
//	arr := strings.Split(s, " ")
//	left := 0
//	right := len(arr) - 1
//	for left < right {
//		if arr[right] == "" {
//			arr = append(arr[:right], arr[right+1:]...)
//			right--
//			continue
//		}
//		if arr[left] == "" {
//			mov := left + 1
//			right--
//			arr = append(arr[:mov], arr[mov+1:]...)
//			continue
//		}
//
//		arr[left], arr[right] = arr[right], arr[left]
//		left++
//		right--
//
//	}
//	return strings.Join(arr, " ")
//}

func reverseWords(s string) string {
	// 处理空输入
	if s == "" {
		return ""
	}

	// 按空格分割字符串
	arr := strings.Split(s, " ")

	// 如果数组为空或只包含空字符串，直接返回空字符串
	if len(arr) == 0 {
		return ""
	}

	// 反转数组，同时跳过空字符串
	left := 0
	right := len(arr) - 1

	for left < right {
		// 跳过右边的空字符串
		for left < right && arr[right] == "" {
			right--
		}
		// 跳过左边的空字符串
		for left < right && arr[left] == "" {
			left++
		}
		// 如果 left 和 right 相遇，退出循环
		if left >= right {
			break
		}

		// 交换非空字符串
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	// 构建结果数组，只包含非空字符串
	var result []string
	for _, word := range arr {
		if word != "" {
			result = append(result, word)
		}
	}

	// 如果结果数组为空，返回空字符串
	if len(result) == 0 {
		return ""
	}

	// 拼接结果
	return strings.Join(result, " ")
}

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] == "" {
			//ss = append(ss[:i-1], ss[i:]...) 的语法用于从切片 ss 中删除索引为 i-1 的元素。
			//这种方法会将切片分成三个部分，然后重新组合在一起：
			//ss[:i-1] 表示从切片开头到索引 i-1（不包括 i-1）的部分。
			//ss[i:] 表示从索引 i 到切片末尾的部分。
			//使用 ... 表示将第 1 和第 2 部分重新组合在一起。
			ss = append(ss[:i], ss[i+1:]...)
		}
	}
	right := len(ss) - 1
	left := 0
	for left < right {
		ss[left], ss[right] = ss[right], ss[left]
		left++
		right--
	}
	return strings.Join(ss, " ")
}

// 其他方法
func reverseWords(s string) string {
	t := strings.Fields(s)          //使用该函数可切割单个/多个空格，提取出单词
	for i := 0; i < len(t)/2; i++ { //遍历数组的前半段直接交换即可
		t[i], t[len(t)-1-i] = t[len(t)-1-i], t[i]
	}
	new := strings.Join(t, " ") //重新使用空格连接多个单词
	return (new)
}

// 不用strings工具类
func reverseWords(s string) string {
	//1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	//删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	//删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	//2.反转整个字符串
	reverse(&b, 0, len(b)-1)
	//3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}

func main() {
	println(reverseWords("a good   example"))
}
