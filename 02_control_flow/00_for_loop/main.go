package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	for k := 0; k < 5; k++ {
		fmt.Println(k)
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i)
			fmt.Printf("The outer loop: %d The inner loop:%d\n", i, j)
		}
	}

	//do while loop ( 3 ways)
	c := 3
	for c < 10 {
		fmt.Println(c)
		c++
	}

	d := 1
	for {
		if d > 9 {
			break
		}
		fmt.Println(d)
		d++
	}

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x \t%#U\n", i, i, i)
	}

}
