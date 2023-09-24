package main

/*
234. 回文链表
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

示例 1：
输入：head = [1,2,2,1]
输出：true
示例 2：
输入：head = [1,2]
输出：false
*/



func isPalindrome(head *ListNode) bool {
	cur := head
	result := make([]int,0)
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}

	for i := 0; i < len(result) / 2; i++ {
		if 	result[i] != result[len(result)-1-i]{
			return false
		}
	}
	return true
}