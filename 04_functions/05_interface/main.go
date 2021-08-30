package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("Person was passed through bar")

	case secretAgent:
		fmt.Println("SecretAgent was passed through bar")
	}
	fmt.Println("Bar")
}
func (p person) speak() {
	fmt.Println("Person speak")
	fmt.Println(p.firstName)
}

func (sa secretAgent) speak() {
	fmt.Println("SecretAgent Speak")
	fmt.Println(sa.firstName)
}

func main() {
	sa := secretAgent{
		person: person{
			firstName: "selva",
			lastName:  "Mohandoss",
		},
		ltk: true,
	}
	sa.speak()
	bar(sa)
	p1 := person{
		firstName: "Selva",
		lastName:  "Senthilnathan",
	}
	bar(p1)

}
