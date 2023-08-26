package main




/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。



示例 1：


输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：

输入：head = []
输出：[]
示例 3：

输入：head = [1]
输出：[1]

*/


/*func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}

	cur := head
	//var pre *ListNode
	for cur != nil && cur.Next != nil{
		next := cur.Next.Next
		cur.Next.Next = cur

		cur.Next = next

		cur = next
	}

	return dummyHead.Next
}*/

/*
上述答案错误。存在以下几个问题：
1.存在一个逻辑错误，会导致链表断裂。在交换两个相邻节点时：cur.Next.Next = cur，cur.Next = next，您需要将当前节点 cur 的 Next 指针指向下一个节点，但是在这行代码中，将 Next 指针指向了当前节点 cur 自身，这会导致链表中断裂，丢失原来的连接关系。
2.需要利用起头节点，不然不好处理前面节点的转换
3.cur = next,这里不应该再指向之前的临时变量，应该是新的位置
*/

func swapPairs(head *ListNode) *ListNode {
	// 设置一个虚拟头结点，将虚拟头结点指向head，这样方面后面做删除操作
	dummyHead := &ListNode{Next: head}

	cur := dummyHead

	for cur.Next != nil && cur.Next.Next != nil{
		// 画图方便理解，步骤主要如下
		// 记录临时节点
		tmp := cur.Next
		tmp1 := cur.Next.Next.Next
		//从0开始举例子
		//1.首先0指向节点2；
		cur.Next = cur.Next.Next
		//2.节点2指向节点1，以为连接断了所以用的之前的临时节点进行连接
		cur.Next.Next = tmp
		//3.节点1指向节点3，以为连接断了所以用的之前的临时节点进行连接
		cur.Next.Next.Next = tmp1
		// cur移动两位，准备下一轮交换
		cur = cur.Next.Next
	}

	return dummyHead.Next
}


// 递归版本,画图调式之后更清晰，递归最主要找边界条件与非边界条件，最后得到就i是最小子问题的结果
func swapPairs(head *ListNode) *ListNode {
	// # 递归边界
	if head == nil || head.Next == nil {
		// 不足两个节点，无需交换
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}


func main(){
	head := &ListNode{Val:1,Next: &ListNode{Val:2,Next: &ListNode{Val:3,Next: &ListNode{Val:4,Next: nil}}}}

	swapPairs(head)
}