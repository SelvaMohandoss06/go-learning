package main

import "fmt"

func main() {
	s := "Hello, 世界"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("---%x\n", "世")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}

	x1 := "Hello, Playground"
	fmt.Println(x1)
	x2 := "Hello,Hawaii"
	fmt.Println(x2)
	bs := []byte(s)
	fmt.Println(bs)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}
	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

}
