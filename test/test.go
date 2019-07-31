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
	fmt.Println(go_stack.Stack.GetStack(1))
}

func main() {
	fmt.Println(`haha`)
	test()
}





