package main

import "fmt"

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。



示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]


提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
*/

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 第一次逻辑错误
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	fast := head
	nodeLen := 0

	// ---------- 第 1 步：fast 先走 n 步 ----------
	for i := 0; i < n; i++ {
		if fast != nil {
			fast = fast.Next
			nodeLen++
		}
		// !!! BUG: 如果 n 大于链表长度，fast 很快变 nil，
		//          但循环仍继续执行，最终 nodeLen < n，
		//          后面逻辑全部失效却没有任何保护或报错。
	}

	// ---------- 第 2 步：fast 和 head 同时走到链尾 ----------
	// !!! BUG: 把 head 当作“慢指针”来移动，导致变量 head 不再指向链表头，
	//          语义混乱；后面删除节点时需要的是“慢指针的前驱”，
	//          此处却直接用 head，位置不对。
	for fast != nil {
		head = head.Next // !!! head 被篡改
		fast = fast.Next
		nodeLen++
	}

	// ---------- 第 3 步：特殊情况——删除头结点 ----------
	if nodeLen == n { // 链表长度等于 n
		return dummy.Next.Next // 删除的是原头节点
		// !!! BUG: 当 n > 长度 时不会进入此分支，却仍继续向下执行，必错。
	}

	// ---------- 第 4 步：常规删除 ----------
	head.Next = head.Next.Next // !!! BUG: head 目前指向待删除节点本身，
	//          而不是它的前驱，最终要么 panic，
	//          要么删除了错误节点。
	return dummy.Next // !!! BUG: 上面误删节点后，dummy.Next
	//          不一定指向真正的链表头
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 1. 使用哑节点，指向原始头节点。这可以轻松处理删除头节点的情况。
	dummy := &ListNode{Next: head}

	// 2. 初始化快慢指针都指向哑节点。
	// fast 指针用于向前探索，slow 指针最终将停在“待删除节点的前一个节点”。
	fast := dummy
	slow := dummy

	// 3. 让 fast 指针先向前移动 n+1 步。
	// 为什么是 n+1？因为我们想让 slow 最终停在待删除节点的前一个节点。
	// 当 fast 到达末尾(nil)时，slow 和 fast 之间就差了 n+1 个节点，
	// 这意味着 slow 正好在倒数第 n+1 个位置，即倒数第 n 个节点的前一个。
	for i := 0; i <= n; i++ {
		// 如果在移动过程中 fast 变为 nil，说明 n 大于链表长度，这是无效输入。
		// （根据 LeetCode 题目保证，通常不会发生这种情况）
		if fast == nil {
			return nil // 或者根据题目要求处理
		}
		fast = fast.Next
	}

	// 4. 同时移动 fast 和 slow 指针，直到 fast 到达链表末尾 (nil)。
	// 由于它们之间保持着固定的 n+1 步距离，当 fast 到达终点时，
	// slow 就自然到达了目标位置。
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 5. 此时，slow 指向的是“待删除节点的前一个节点”。
	// 我们通过跳过 slow.Next 来删除目标节点。
	slow.Next = slow.Next.Next

	// 6. 返回哑节点的 Next，即新链表的头节点。
	return dummy.Next
}

/*func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, quick := head,head
	for i := 0; i < n; i ++ {
		slow = slow.Next
	}
	for ;slow != nil;		slow = slow.Next  {

		quick = quick.Next
	}
	quick.Next = quick.Next.Next
	return head
}*/
//
/*
上述思路错误
但是在处理特殊情况（即删除头节点）时，会出现空指针异常。
当 n 等于链表的长度时，slow 移动到了最后一个节点，而 quick 移动到了倒数第 n+1 个节点，
如果要删除的是头节点，quick.Next 将会是空指针。为了解决这个问题，
需要为链表头节点添加一个虚拟节点，并在返回时返回虚拟节点的下一个节点。
另外，为了确保链表的完整性，还需要添加对 slow.Next 是否为空的检查，以避免空指针异常。
*/

func removeNthFromEnd0(head *ListNode, n int) *ListNode {
	// 添加虚拟头节点
	dummy := &ListNode{0, head}
	slow, quick := dummy, dummy

	// 将 slow 移动到 n+1 个节点处
	for i := 0; i <= n; i++ {
		quick = quick.Next
	}

	// 同时移动 slow 和 quick，直到 quick 到达链表末尾
	for quick != nil {
		slow = slow.Next
		quick = quick.Next
	}

	// 删除倒数第 n 个节点
	slow.Next = slow.Next.Next

	// 返回虚拟头节点的下一个节点，保持链表完整性
	return dummy.Next
}

// 直接先遍历求得总长度，然后再求倒数n个的位置
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	tmp := head
	cur := dummyHead
	len := 0
	for tmp != nil {
		len++
		tmp = tmp.Next
	}
	for i := 0; i < len-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummyHead.Next
}

/*
三、双指针，快慢指针
思路与算法
我们也可以在不预处理出链表的长度，以及使用常数空间的前提下解决本题。
由于我们需要找到倒数第 nnn 个节点，因此我们可以使用两个指针 first 和 second 同时对链表进行遍历，并且 first 比 second 超前 nnn 个节点。当 first 遍历到链表的末尾时，second 就恰好处于倒数第 nnn 个节点。
具体地，初始时 first 和 second均指向头节点。我们首先使用 first 对链表进行遍历，遍历的次数为 nnn。此时，first 和 second之间间隔了 n−1n-1n−1 个节点，即 first 比 second 超前了 nnn 个节点。

在这之后，我们同时使用 first 和 second对链表进行遍历。当 first 遍历到链表的末尾（即 first 为空指针）时，second 恰好指向倒数第 nnn 个节点。
根据方法一和方法二，如果我们能够得到的是倒数第 nnn 个节点的前驱节点而不是倒数第 nnn 个节点的话，删除操作会更加方便。因此我们可以考虑在初始时将 second 指向哑节点，其余的操作步骤不变。这样一来，当 first 遍历到链表的末尾时，second 的下一个节点就是我们需要删除的节点。
*/
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

/*
 */
// 栈
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	n := 2
	result := removeNthFromEnd(head, n)
	for result != nil {
		fmt.Printf("result: %v", result.Val)
		result = result.Next

	}

}
