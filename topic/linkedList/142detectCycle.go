package main

/*
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改 链表。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。

提示：

链表中节点的数目范围在范围 [0, 104] 内
-105 <= Node.val <= 105
pos 的值为 -1 或者链表中的一个有效索引
*/

type ListNode struct {
	Val int
	Next *ListNode
}

/*
解题思路
我们使用两个指针，fast\{fast}fast 与 slow\{slow}slow。它们起始都位于链表的头部。随后，slow\{slow}slow 指针每次向后移动一个位置，而 fast\{fast}fast 指针向后移动两个位置。如果链表中存在环，
则 fast\{fast}fast 指针最终将再次与 slow\{slow}slow 指针在环中相遇。
如下图所示，设链表中环外部分的长度为 aaa。slow\{slow}slow 指针进入环后，又走了 b 的距离与 fast\{fast}fast 相遇。此时，fast\{fast}fast 指针已经走完了环的 n 圈，
因此它走过的总距离为 a+n(b+c)+b=a+(n+1)b+nc

根据题意，任意时刻，fast\{fast}fast 指针走过的距离都为 slow\{slow}slow 指针的 2 倍。因此，我们有
a+(n+1)b+nc=2(a+b)  ⟹  a=c+(n−1)(b+c)a+(n+1)b+nc=2(a+b) \implies a=c+(n-1)(b+c)
a+(n+1)b+nc=2(a+b)⟹a=c+(n−1)(b+c)
有了 a=c+(n−1)(b+c)a=c+(n-1)(b+c)a=c+(n−1)(b+c) 的等量关系，我们会发现：从相遇点到入环点的距离加上 n−1 圈的环长，恰好等于从链表头部到入环点的距离。

因此，当发现 slow\{slow}slow 与 fast\{fast}fast 相遇时，我们再额外使用一个指针 ptr\{ptr}ptr。起始，它指向链表头部；随后，它和 slow\{slow}slow 每次向后移动一个位置。最终，它们会在入环点相遇。

*/
// 下述代码基于快慢指针，存在一个错误点，已在代码中更正
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return nil
	}
	// 定义快慢指针，起点一致，快指针是慢指针的一倍速；
	slow := head
	fast := head
	flag := false

	for fast != nil &&  fast.Next != nil {
		//在第一个循环中，当快慢指针相遇时（即 fast == slow），您使用 break 关键字终止了循环。
		//最开始的时候肯定相等以为都指向head
		//if fast == slow {
		//	flag = true
		//	break
		//}
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			flag = true
			break
		}
	}
	for flag {
		fast = head
		if fast == slow {
			return fast
		}
		fast = fast.Next
		slow = slow.Next
	}

	return nil
}


//2.哈希表实现方式:一个非常直观的思路是：我们遍历链表中的每个节点，并将它记录下来；一旦遇到了此前遍历过的节点，就可以判定链表中存在环。借助哈希表可以很方便地实现。
func detectCycle(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}
