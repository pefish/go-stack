package main

import (
	"fmt"
	"github.com/pefish/go-stack"
)

func test() {
	test1()
}

func test1() {
	test2()
}

func test2() {
	fmt.Println(go_stack.Stack.GetStack(go_stack.Option{
		Skip: 0,
		FilenameMustInclude: "Work",
	}))
}

func main() {
	fmt.Println(`haha`)
	test()
}





