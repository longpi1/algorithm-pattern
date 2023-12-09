package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

/*func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: nil}
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil {
		if cur.Val == val{
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}

	}
	return head
}*/

/*
上述代码犯了两个错误：
1.不应该用cur != nil和cur.Val == val ，因为我们这里要用到cur.Next = cur.Next.Next，如果是末尾元素会造成空指针报错因为没有对应的cur.Next.Next ，所以要判断cur.Next才行
2.不应该return head，应该返回的是dummyHead.Next。没有使用更新后的 dummyHead.Next 指针。这可能导致内存泄漏而且并没有更新链表，因为删除节点时原始 head 指针不再指向链表的头部。
*/

func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: nil}
	dummyHead.Next = head
	cur := dummyHead
	//1.不应该用cur != nil和cur.Val == val ，因为我们这里要用到cur.Next = cur.Next.Next，如果是末尾元素会造成空指针报错因为没有对应的cur.Next.Next ，所以要判断cur.Next才行
	for cur.Next != nil {
	//for cur != nil {
		if cur.Next.Val == val{
		//if cur.Val == val{
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}

	}
	//不应该return head，应该返回的是dummyHead.Next。没有使用更新后的 dummyHead.Next 指针。这可能导致内存泄漏而且并没有更新链表，因为删除节点时原始 head 指针不再指向链表的头部。
	//return head
	return dummyHead.Next
}


func main(){
	// 创建链表: 1 -> 2 -> 6 -> 3 -> 4 -> 5 -> 6
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6,
		Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}

	// 移除元素 6
	newHead := deleteNode(head, 6)
	fmt.Printf("result: %v", newHead)
}