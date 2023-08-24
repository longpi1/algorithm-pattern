package main

/*

你可以选择使用单链表或者双链表，设计并实现自己的链表。

单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。

如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。

实现 MyLinkedList 类：

MyLinkedList() 初始化 MyLinkedList 对象。
int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。


示例：

输入
["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
[[], [1], [3], [1, 2], [1], [1], [1]]
输出
[null, null, null, null, 2, null, 3]

解释
MyLinkedList myLinkedList = new MyLinkedList();
myLinkedList.addAtHead(1);
myLinkedList.addAtTail(3);
myLinkedList.addAtIndex(1, 2);    // 链表变为 1->2->3
myLinkedList.get(1);              // 返回 2
myLinkedList.deleteAtIndex(1);    // 现在，链表变为 1->3
myLinkedList.get(1);              // 返回 3
*/
/*
type MyLinkedList struct {
	Val int
	Next *MyLinkedList
}


func Constructor() MyLinkedList {
	linkedList := new(MyLinkedList)
	return *linkedList
}


func (this *MyLinkedList) Get(index int) int {
	count := 0

	cur := this
	for count <= index {
		count++
		cur = cur.Next
	}
	return cur.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
 head :=	MyLinkedList{Val: val}
 head.Next = this
}


func (this *MyLinkedList) AddAtTail(val int)  {
	head :=	MyLinkedList{Val: val}
	cur := this
	for cur != nil{
		if cur.Next == nil{
			cur.Next = &head
		}
		cur = cur.Next
	}
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	count := 0
	tmp := MyLinkedList{Val: val}
	cur := this
	for count < index {
		count++
		cur = cur.Next
	}
	cur.Next = &tmp
	tmp.Next = cur.Next.Next
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	count := 0

	cur := this
	for count < index {
		count++
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
}
*/

/*
上述思路错误：
1.
*/

// 节点结构体
type Node struct {
	Val int
	Next *Node
}
// MylinkedList是一个对象，需要去实现LinkedList接口的所有方法
type MyLinkedList struct {
	DummyHead *Node //虚拟头节点
	Size int //链表长度
}

//创建一个链表
func Constructor() MyLinkedList { //成功
	// 创建一个虚拟头节点，非真正的头节点(真正的头节点是数组第一个节点)
	dummyHead := &Node{
		Val : -1,
		Next : nil,
	}
	return MyLinkedList{DummyHead: dummyHead,Size: 0}
}
// 获取index位置的元素Val
// 获取到第index个节点数值，如果index是非法数值直接返回-1， 注意index是从0开始的，第0个节点就是头结点
func (this *MyLinkedList) Get(index int) int {
	// 判断index是否非法
	if (index < 0) || (index > (this.Size - 1)) {
		return -1
	}
	// 查找
	var cur  = this.DummyHead
	// dummy节点的index = -1
	for i := 0;i <= index;i++ {
		//找到index为 index的节点
		cur = cur.Next
		//0,1,2,3,4....index
	}
	return cur.Val
}

// 在头节点前面再次添加一个节点
func (this *MyLinkedList) AddAtHead(val int)  {
	// 在dummy节点后面直接添加一个节点
	var newNode  = &Node{Val:val, Next:nil}
	//所有变量都要显示的初始化
	newNode.Next = this.DummyHead.Next
	this.DummyHead.Next = newNode
	this.Size++
}


// 在尾结点添加节点
func (this *MyLinkedList) AddAtTail(val int)  {
	var newNode  = &Node{val, nil}
	cur := this.DummyHead
	for cur.Next != nil { //找到末尾节点
		cur = cur.Next
	}
	cur.Next = newNode //新元素添加到末尾节点后面
	this.Size++
}

// 在index节点前面添加一个节点
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.Size {
		return
	}else if index == this.Size {
		//直接添加到末尾
		this.AddAtTail(val)
		return
	}else if index < 0 {
		index = 0
	}
	var newNode  = &Node{val, nil}
	cur := this.DummyHead
	for i:=0; i<=index-1; i++ {
		//找到index为 index-1的节点,如果index=0，则不会进入循环，直接插入到dummy后面
		cur = cur.Next
	}
	newNode.Next = cur.Next
	cur.Next = newNode
	this.Size++
}
// 删除index节点
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	// 判断是否有效
	if index > this.Size-1 || index < 0 {
		return
	}
	cur := this.DummyHead
	for i := 0; i <= index-1; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	this.Size--
}


// 双链表做法：
type node struct {
	val        int
	next, prev *node
}

type MyLinkedList struct {
	head, tail *node
	size       int
}

func Constructor() MyLinkedList {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head
	return MyLinkedList{head, tail, 0}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	var curr *node
	if index+1 < l.size-index {
		curr = l.head
		for i := 0; i <= index; i++ {
			curr = curr.next
		}
	} else {
		curr = l.tail
		for i := 0; i < l.size-index; i++ {
			curr = curr.prev
		}
	}
	return curr.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) AddAtIndex(index, val int) {
	if index > l.size {
		return
	}
	index = max(0, index)
	var pred, succ *node
	if index < l.size-index {
		pred = l.head
		for i := 0; i < index; i++ {
			pred = pred.next
		}
		succ = pred.next
	} else {
		succ = l.tail
		for i := 0; i < l.size-index; i++ {
			succ = succ.prev
		}
		pred = succ.prev
	}
	l.size++
	toAdd := &node{val, succ, pred}
	pred.next = toAdd
	succ.prev = toAdd
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	var pred, succ *node
	if index < l.size-index {
		pred = l.head
		for i := 0; i < index; i++ {
			pred = pred.next
		}
		succ = pred.next.next
	} else {
		succ = l.tail
		for i := 0; i < l.size-index-1; i++ {
			succ = succ.prev
		}
		pred = succ.prev.prev
	}
	l.size--
	pred.next = succ
	succ.prev = pred
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}


// func main(){
//     obj := Constructor()
// 	obj.AddAtHead(1)
// 	obj.AddAtTail(2)
// 	obj.AddAtIndex(0,0)
// 	obj.DeleteAtIndex(0)
// 	fmt.Println(obj.Get(0),obj.Get(1),obj.Get(2))

// }
