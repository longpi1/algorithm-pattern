package main

/*
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

示例 1：

输入：s = "abcdefg", k = 2
输出："bacdfeg"
示例 2：

输入：s = "abcd", k = 2
输出："bacd"

提示：

1 <= s.length <= 104
s 仅由小写英文组成
1 <= k <= 104
*/
func reverseStr(s string, k int) string {
	arr := []byte(s)
	l := len(arr) - 1
	n := l / (2 * k)
	//result := []byte
	for i := 0; i <= n; i++ {
		left := i * 2 * k

		right := i*2*k + k - 1
		if i == n && l-i*2*k < k {
			right = l
		}
		for left < right {
			tmp := arr[left]
			arr[left] = arr[right]
			arr[right] = tmp
			left++
			right--
		}
	}
	return string(arr)
}

// 需要注意字符串与数组转换
func reverseStr(s string, k int) string {
	//这里需要用byte数组
	ss := []byte(s)
	//ss := []string(s)

	length := len(s)
	for i := 0; i < length; i += 2 * k {
		if i+k <= length {
			reverse(ss[i : i+k])
		} else {
			reverse(ss[i:length])
		}
	}
	//最后用string（）返回即可
	return string(ss)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
