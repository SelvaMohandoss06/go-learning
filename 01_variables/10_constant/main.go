package main

import (
	"fmt"
)

const (
	a int     = 42
	b float64 = 42.78
	c string  = "James Bond"
)

const typedHello string = "selva"
const untyped = "dsd"

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", b)
	fmt.Println(typedHello)
	fmt.Println(untyped)

}
