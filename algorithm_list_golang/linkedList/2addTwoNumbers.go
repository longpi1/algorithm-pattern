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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 下述代码思路错误，已添加错误注释
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 初始化一个头节点，这很好，是哑节点（dummy node）的思路。
	// 但它的命名 'head' 容易让人误解为最终返回的头节点，
	// 实际上我们应该返回 head.Next。通常命名为 'dummy'。
	head := &ListNode{}
	cur := head
	// 'needAddOne' 用来表示进位，命名清晰，很好。
	needAddOne := false

	// --- 问题 1: 循环条件是错误的 ---
	// for l1.Next != nil && l2.Next != nil {
	// 这个条件会提前一个节点结束循环。
	// 例如，如果 l1 = [2,4,3] l2 = [5,6,4]，
	// 当 l1 指向 4, l2 指向 6 时，它们的 Next 都不是 nil，循环继续。
	// 当 l1 指向 3, l2 指向 4 时，它们的 Next 都是 nil，循环就退出了！
	// 这导致最后一个节点 (3 和 4) 的相加操作被完全跳过。
	// 正确的条件应该是 for l1 != nil || l2 != nil，只要有一个链表还有节点，就应该继续。
	for l1.Next != nil && l2.Next != nil {
		tmpVal := l1.Val + l2.Val
		if needAddOne {
			tmpVal += 1
		}

		// --- 问题 2: 判断进位的逻辑错误 ---
		// if tmpVal > 10 {
		// 当两个数相加等于 10 时（比如 5+5），也应该进位。
		// 正确的判断应该是 if tmpVal >= 10。
		if tmpVal > 10 {
			needAddOne = true
			tmpVal %= 10
		} else {
			needAddOne = false
		}

		// --- 问题 3: 对头节点/哑节点的使用不正确 ---
		// cur.Val = tmpVal
		// 第一次循环时，你修改了 `head` 节点本身的值。
		// 这意味着你的哑节点被用作了结果链表的第一个实际节点。
		// 这样做虽然不一定会导致最终结果错误，但它不是哑节点的标准用法，
		// 并且在最后返回时容易出错。标准的做法是始终操作 cur.Next。
		// 另外，在创建新节点时，应该是 cur.Next = &ListNode{Val: tmpVal}，然后移动 cur。
		// 你的写法是先给当前节点赋值，再创建一个空的 Next，这更繁琐。
		cur.Val = tmpVal
		cur.Next = &ListNode{}

		// --- 问题 4: 指针移动逻辑是正确的，但是... ---
		// 由于循环条件错误，这里的移动是不完整的。
		l1 = l1.Next
		l2 = l2.Next

		// --- 问题 5: cur 指针没有移动！---
		// 这是个非常严重的错误。你每次循环都只更新了同一个节点 `head` 的值和它的 Next。
		// cur 指针应该在每次循环后移动到新创建的节点上：cur = cur.Next。
		// 否则，你的链表永远只有一个节点。
	}

	// --- 问题 6: 处理剩余链表的逻辑完全错误 ---
	// 假设 l1 = [9,9], l2 = [1]。循环不会执行。
	// l1.Next != nil 为 true，l2.Next != nil 为 false，循环不进入。
	// if l1.Next != nil {  // 这个条件会为 true
	//	 cur.Next = l1   // cur.Next 指向了 [9,9] 这个链表
	// }
	// `cur` 仍然是 `head`，所以 `head.Next` 指向了 l1。
	// 那么 l2 的 [1] 呢？完全被忽略了。
	// 并且，最重要的进位逻辑在这里完全没有体现。如果 l1 是 [9,9,9] 而 l2 是 [1]，
	// 结果应该是 [0,0,0,1]，你的代码只会错误地拼接链表。
	// 正确的做法是，在主循环结束后，继续写循环来处理 l1 或 l2 的剩余部分，并考虑进位。
	if l1.Next != nil {
		cur.Next = l1
	}
	if l2.Next != nil {
		cur.Next = l2
	}

	// --- 问题 7: 返回值错误 ---
	// return head
	// 你返回了哑节点本身，而不是新链表的真正头节点(head.Next)。
	// 由于你也修改了 head.Val，所以你返回的链表第一个节点的值可能是对的，
	// 但这是一种不规范且有潜在风险的写法。
	return head
}

// 正确答案如下
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 使用哑节点（dummy node）作为新链表的伪头节点，这是一个好习惯。
	// 这样可以统一处理所有节点，无需对头节点做特殊判断。
	dummy := &ListNode{}
	// cur 是一个移动的指针，用于构建新链表。初始指向哑节点。
	cur := dummy
	// carry 用于存储进位值，初始化为 0。
	carry := 0

	// 循环条件：只要 l1 或 l2 还有一个不为空，或者还有进位，就继续循环。
	// 这可以优雅地处理两个链表不等长以及最后一次相加产生进位的情况。
	for l1 != nil || l2 != nil || carry > 0 {
		// 初始化当前位的和为进位值。
		sum := carry

		// 如果 l1 不为空，则将其值加入 sum，并将 l1 指针后移。
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		// 如果 l2 不为空，则将其值加入 sum，并将 l2 指针后移。
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// 计算新的进位。例如 sum=15, carry=1。
		carry = sum / 10
		// 计算当前节点应该存储的值。例如 sum=15, val=5。
		val := sum % 10

		// 创建新节点，并将其连接到新链表的末尾。
		cur.Next = &ListNode{Val: val}
		// 将 cur 指针移动到新创建的节点上，为下一次循环做准备。
		cur = cur.Next
	}

	// 哑节点的 Next 指针指向的才是新链表的真正头节点。
	return dummy.Next
}
