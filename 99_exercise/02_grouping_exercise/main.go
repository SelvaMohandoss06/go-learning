package main

import "fmt"

func main() {
	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4

	fmt.Println(a)

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[1:3])

}
