package main

import "fmt"

const (
	a11        = 454
	b11 string = "b11"
)

const (
	x = 2017 + iota
	y = 2017 + iota
	z = 2017 + iota
)

func main() {
	a := 42
	fmt.Printf("%d\t%b\t#x", a, a, a)
	b := (42 >= 45)
	fmt.Println(b)
	fmt.Println(a11)
	fmt.Println(b11)
	fmt.Println(x)
	fmt.Println(y)

}
