package main

import "fmt"

type test struct {
	str string
}

func main() {
	ss := test{
		str: "test",
	}
	println(string(ss))
	s := "longping"
	for i := 0; i <len(s) ; i++ {
		fmt.Println("result: ", s[i])
	}

}
