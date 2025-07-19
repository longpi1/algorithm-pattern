package main

/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

示例 1：

输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：

输入：head = []
输出：[]
示例 3：

输入：head = [1]
输出：[1]
*/
// 下述思路错误
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}

	cur := dummyHead.Next // !!! BUG: cur 指向的是第 1 个待交换节点，
	//          但真正需要的是“前驱节点”，
	//          这样才能把前驱 .Next 指到 newHead（第二个节点）。
	//          否则交换后无法把前驱接回去，链表断裂。

	for cur != nil && cur.Next != nil {
		next := cur.Next // 待交换的第二个节点

		cur.Next = next.Next // 把第 1 个节点连到第 2 个节点之后的链表
		next.Next = cur      // 把第 2 个节点连到第 1 个节点前面

		// !!! BUG: 少了一步把“前驱节点”的 Next 指向 next，
		//          导致原来 dummyHead（或上一对的尾巴）仍指向 cur，
		//          交换好的 next/cur 对被甩出主链。
		// dummyHead.Next = next      // 第 1 轮本应这样连回
		// pre.Next       = next      // 后续轮次应把上次交换后的尾巴连回

		cur = cur.Next // !!! BUG: cur 已经被移动到 pair 里的“第 1 个节点”，
		//          其 Next 现在指向下一个 pair 的“第 1 个节点”没错，
		//          但因为前驱丢失，上一步错误被放大：后续循环依旧断链
	}
	return dummyHead.Next
}

/*func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}

	cur := head
	//var pre *ListNode
	for cur != nil && cur.Next != nil{
		next := cur.Next.Next
		cur.Next.Next = cur

		cur.Next = next

		cur = next
	}

	return dummyHead.Next
}*/

/*
上述答案错误。存在以下几个问题：
1.存在一个逻辑错误，会导致链表断裂。在交换两个相邻节点时：cur.Next.Next = cur，cur.Next = next，您需要将当前节点 cur 的 Next 指针指向下一个节点，但是在这行代码中，将 Next 指针指向了当前节点 cur 自身，这会导致链表中断裂，丢失原来的连接关系。
2.需要利用起头节点，不然不好处理前面节点的转换
3.cur = next,这里不应该再指向之前的临时变量，应该是新的位置
*/

func swapPairs(head *ListNode) *ListNode {
	// 使用哑节点，这是一个非常好的实践，可以避免对头节点进行特殊处理。
	dummyHead := &ListNode{Next: head}

	// --- 问题 1: 'cur' 的角色不正确 ---
	// 你让 'cur' 指向了第一个要交换的节点（即 head）。
	// 在交换过程中，你需要一个指针指向“待交换对”的前一个节点，
	// 这样才能将交换后的新头节点连接回主链表。
	// 你的 'cur' 自身就是参与交换的一部分，所以它无法完成“连接”这个任务。
	// 我们需要一个像 `prev` 这样的指针，初始时指向 dummyHead。
	cur := dummyHead.Next

	// 'pre' 被注释掉了，但实际上一个类似 'pre' 的指针正是解决问题的关键。

	// 循环条件是正确的，确保至少有两个节点可供交换。
	for cur != nil && cur.Next != nil {
		// 假设链表是 dummy -> 1 -> 2 -> 3 -> 4
		// 当前 cur 是 1

		// next 指向节点 2，正确。
		next := cur.Next

		// --- 核心交换逻辑 (部分正确) ---
		// cur.Next = next.Next  (1.Next 指向了 3)
		// 此时链表片段变为: 1 -> 3
		cur.Next = next.Next

		// next.Next = cur (2.Next 指向了 1)
		// 此时我们有了交换好的一对: 2 -> 1
		next.Next = cur

		// --- 问题 2: 链条断裂！(最关键的错误) ---
		// 我们现在有两段链表：
		// 1. dummy -> 1 -> 3 -> 4
		// 2. 2 -> 1 (我们刚刚交换好的一对)
		//
		// `dummy` 节点仍然指向 `1`！我们从未更新 `dummy.Next` 让它指向交换后的新头 `2`。
		// 这个连接操作是整个算法中最关键的、缺失的一步。
		// 正确的操作应该是： `prev.Next = next` (其中 prev 是交换前指向 1 的那个节点)。

		// --- 问题 3: 'cur' 指针的更新逻辑错误 ---
		// cur = cur.Next
		// 在上面的交换后，cur (即节点 1) 的 Next 已经被我们设置为了节点 3。
		// 所以这行代码让 cur 直接跳到了节点 3。
		//
		// 这样会跳过节点 2，导致下一次循环的起始点错误。
		// 实际上，完成一对交换 (1,2 -> 2,1) 后，下一对要交换的是 (3,4)。
		// 为了准备下一次交换，我们的前置指针 (`prev`) 应该移动到节点 1 的位置。
		// 所以更新应该是 `prev = cur`。
		cur = cur.Next
	}

	// 因为 dummyHead.Next 从未被更新，它始终指向原始的 head (节点 1)。
	// 所以函数会返回 1 -> 3 -> 4 -> ... 这样的错误结果。
	// 第一次交换的成果 (节点 2) 完全丢失了。
	return dummyHead.Next
}

// 递归版本,画图调式之后更清晰，递归最主要找边界条件与非边界条件，最后得到就i是最小子问题的结果
func swapPairs(head *ListNode) *ListNode {
	// # 递归边界
	if head == nil || head.Next == nil {
		// 不足两个节点，无需交换
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy // cur 当作“前驱节点”

	for cur.Next != nil && cur.Next.Next != nil {
		first := cur.Next
		second := first.Next

		// 交换
		first.Next = second.Next
		second.Next = first
		cur.Next = second

		// cur 指向下一对的前驱
		cur = first
	}
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}

	swapPairs(head)
}
