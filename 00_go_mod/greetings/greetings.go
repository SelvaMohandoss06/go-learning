package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v.Welcome", name)
	return message
}

func Hello2() {
	fmt.Println("rer")
}
