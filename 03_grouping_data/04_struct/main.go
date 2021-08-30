package main

import "fmt"

type person struct {
	first string
	last  string
	email string
}

type secretagent struct {
	person
	ltk bool
	//first string
}

func main() {
	//Need to check with Kalai
	p1 := person{"Selva", "Mohandoss", "selvamohandoss@gmail.com"}
	p2 := person{"Janesh", "Senthilnathan", "janeshsenthilnathan@gmail.com"}
	fmt.Println(p1)
	fmt.Println(p2)

	sa1 := secretagent{
		person: person{
			first: "Janesh",
			last:  "Senthilnathan",
			email: "JaneshSenthilnathan@gmail.com",
		},
		ltk: true,
		//first: "Senthil",
	}

	fmt.Println("test")
	fmt.Println(sa1.first)
	//fmt.Println(sa1.first)
	p3 := person{
		first: "Overide",
		last:  "test",
		email: "fsdfdsf@gmail.com",
	}
	fmt.Println(sa1)
	fmt.Println(p3)
	fmt.Println(sa1.person.first)
	fmt.Println(sa1.first)

	Elementstruct := struct {
		name     string
		branch   string
		language string
	}{
		name:     "Pikachu",
		branch:   "sddsfs",
		language: "English",
	}
	fmt.Println(Elementstruct)

}
