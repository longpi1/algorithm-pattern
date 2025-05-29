package main

/*
160. 相交链表
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
图示两个链表在节点 c1 开始相交：

题目数据 保证 整个链式结构中不存在环。

注意，函数返回结果后，链表必须 保持其原始结构 。

自定义评测：

评测系统 的输入如下（你设计的程序 不适用 此输入）：

intersectVal - 相交的起始节点的值。如果不存在相交节点，这一值为 0
listA - 第一个链表
listB - 第二个链表
skipA - 在 listA 中（从头节点开始）跳到交叉节点的节点数
skipB - 在 listB 中（从头节点开始）跳到交叉节点的节点数
评测系统将根据这些输入创建链式数据结构，并将两个头节点 headA 和 headB 传递给你的程序。如果程序能够正确返回相交节点，那么你的解决方案将被 视作正确答案 。



示例 1：
输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
输出：Intersected at '8'
解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,6,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
— 请注意相交节点的值不为 1，因为在链表 A 和链表 B 之中值为 1 的节点 (A 中第二个节点和 B 中第三个节点) 是不同的节点。换句话说，它们在内存中指向两个不同的位置，而链表 A 和链表 B 中值为 8 的节点 (A 中第三个节点，B 中第四个节点) 在内存中指向相同的位置。


示例 2：
输入：intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Intersected at '2'
解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [1,9,1,2,4]，链表 B 为 [3,2,4]。
在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	/*
		// 第一次做错误，应该改为直接用对应节点为map的key
		m := make(map[int]bool)
		for headA != nil {
			m[headA.Val] = true
			headA = headA.Next
		}

		for headB != nil {
			if _, ok := m[headB.Val]; ok {
				return headB
			}
			headB = headB.Next
		}*/
	m := make(map[*ListNode]bool)
	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for headA.Next != nil {
		m[headA] = true
		headA = headA.Next
	}
	for headB.Next != nil {
		if m[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

/*
双指针：

*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 如果任一链表为空，无交点，返回 nil
	if headA == nil || headB == nil {
		return nil
	}

	// 创建两个指针 pa 和 pb 分别指向链表 headA 和 headB 的头节点
	pa, pb := headA, headB

	// 遍历两个链表，直到找到交点或者遍历完所有节点
	for pa != pb {
		// 如果 pa 到达链表 A 的末尾，则重定位到链表 B 的头部
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		// 如果 pb 到达链表 B 的末尾，则重定位到链表 A 的头部
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	// 返回交点或者 nil（如果没有交点）
	return pa
}
