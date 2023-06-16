package main

import "sort"

// go 通过切片模拟栈和队列
func stack(){
	stack := make([]string, 10)
	// 入栈
	stack = append(stack, "test")
	// 出栈
	value :=stack[len(stack)-1]
	print(value)
	stack = stack[0:len(stack)-1]

	stack=stack[:len(stack)-1]
}

func queue()  {
	queue := make([]string, 10)
	// enqueue
	queue = append(queue, "test")
	// dequeue
	value := queue[0]
	print(value)
	queue = queue[1:]
}

// 字典
func dict(){
	dict := make(map[string]int, 0)
	//set
	dict["test"] = 0
	//get
	value := dict["test"]
	print(value)
	// 删除k
	delete(dict,"hello")
}


// 标准库
func libary(){
	// sort
	sort.Float64s([]float64{})
	sort.Ints([]int{})
	sort.Strings([]string{})
	data := make([]int, 10)
	sort.Slice(data, func(i, j int) bool {
		return  data[i] > data[j]
	})

}




//> 给定一个  haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。如果不存在，则返回  -1。
func findStr(srcStr string, dstStr string) int {
	if len(dstStr) == 0 {
		return 0
	}
	var i, j int
	for i = 0; i < len(srcStr) - len(dstStr)+1; i++ {
		for j = 0; j <len(dstStr) ; j++ {
			if srcStr[i+j] != dstStr[j] {
				break
			}
		}
		if len(dstStr) == j{
			return i
		}
	}

	return -1
}


// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
/*
回溯模板
result = []
func backtrack(选择列表,路径):
    if 满足结束条件:
        result.add(路径)
        return
    for 选择 in 选择列表:
        做选择
        backtrack(选择列表,路径)
        撤销选择
```
*/
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	// 保存中间结果
	list := make([]int, 0)
	backtrack(nums, 0, list, &result)
	return result
}

func backtrack(nums []int, pos int, list []int, result *[][]int){
	// 把临时结果复制出来保存到最终结果
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	for i := pos; i <len(nums) ; i++ {
		list = append(list, nums[i])
		backtrack(nums,i+1,list,result)
		list = list[0 : len(list)-1]
	}
}




func main(){
	nums := []int{1,2,3}
	print(subsets(nums))
}
