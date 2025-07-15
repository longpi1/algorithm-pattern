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
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

}

//// 利用切片判断
//func isPalindrome(head *ListNode) bool {
//	cur := head
//	result := make([]int, 0)
//	for cur != nil {
//		result = append(result, cur.Val)
//		cur = cur.Next
//	}
//
//	for i := 0; i < len(result)/2; i++ {
//		if result[i] != result[len(result)-1-i] {
//			return false
//		}
//	}
//	return true
//}

// 通过获取中间节点和反转链表
/*
首先找到链表的中间结点 mid：

如果链表有奇数个节点，那么找的是正中间的节点。
如果链表有偶数个节点，那么找的是正中间右边的节点。
然后反转从 mid 到链表末尾的这段。如上图，反转后得到链表 6→5→4，其头节点记作 head

最后，同时遍历 head 和 head这两个链表，直到 head链表遍历结束。每次循环判断 head.val 是否等于 head.val，若不相等，则返回 false。
如果循环中没有返回 false，说明链表是回文的，返回 true。

作者：灵茶山艾府
链接：https://leetcode.cn/problems/palindrome-linked-list/solutions/2952645/o1-kong-jian-zuo-fa-xun-zhao-zhong-jian-rv0f3/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 876. 链表的中间结点
func getMidNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 206. 反转链表
func reversal(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func isPalindrome(head *ListNode) bool {
	midNode := getMidNode(head)
	reversalNode := reversal(midNode)
	for reversalNode != nil {
		// 不是回文链表
		if head.Val != reversalNode.Val {
			return false
		}
		head = head.Next
		reversalNode = reversalNode.Next
	}
	return true
}

func main() {
	testNode := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	isPalindrome(testNode)
}
