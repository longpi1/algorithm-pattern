package main



/*

给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。


示例 1：


输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
示例 2：


输入：head = [1,2]
输出：[2,1]
示例 3：

输入：head = []
输出：[]


提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
在遍历链表时，将当前节点的 next  指针改为指向前一个节点。由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。在更改引用之前，还需要存储后一个节点。最后返回新的头引用。

*/

/*func reverseList(head *ListNode) *ListNode {
	cur := head
	pre := &ListNode{}

	for cur.Next != nil {
		next :=	cur.Next
		pre = cur
		cur.Next = pre
		cur = next
	}
	return head
}*/
/*
上述代码存在以下几个问题：
1.不应该是cur.Next != nil，应该为cur ！= nil
2.不应该返回head,最后返回新的头引用pre
3.
4.不应该初始pre为&ListNode{},应该直接var定义变量即可

*/

//迭代+双指针
func reverseList(head *ListNode) *ListNode {
	cur := head
	// 不应该初始pre为&ListNode{},应该直接var定义变量即可
	//pre := &ListNode{}
	var pre *ListNode
	for cur != nil {
		//for cur.Next != nil {
		next :=	cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
	// 不应该返回head,最后返回新的头引用pre
	//return head
}


// 递归（判断是否有最小子问题）递归最主要找边界条件与非边界条件，最后得到就i是最小子问题的结果
func reverseList(head *ListNode) *ListNode {
	// 1. 递归终止条件也就是边界条件
	if head == nil || head.Next == nil {
		return head
	}
	var p = reverseList(head.Next)
	//反转
	head.Next.Next = head
	head.Next = nil
	return p
}
