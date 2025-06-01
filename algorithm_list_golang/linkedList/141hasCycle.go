package main

/*
给你一个链表的头节点 head ，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
如果链表中存在环 ，则返回 true 。 否则，返回 false 。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。

提示：
链表中节点的数目范围是 [0, 104]
-105 <= Node.val <= 105
pos 为 -1 或者链表中的一个 有效索引 。

*/

/*

 */

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		if slow == fast {
			// 错误1：第一次循环时，slow 和 fast 都指向 head，必然相等，会错误地返回 true
			// 修正建议：在比较之前，应该先移动指针
			return true
		}
		// 错误2：这里使用了 head.Next 和 head.Next.Next，而不是 slow 和 fast
		// 这会导致指针始终基于 head 移动，而不是基于当前 slow 和 fast 的位置
		// 修正建议：应改为 slow = slow.Next 和 fast = fast.Next.Next
		slow = head.Next
		fast = head.Next.Next
	}
	return false
}
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		// 先移动指针，再进行比较
		slow = slow.Next
		fast = fast.Next.Next
		// 在移动后再比较，避免第一次循环时 slow == fast 的误判
		if slow == fast {
			return true
		}
	}
	return false
}

// 双指针，（快慢指针的思路）
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	result := hasCycle(head)
	println(result)
}
