package main

import "fmt"

func main() {
	fmt.Println("")
	n := factorial(4)
	fmt.Println(n)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
