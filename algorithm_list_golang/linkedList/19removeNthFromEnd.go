package main


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


func removeNthFromEnd(head *ListNode, n int) *ListNode {
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

//直接先遍历求得总长度，然后再求倒数n个的位置
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	tmp := head
	cur := dummyHead
	len := 0
	for tmp != nil {
		len ++
		tmp = tmp.Next
	}
	for i := 0; i < len - n ; i++ {
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}



func main(){
	head := &ListNode{Val:1,Next: &ListNode{Val:2,Next: &ListNode{Val:3,Next: &ListNode{Val:4,Next: nil}}}}
	n := 2
	result := removeNthFromEnd(head,n)
	println(result)
}