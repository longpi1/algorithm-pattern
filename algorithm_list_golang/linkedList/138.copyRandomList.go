package main

/*
138. 随机链表的复制
给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。
新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。
复制链表中的指针都不应指向原链表中的节点 。
例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。
返回复制链表的头节点。
用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：
val：一个表示 Node.val 的整数。
random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
你的代码 只 接受原链表的头节点 head 作为传入参数。


示例 1：
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
示例 2：
输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]
示例 3：
输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]
*/

/*func copyRandomList(head *Node) *Node {
	// memo用于存储原节点和复制节点的映射关系
	memo := map[*Node]*Node{}

	// 第一次遍历，创建复制节点并存入memo
	for tmp := head; tmp != nil; tmp = tmp.Next {
		memo[tmp] = &Node{tmp.Val, tmp.Next, tmp.Random} // 复制当前节点，并初始化Random和Next指针为nil
	}

	//// 第二次遍历，连接复制节点的Random和Next指针
	//for node, cpy := range memo {
	//	if node.Random != nil {
	//		cpy.Random = memo[node.Random] // 连接复制节点的Random指针
	//	}
	//	if node.Next != nil {
	//		cpy.Next = memo[node.Next] // 连接复制节点的Next指针
	//	}
	//}

	// 返回复制链表的头节点
	return memo[head]
}*/
/*
上述思路错误：
提供的代码中，在第一次遍历时，直接将新节点的Next和Random指针指向了原链表中对应节点的Next和Random指针。
这样做的问题是，新链表中的节点的Next指针和Random指针依然指向了原链表的节点，而不是新链表中的节点。在第二次遍历中，由于新链表的节点仍然指向原链表的节点，连接操作就无法正确地建立新链表内部的指针关系。
正确的做法是，在第一次遍历中，只创建新节点，并将新节点存储在memo中，
不设置新节点的Next和Random指针。在第二次遍历中，再通过memo中的映射关系，建立新链表的Next和Random指针。
*/

type Node struct {
     Val int
     Next *Node
     Random *Node
 }

/*
1.创建节点映射关系： 第一次遍历原链表，创建原节点和复制节点的映射关系，存储在memo中。复制节点的Random和Next指针暂时初始化为nil。
2.连接Random和Next指针： 第二次遍历原链表，通过memo中的映射关系，连接复制节点的Random和Next指针。
3.返回结果： 返回复制链表的头节点，即memo[head]。
*/
// copyRandomList 复制带随机指针的链表

func copyRandomList(head *Node) *Node {
	memo := map[*Node]*Node{} // 用于存储原节点和复制节点的映射关系

	// 第一次遍历，创建新节点并存入memo
	for tmp := head; tmp != nil; tmp = tmp.Next {
		memo[tmp] = &Node{tmp.Val, nil, nil} // 只创建新节点，不设置Next和Random指针
	}

	// 第二次遍历，连接新节点的Next和Random指针
	for node, cpy := range memo {
		cpy.Next = memo[node.Next]   // 连接新节点的Next指针
		cpy.Random = memo[node.Random] // 连接新节点的Random指针
	}

	// 返回新链表的头节点
	return memo[head]
}
