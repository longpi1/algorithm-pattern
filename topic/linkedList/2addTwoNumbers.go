package main

/*
2. 两数相加
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
*/


func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode // 用于追踪结果链表的尾部
	carry := 0 // 用于记录进位

	// 遍历 l1 和 l2，直到两个链表都为空
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0 // 用于存储当前节点的值
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		// 计算当前节点的和以及进位
		sum := n1 + n2 + carry
		sum, carry = sum % 10, sum / 10

		// 创建新节点，并根据头部是否为空初始化头部或者追加到尾部
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	// 如果最高位有进位，则需要在结果链表末尾添加一个新节点
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return head // 返回结果链表的头部
}
