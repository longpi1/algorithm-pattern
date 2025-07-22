package main

import "math"

/*
155. 最小栈

设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:

MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。

示例 1:
输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
*/
type MinStack1 struct {
}

func Constructor() MinStack {

}

func (this *MinStack) Push(val int) {

}

func (this *MinStack) Pop() {

}

func (this *MinStack) Top() int {

}

func (this *MinStack) GetMin() int {

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
/*type MinStack struct {
	stack []int
	minStack []int
}


func Constructor() MinStack {
	return MinStack{stack: []int{}, minStack: []int{math.MaxInt64}}
}


func (this *MinStack) Push(val int)  {

}


func (this *MinStack) Pop()  {

}


func (this *MinStack) Top() int {
	return this.stack[0]
}


func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1]
}

func min(b int, a int) int{
	if a > b {
		return b
	}else {
		return a
	}
}*/

/*
上述思路错误，
题目只要求常数时间内获取最小元素即可
*/

// MinStack 结构体表示一个栈数据结构，具有获取栈中最小元素的特殊功能。
type MinStack struct {
	stack    []int // 主栈，用于存储元素
	minStack []int // 辅助栈，用于存储最小元素
}

// Constructor 构造函数，创建一个新的 MinStack 实例并初始化栈。
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64}, // 将 minStack 初始化为最大整数
	}
}

// Push 将新元素加入栈中，并同时更新 minStack 以存储目前为止的最小元素。
func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)                 // 将元素加入主栈
	top := this.minStack[len(this.minStack)-1]         // 获取 minStack 的栈顶元素
	this.minStack = append(this.minStack, min(x, top)) // 将 x 和栈顶元素的最小值加入 minStack
}

// Pop 从主栈和 minStack 中弹出栈顶元素。
func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]          // 从主栈中弹出栈顶元素
	this.minStack = this.minStack[:len(this.minStack)-1] // 从 minStack 中弹出栈顶元素
}

// Top 返回主栈的栈顶元素。
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1] // 获取主栈的栈顶元素
}

// GetMin 返回栈中的最小元素。
func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1] // 获取 minStack 的栈顶元素
}

// min 是一个辅助函数，用于返回两个整数的最小值。
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
