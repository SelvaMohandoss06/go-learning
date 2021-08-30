package main

import "fmt"

func main() {

	x := sum(1, 2, 3, 33, 3)
	fmt.Println(x)

}
func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
