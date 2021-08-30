package main

import "fmt"

func main() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}

	if x := 42; x == 45 {
		fmt.Println("001")
	}

	x1 := 40
	if x1 == 40 {
		fmt.Println("our  value was 40")
	} else if x1 == 41 {
		fmt.Println("our  value was 42")
	}

}
