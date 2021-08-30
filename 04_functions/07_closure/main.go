package main

import "fmt"

var i int

func main() {
	var i int
	fmt.Println("Closure")
	fmt.Println(i)
	i++
	fmt.Println(i)
	foo()
	fmt.Println(i)
	{
		y := 560
		fmt.Println(y)
	}
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(b())

}
func foo() {
	i++
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
