package main

import (
	"fmt"
	"runtime"
)

var x1 int
var y1 rune
var z1 byte

func main() {
	x := 42
	y := 45.55
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	x1 = 454
	fmt.Println(x1)
	y1 = 343
	fmt.Println(y1)
	z1 = 0
	fmt.Println(z1)
	fmt.Println(runtime.GOMAXPROCS)
	fmt.Println(runtime.GOOS)
}
