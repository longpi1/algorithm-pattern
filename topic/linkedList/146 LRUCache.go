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

/*
采用了双向链表和哈希表的结合，使用双向链表维护缓存项的访问顺序，使用哈希表实现O(1)时间复杂度的查找和删除操作。当插入新缓存项时，
如果缓存容量超过限制，则删除最近最少访问的节点。这样，LRU（Least Recently Used）策略保证了最近访问的缓存项会被保留在缓存中，
而很久没有访问的缓存项会被移除。
*/
// entry 结构体用于表示缓存中的键值对
type entry struct {
	key, value int
}

// LRUCache 结构体表示LRU缓存
type LRUCache struct {
	capacity  int              // 缓存容量
	list      *list.List       // 双向链表，用于按访问顺序存储缓存项
	keyToNode map[int]*list.Element // 哈希表，用于存储键到双向链表节点的映射关系
}

// Constructor 初始化LRUCache
func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, list.New(), map[int]*list.Element{}}
}

// Get 根据键获取缓存值，如果键不存在返回-1，否则返回对应的值，并将该键值对移到链表头部表示最近访问
func (c *LRUCache) Get(key int) int {
	node := c.keyToNode[key]
	if node == nil {
		return -1
	}
	c.list.MoveToFront(node) // 将节点移到链表头部，表示最近访问
	return node.Value.(entry).value
}

// Put 将键值对放入缓存，如果键已存在，则更新值并将该键值对移到链表头部表示最近访问；
// 如果键不存在，创建新节点，放入链表头部，如果缓存容量超过限制，则删除最近最少访问的节点（链表尾部）。
func (c *LRUCache) Put(key, value int) {
	if node := c.keyToNode[key]; node != nil { // 如果键已存在
		node.Value = entry{key, value} // 更新值
		c.list.MoveToFront(node)       // 将节点移到链表头部，表示最近访问
		return
	}
	c.keyToNode[key] = c.list.PushFront(entry{key, value}) // 创建新节点，放入链表头部
	if len(c.keyToNode) > c.capacity {                      // 如果缓存容量超过限制
		delete(c.keyToNode, c.list.Remove(c.list.Back()).(entry).key) // 删除最近最少访问的节点（链表尾部）
	}
}
