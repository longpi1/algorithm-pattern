package main

import "container/list"

/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。



示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4
*/


type entry struct {
	key, value int
}

type LRUCache struct {
	capacity  int
	list      *list.List // 双向链表
	keyToNode map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, list.New(), map[int]*list.Element{}}
}

func (c *LRUCache) Get(key int) int {
	node := c.keyToNode[key]
	if node == nil { // 没有这本书
		return -1
	}
	c.list.MoveToFront(node) // 把这本书放在最上面
	return node.Value.(entry).value
}

func (c *LRUCache) Put(key, value int) {
	if node := c.keyToNode[key]; node != nil { // 有这本书
		node.Value = entry{key, value} // 更新
		c.list.MoveToFront(node) // 把这本书放在最上面
		return
	}
	c.keyToNode[key] = c.list.PushFront(entry{key, value}) // 新书，放在最上面
	if len(c.keyToNode) > c.capacity { // 书太多了
		delete(c.keyToNode, c.list.Remove(c.list.Back()).(entry).key) // 去掉最后一本书
	}
}
