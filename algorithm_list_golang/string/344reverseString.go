package main

/*
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。



示例 1：

输入：s = ["h","e","l","l","o"]
输出：["o","l","l","e","h"]
示例 2：

输入：s = ["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]


提示：

1 <= s.length <= 105
s[i] 都是 ASCII 码表中的可打印字符
*/

func reverseString(s []byte) {
	right := len(s) - 1
	left := 0
	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp
		left++
		right--
	}

}

func reverseString(s []byte) {
	right := len(s) - 1
	left := 0
	for left < right {
		tmp := s[right]
		s[right] = s[left]
		s[left] = tmp
		left++
		right--
	}
}
