package main

import "fmt"

type person struct {
	firstname  string
	lastname   string
	favFlavors []string
}

func main() {
	p1 := person{
		firstname: "James",
		lastname:  "Bond",
		favFlavors: []string{
			"chocolate",
			"vennnila",
		},
	}

	p2 := person{
		firstname: "Miss",
		lastname:  "MoneyPenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
		},
	}
	fmt.Println(p1.firstname)
	fmt.Println(p1.lastname)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}
	fmt.Println(p2)

	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstname)
		fmt.Println(v.lastname)
	}

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"MoneyPenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

}
