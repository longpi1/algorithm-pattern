package main


/*
解题思路
根据题意的充分理解，我们可分析如下：

多组相邻重复项，我们无论先删除哪一项，都不会影响最终结果。
删除当前项是需要拿上一项出来对比的，所以我们需要用临时栈存放之前的内容。
当前项和栈顶一致，弹出栈顶抵消即可。若不一致，压入栈留存，供后续使用。
*/



func removeDuplicates(s string) string {
	stack := []byte{}
	for i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
