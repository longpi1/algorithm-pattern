package main

import "sort"

// go 通过切片模拟栈和队列
func stack() {
	stack := make([]string, 10)
	// 入栈
	stack = append(stack, "test")
	// 出栈
	value := stack[len(stack)-1]
	print(value)
	stack = stack[0 : len(stack)-1]

	stack = stack[:len(stack)-1]
}

func queue() {
	queue := make([]string, 10)
	// enqueue
	queue = append(queue, "test")
	// dequeue
	value := queue[0]
	print(value)
	queue = queue[1:]
}

// 字典
func dict() {
	dict := make(map[string]int, 0)
	//set
	dict["test"] = 0
	//get
	value := dict["test"]
	print(value)
	// 删除k
	delete(dict, "hello")
}

// 标准库
func libary() {
	// sort
	sort.Float64s([]float64{})
	sort.Ints([]int{})
	sort.Strings([]string{})
	data := make([]int, 10)
	sort.Slice(data, func(i, j int) bool {
		return data[i] > data[j]
	})

}

// > 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。如果不存在，则返回 -1。
func findStr(srcStr string, dstStr string) int {
	if len(dstStr) == 0 {
		return 0
	}
	var i, j int
	for i = 0; i < len(srcStr)-len(dstStr)+1; i++ {
		for j = 0; j < len(dstStr); j++ {
			if srcStr[i+j] != dstStr[j] {
				break
			}
		}
		if len(dstStr) == j {
			return i
		}
	}

	return -1
}

// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
/*
// 大佬的题解思路：https://leetcode.cn/problems/subsets/solution/shou-hua-tu-jie-zi-ji-hui-su-fa-xiang-jie-wei-yun-/
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

func backtrack(nums []int, pos int, list []int, result *[][]int) {
	// 把临时结果复制出来保存到最终结果
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}

/*
用 for 枚举出当前可选的数，比如选第一个数时：1、2、3 可选。
如果第一个数选 1，选第二个数，2、3 可选；
如果第一个数选 2，选第二个数，只有 3 可选（不能选1，产生重复组合）
如果第一个数选 3，没有第二个数可选
即，每次传入子递归的 index 是：当前你选的数的索引 + 1。
每次递归枚举的选项变少，一直递归到没有可选的数字，那就进入不了for循环，落入不了递归，整个DFS结束。
可见我们没有显式地设置递归的出口，而是通过控制循环的起点，使得最后递归自然结束。

作者：xiao_ben_zhu
链接：https://leetcode.cn/problems/subsets/solution/shou-hua-tu-jie-zi-ji-hui-su-fa-xiang-jie-wei-yun-/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func subsets(nums []int) [][]int {
	res := [][]int{}

	var dfs func(i int, list []int)
	dfs = func(i int, list []int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)
		for j := i; j < len(nums); j++ {
			list = append(list, nums[j])
			dfs(j+1, list)
			list = list[:len(list)-1]
		}
	}

	dfs(0, []int{})
	return res
}

/*
单看每个元素，都有两种选择：选入子集，或不选入子集。

比如[1,2,3]，先看1，选1或不选1，都会再看2，选2或不选2，以此类推。

考察当前枚举的数，基于选它而继续，是一个递归分支；基于不选它而继续，又是一个分支。


用索引index代表当前递归考察的数nums[index]。

当index越界时，说明所有数字考察完了，得到一个解，把它加入解集，结束当前递归分支。

为什么要回溯？
因为不是找到一个子集就完事。
找到一个子集，结束递归，要撤销当前的选择，回到选择前的状态，做另一个选择——不选当前的数，基于不选，往下递归，继续生成子集。
回退到上一步，才能在包含解的空间树中把路走全，回溯出所有的解。

*/

func subsets(nums []int) [][]int {
	res := [][]int{}

	var dfs func(i int, list []int)
	dfs = func(i int, list []int) {
		if i == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}
		list = append(list, nums[i])
		dfs(i+1, list)
		list = list[:len(list)-1]
		dfs(i+1, list)
	}
	dfs(0, []int{})

	return res
}

func main() {
	nums := []int{1, 2, 3}
	print(subsets(nums))
}
