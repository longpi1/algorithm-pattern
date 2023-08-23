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

//  直接使用原来的链表来进行删除操作
func removeElements(head *ListNode, val int) *ListNode {
	// 1.需要首先对头结点进行处理，判断头节点是否相等
 	for head !=nil && head.Val == val {
 		head = head.Next
	}
    // 2不能直接操作头节点，头结点遍历会导致值不断变化，最后返回的就是最后一个值，这里需要用cur作为临时指针节点指向头节点；
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val{
			cur.Next = cur.Next.Next
		}else{
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


func main(){
	// 创建链表: 1 -> 2 -> 6 -> 3 -> 4 -> 5 -> 6
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6,
		Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}

	// 移除元素 6
	newHead := removeElements(head, 6)
	fmt.Printf("result: %v", newHead)
}