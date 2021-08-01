package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	msg := greetings.Hello("Selva")
	fmt.Println(msg)
}
