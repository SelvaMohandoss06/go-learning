package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < 6; i++ {
		fmt.Println(i, x[i])
	}

	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	y := []int{1, 2, 3, 4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	marks := make([]int, 10, 12)
	fmt.Println(marks)
	marks[0] = 34

	jb := []string{"James", "Bond", "Chocolate"}
	fmt.Println(jb)

	mp := []string{"Miss", "MoneyPenny"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

}
