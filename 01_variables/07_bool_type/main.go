package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	a := 7
	b := 8
	fmt.Println(a == b)

}
