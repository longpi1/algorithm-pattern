package main

/*
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：

void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
说明：

你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。


示例 1：

输入：
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 1, 1, false]

解释：
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false


提示：
1 <= x <= 9
最多调用 100 次 push、pop、peek 和 empty
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）

*/

type MyQueue struct {
	in, out []int
}

func Constructor() MyQueue {
	return MyQueue{in: make([]int, 0), out: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	for len(this.out) != 0 {
		num := this.out[len(this.out)-1]
		this.in = append(this.in, num)
		this.out = this.out[0 : len(this.out)-1]
	}
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	for len(this.in) != 0 {
		num := this.in[len(this.in)-1]
		this.out = append(this.out, num)
		this.in = this.in[0 : len(this.in)-1]
	}
	result := this.out[len(this.out)-1]
	this.out = this.out[0 : len(this.out)-1]
	return result
}

func (this *MyQueue) Peek() int {
	for len(this.in) != 0 {
		num := this.in[len(this.in)-1]
		this.out = append(this.out, num)
		this.in = this.in[0 : len(this.in)-1]
	}
	if len(this.out) == 0 {
		return -1
	}
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

/*type MyQueue struct {
	Queue []int
}


func Constructor() MyQueue {
	queue := make([]int,0)
	return MyQueue{Queue: queue}
}


func (this *MyQueue) Push(x int)  {
	this.Queue = append(this.Queue, x)
}


func (this *MyQueue) Pop() int {
	n := len(this.Queue)
	result := this.Queue[0]
	this.Queue = this.Queue[1:n]
	return result
}


func (this *MyQueue) Peek() int {
	result := this.Queue[0]
	return result
}


func (this *MyQueue) Empty() bool {
	flag := len(this.Queue) == 0
	return flag
}*/

// 上述答案是直接通过切片实现，下面通过双栈
type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	in := make([]int, 0)
	out := make([]int, 0)
	return MyQueue{in: in, out: out}
}

// push需要注意的点，由于是双栈实现，所以需要先将out切片中的数据全部传到in栈中才能添加元素，pop也是一个道理
func (this *MyQueue) Push(x int) {
	for len(this.out) != 0 {
		val := this.out[len(this.out)-1]
		this.out = this.out[:len(this.out)-1]
		this.in = append(this.in, val)
	}
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	for len(this.in) != 0 {
		val := this.in[len(this.in)-1]
		this.in = this.in[:len(this.in)-1]
		this.out = append(this.out, val)
	}
	result := this.out[len(this.out)-1]
	this.out = this.out[0 : len(this.out)-1]
	return result
}

func (this *MyQueue) Peek() int {
	if len(this.in) == 0 && len(this.out) == 0 {
		return -1
	}

	for len(this.in) != 0 {
		val := this.in[len(this.in)-1]
		this.in = this.in[:len(this.in)-1]
		this.out = append(this.out, val)
	}
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
