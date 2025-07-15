package main

/*
21. 合并两个有序链表

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	newNode := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			newNode.Next = list1
			list1 = list1.Next
		} else {
			newNode.Next = list2
			list2 = list2.Next
		}
		newNode = newNode.Next
	}
	if list1 != nil {
		newNode.Next = list1
	}
	if list2 != nil {
		newNode.Next = list2
	}
	return dummy.Next
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // 用哨兵节点简化代码逻辑
	cur := dummy         // cur 指向新链表的末尾
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1 // 把 list1 加到新链表中
			list1 = list1.Next
		} else { // 注：相等的情况加哪个节点都是可以的
			cur.Next = list2 // 把 list2 加到新链表中
			list2 = list2.Next
		}
		cur = cur.Next
	}
	// 拼接剩余链表
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}
