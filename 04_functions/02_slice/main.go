package main

import "fmt"

func main() {
	x := []int{1, 23, 2323}
	output := sum(x...)
	fmt.Println(output)

}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
