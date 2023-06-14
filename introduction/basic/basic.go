package basic

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
	sort.Float64s([]float64{})
	sort.Ints([]int{})
	sort.Strings([]string{})
	data := make([]int, 10)
	sort.Slice(data, func(i, j int) bool {
		return  data[i] > data[j]
	})

}

