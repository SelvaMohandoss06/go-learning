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

func (s secretAgent) speak() {
	fmt.Println(s.firstName)
	fmt.Println(s.lastName)
}

func main() {
	sa := secretAgent{
		person: person{
			firstName: "Selva",
			lastName:  "Mohandoss",
		},
		ltk: true,
	}
	sa.speak()
	fmt.Println(sa.firstName)

}
