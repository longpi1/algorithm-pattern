package main

import "strconv"

/*

给你一个字符串数组 tokens ，表示一个根据 逆波兰表示法 表示的算术表达式。

请你计算该表达式。返回一个表示表达式值的整数。

注意：

有效的算符为 '+'、'-'、'*' 和 '/' 。
每个操作数（运算对象）都可以是一个整数或者另一个表达式。
两个整数之间的除法总是 向零截断 。
表达式中不含除零运算。
输入是一个根据逆波兰表示法表示的算术表达式。
答案及所有中间计算结果可以用 32 位 整数表示。


示例 1：

输入：tokens = ["2","1","+","3","*"]
输出：9
解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
示例 2：

输入：tokens = ["4","13","5","/","+"]
输出：6
解释：该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6
示例 3：

输入：tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
输出：22
解释：该算式转化为常见的中缀算术表达式为：
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/

func evalRPN(tokens []string) int {
	m := make(map[string]bool)
	m["/"] = true
	m["*"] = true
	m["+"] = true
	m["-"] = true
	stack := make([]string,0)
	for i :=0; i< len(tokens); i++ {
		if !m[tokens[i]] {
			stack = append(stack, tokens[i])
			continue
		}
		n :=len(stack)
		//1. 应该使用strconv.Atoi将字符串转换为整数  !!! 刚开始做的时候忘记了
		val1,_ := strconv.Atoi(stack[n-1])
		val2,_ := strconv.Atoi(stack[n-2])
		stack = stack[:n-2]
		switch tokens[i] {
		case "/":
			// 运算符操作顺序错误，这里应该是val2 在左边，val1在右边，因为后续遍历是左右中，左边的数字应该是stack[n-2]
			result := val2 / val1
			//result := val1 / val2
			//2. 整数转字符串要用strconv.itoa 方法或者使用 fmt.Sprintf，不能直接用string() ，string函数的参数若是一个整型数字，它将该整型数字转换成ASCII码值等于该整形数字的字符。
			stack = append(stack, strconv.Itoa(result))
		case "*":
			result := val1 * val2
			stack = append(stack, strconv.Itoa(result))
		case "+":
			result := val1 + val2
			stack = append(stack, strconv.Itoa(result))
		case "-":
			// 运算符操作顺序错误，这里应该是val2 在左边，val1在右边，因为后续遍历是左右中，左边的数字应该是stack[n-2]
			//result := val1 - val2
			result := val2 - val1
			stack = append(stack, strconv.Itoa(result))
		}

	}
	result, _ := strconv.Atoi(stack[0])
	return result
}

// 上述答案刚开始存在两个问题：
// 1. 应该使用strconv.Atoi将字符串转换为整数  !!! 刚开始做的时候忘记了
// 2. 整数转字符串要用strconv.itoa 方法或者使用 fmt.Sprintf，不能直接用string() ，string函数的参数若是一个整型数字，它将该整型数字转换成ASCII码值等于该整形数字的字符。
// 3. 运算符操作顺序错误，这里应该是val2 在左边，val1在右边，因为后续遍历是左右中，左边的数字应该是stack[n-2]

func main(){
	s := []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}
	println(evalRPN(s))
}