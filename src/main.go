package main

import "fmt"

type test struct {
	str string
}

func main() {
	//ss := test{
	//	str: "test",
	//}

	s := "longping"
	for i := 0; i <len(s) ; i++ {
		fmt.Println("result: ", s[i])
	}

	println(s[1:])
	/*
	result:  108
	result:  111
	result:  110
	result:  103
	result:  112
	result:  105
	result:  110
	result:  103
	ongping
	*/
	/*
	在循环中，你使用 s[i] 来访问字符串中的字符，其中 i 表示字符在字符串中的索引。这将逐个访问字符串中的字符，并且每个字符都被视为一个独立的字符。因此，fmt.Println("result: ", s[i])
	打印的是字符而不是字符串。这是因为 s[i] 表示字符串 s 中索引为 i 的字符。

	在 println(s[1:]) 中，你使用了切片操作 s[1:]，它表示从字符串 s 中的索引1开始，截取到字符串的末尾。这将返回一个新的字符串，因此 println(s[1:]) 打印的是一个新的字符串，而不是单个字符。
	*/
}
