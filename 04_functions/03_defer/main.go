package main

import "fmt"

func main() {
	fmt.Println("defer")
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
