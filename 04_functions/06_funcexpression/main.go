package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("func demo")
	}
	f()
	func() {
		fmt.Println("self executing func")
	}()

}
