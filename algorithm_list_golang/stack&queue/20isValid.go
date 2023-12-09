package main

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。


示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false


提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/
/*
解题思路：
算法原理

栈先入后出特点恰好与本题括号排序特点一致，即若遇到左括号入栈，遇到右括号时将对应栈顶左括号出栈，则遍历完所有括号后 stack 仍然为空；
建立哈希表 dic 构建左右括号对应关系：key左括号，value右括号；这样查询 222 个括号是否对应只需 O(1) 时间复杂度；建立栈 stack，遍历字符串 s 并按照算法流程一一判断。
算法流程

如果 c 是左括号，则入栈 push；
否则通过哈希表判断括号对应关系，若 stack 栈顶出栈括号 stack.pop() 与当前遍历括号 c 不对应，则提前返回 false。
提前返回 false

提前返回优点： 在迭代过程中，提前发现不符合的括号并且返回，提升算法效率。
解决边界问题：
栈 stack 为空： 此时 stack.pop() 操作会报错；因此，我们采用一个取巧方法，给 stack 赋初值 ??? ，并在哈希表 dic 中建立 key:′?′，value:′?′
′
  的对应关系予以配合。此时当 stack 为空且 c 为右括号时，可以正常提前返回 falsefalsefalse；
字符串 s 以左括号结尾： 此情况下可以正常遍历完整个 s，但 stack 中遗留未出栈的左括号；因此，最后需返回 len(stack) == 1，以判断是否是有效的括号组合。
复杂度分析

时间复杂度 O(N))：正确的括号组合需要遍历 111 遍 s；
空间复杂度 O(N)：哈希表和栈使用线性的空间大小。

*/
func isValid(s string) bool {
	m := make(map[byte]byte,3)
	m['('] = ')'
	m['{'] = '}'
	m['['] = ']'
	stack := make([]byte, 0)
	str := []byte(s)
	for i :=0; i < len(str); i++ {
		val := str[i]
		if len(stack) == 0 {
			//当栈为空时，出现右括号则直接退出
			if _,flag := m[val]; !flag {
				return false
			}
		}
		// 左括号直接添加，右括号则需要判断
		if _,flag := m[val];flag {
			stack = append(stack, val)
		}else{
			pop := stack[len(stack)-1]
			if val != m[pop] {
				return false
			}else{
				stack = stack[0:len(stack)-1]
			}
		}
	}
	if len(stack) > 0 {
		return false
	}

	return true
}

func main(){
	s := "()[]{"
	print(isValid(s))
}
