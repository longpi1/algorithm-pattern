package main

/*

给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。


示例 1：


输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]
示例 2：

输入：head = [5], left = 1, right = 1
输出：[5]


提示：

链表中节点数目为 n
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n
*/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func  reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{Next: head}
	p0 := dummyHead
	//定义前后指针节点
	var pre *ListNode
	cur := p0.Next
	count := 1
	for cur != nil {
		if count < left {
			count ++
			p0 = p0.Next
			cur = cur.Next
		}else if left <= count && count <= right {
			count ++
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}else {
			break
		}
	}
	p0.Next.Next = cur
	p0.Next = pre

	return dummyHead.Next
}

func main(){
	head := &ListNode{Val:1,Next: &ListNode{Val:2,Next: &ListNode{Val:3,Next: &ListNode{Val:4,Next: &ListNode{Val: 5,Next: nil}}}}}
	left := 2
	right := 4
	reverseBetween(head,left,right)
}