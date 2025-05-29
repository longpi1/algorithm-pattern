package main

import "fmt"

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
203. 移除链表元素
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

示例 1：
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]
示例 2：

输入：head = [], val = 1
输出：[]
示例 3：

输入：head = [7,7,7,7], val = 7
输出：[]
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{}
	newHead.Next = head
	curNode := head
	for curNode.Next != nil {
		if curNode.Val == val {
			curNode.Next = curNode.Next.Next
		} else {
			curNode = curNode.Next
		}
	}
	return newHead.Next
}

//type ListNode struct {
//	Val int
//	Next *ListNode
//}
/*func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
 for head.Next != nil{
 	if head.Next.Val == val{
		head.Next = head.Next.Next
	}else{
		head = head.Next
	}
 }
	return head
}*/
/*
上述思路错误：
1.不能直接操作头节点，头结点遍历会导致值不断变化，最后返回的就是最后一个值，这里需要用cur作为临时指针节点指向头节点；
2. 需要首先对头结点进行处理，判断头节点是否相等
*/

// 直接使用原来的链表来进行删除操作
func removeElements(head *ListNode, val int) *ListNode {
	// 1.需要首先对头结点进行处理，判断头节点是否相等
	// 这个循环会移除所有值为 val 的前导节点。执行完毕后，head 要么是 nil (如果所有节点都被移除了)，要么指向第一个值不等于 val 的节点。
	for head != nil && head.Val == val {
		head = head.Next
	}
	// 2不能直接操作头节点，头结点遍历会导致值不断变化，最后返回的就是最后一个值，这里需要用cur作为临时指针节点指向头节点；
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 虚拟头节点做法，
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

func main() {
	// 创建链表: 1 -> 2 -> 6 -> 3 -> 4 -> 5 -> 6
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6,
		Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}

	// 移除元素 6
	newHead := removeElements(head, 6)
	fmt.Printf("result: %v", newHead)
}
