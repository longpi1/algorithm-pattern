package main

/*
208. 实现 Trie (前缀树)
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：
Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

示例：
输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True
*/

type Trie struct {
	children [26]*Trie // 26个字母的子节点数组
	isEnd    bool // 标志当前节点是否为单词结尾
}

// Constructor 初始化 Trie 结构
func Constructor() Trie {
	return Trie{}
}

// Insert 插入单词到 Trie 中
func (t *Trie) Insert(word string) {
	node := t
	// 遍历单词的每个字符
	for _, ch := range word {
		ch -= 'a' // 计算字符的相对位置（a 的 ASCII 值是 97，将其映射到 0-25 的范围）
		// 如果当前字符的节点不存在，创建一个新节点
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch] // 移动到下一个节点
	}
	node.isEnd = true // 标志当前节点为单词结尾
}

// SearchPrefix 搜索前缀是否存在于 Trie 中
func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	// 遍历前缀的每个字符
	for _, ch := range prefix {
		ch -= 'a' // 计算字符的相对位置（a 的 ASCII 值是 97，将其映射到 0-25 的范围）
		// 如果当前字符的节点不存在，返回 nil
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch] // 移动到下一个节点
	}
	return node // 返回最后一个字符的节点
}

// Search 搜索完整单词是否存在于 Trie 中
func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word) // 搜索前缀
	return node != nil && node.isEnd // 如果前缀存在且标志当前节点为单词结尾，返回 true
}

// StartsWith 检查 Trie 中是否有以给定前缀开头的单词
func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil // 检查前缀是否存在
}

