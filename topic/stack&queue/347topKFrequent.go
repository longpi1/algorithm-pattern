package main

import (
	"container/heap"
	"sort"
)
/*
给定一个整数数组 nums 和一个整数 k ，请返回其中出现频率前 k 高的元素。可以按 任意顺序 返回答案。

示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]

示例 2:
输入: nums = [1], k = 1
输出: [1]

提示：
1 <= nums.length <= 105
k 的取值范围是 [1, 数组中不相同的元素的个数]
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
*/

/*
func topKFrequent(nums []int, k int) []int {
	// 创建一个切片来存储 map 的键值对
	var pairs []struct {
		Key   int
		Value int
	}
	m := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		pairs += 1
	}
	//下述对map排序方式错误
	//sort.Slice(m, func(i, j int) bool {
	//	return  m[i] > m[j]
	//})
	//正确方式 使用自定义排序函数对切片按值排序
	sort.Slice(m, func(i, j int) bool {
		return m[i].Value < m[j].Value
	})
	result := make([]int,0,k)
	count := 0
	for _, vaule := range m {
		result = append(result, vaule)
		count ++
		if count == k {
			break
		}
	}
	return result
}*/

/*
type MinHeap struct {
	Heap []int
}

func (heap *MinHeap) Len() int{
	return len(heap.Heap)
}

func (heap *MinHeap) Less(i int,j int) bool{
	return heap.Heap[i] < heap.Heap[j]
}

func (heap *MinHeap) Swap(i int,j int){
	heap.Heap[i] ,heap.Heap[j] = heap.Heap[j],	heap.Heap[i]
}

func  (heap *MinHeap) Pop() interface{}{

	val := heap.Heap[0]
	heap.Heap = heap.Heap[1:len(heap.Heap)]
	return val
}

func  (heap *MinHeap) Push(val interface{}){
	// 类型断言（Type Assertion）是一个使用在接口值上的操作，用于检查接口类型变量所持有的值是否实现了期望的接口或者具体的类型。
	// 示例如下：value, ok := x.(T)
	heap.Heap = append(heap.Heap, val.(int))
}


func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	m := make(map[int]int)
	result := make([]int,0,k)
	for i := 0; i < n; i++ {
		m[nums[i]] += 1
	}
	minHeap := MinHeap{Heap: make([]int, 0)}
	heap.Init(&minHeap)

	for key, value := range m {
		minHeap.Push(value)
	}

	return result
}
*/

// 上述思路错误

/*
解题思路：
优先级队列，大顶堆/小顶堆 利用"container/heap"包实现
*/

//方法一：小顶堆 重点！！ [][2]int 定义了一个自定义类型 IHeap，它是一个切片（slice），其中每个元素都是一个包含两个整数的数组 [2]int。这个自定义类型被用作堆（优先队列）的数据结构，用于存储元素和它们的频率。
func topKFrequent(nums []int, k int) []int {
	map_num:=map[int]int{}
	//记录每个元素出现的次数
	for _,item:=range nums{
		map_num[item]++
	}
	h:=&IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key,value:=range map_num{
		heap.Push(h,[2]int{key,value})
		if h.Len()>k{
			heap.Pop(h)
		}
	}
	res:=make([]int,k)
	//按顺序返回堆中的元素
	for i:=0;i<k;i++{
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

// ！！！构建小顶堆  [][2]int 定义了一个自定义类型 IHeap，它是一个切片（slice），其中每个元素都是一个包含两个整数的数组 [2]int。这个自定义类型被用作堆（优先队列）的数据结构，用于存储元素和它们的频率。
type IHeap [][2]int

func (h IHeap) Len()int {
	return len(h)
}

func (h IHeap) Less (i,j int) bool {
	return h[i][1]<h[j][1]
}

func (h IHeap) Swap(i,j int) {
	h[i],h[j]=h[j],h[i]
}

func (h *IHeap) Push(x interface{}){
	*h=append(*h,x.([2]int))
}


func (h *IHeap) Pop() interface{}{
	n := len(*h)
	x := (*h)[n-1] // 获取堆顶元素
	*h = (*h)[:n-1] // 移除堆顶元素
	return x
}


//方法二:利用O(logn)排序
func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[nums[i]] += 1
	}
	result := make([]int,0)

	for key, _ := range m {
		result = append(result, key)
	}
	//核心思想：排序
	//可以不用包函数，自己实现快排,这里sort.Slice默认实现了快排
	sort.Slice(result, func(i, j int) bool {
		return m[result[i]] > m[result[j]]
	})

	return result[:k]
}








/*
以下是代码的详细解析：
1.map_num 是一个映射，用于存储输入 nums 中每个唯一元素作为键以及它们对应的频率作为值。
2.创建了一个自定义整数堆 IHeap 的实例，并使用 heap.Init 函数对其进行初始化。
3.代码遍历 map_num 中的元素。对于每个键值对（元素和其频率），它将该对推入堆中。如果堆的大小超过了 "k"，则会从堆中弹出最小的元素。这确保堆中始终包含前 "k" 个最频繁的元素。
4.在处理完所有 map_num 中的元素后，初始化了用于存储结果的 res 切片。
5.代码然后以逆序遍历堆，并将前 "k" 个最频繁的元素填充到 res 切片中。它从堆中弹出元素，并以逆序的方式将它们存储在 res 切片中。
6.最后，函数返回 res 切片，其中包含按频率降序排列的前 "k" 个最频繁元素。
*/
