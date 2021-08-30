package main

import (
	"fmt"

	null "gopkg.in/guregu/null.v3"
)

type hotdog int

func main() {
	var x int
	fmt.Println(x)
	var s string
	fmt.Println(s)
	arr := [3]int{}
	fmt.Println(arr)
	slicevar := []int{}
	fmt.Println(slicevar)
	//intptr := 90
	// var x1 hotdog = 56
	// nullableInt := Int8fromPtr(*x1)

	fmt.Println(nullableInt)
}

type Int8 struct {
	null.Int
}

func NewInt8(i int8, valid bool) Int8 {
	return Int8{null.NewInt(int64(i), valid)}
}

func Int8from(i int8) Int8 {
	return NewInt8(i, true)
}

func Int8fromPtr(i *int8) Int8 {
	if i == nil {
		return NewInt8(0, false)
	}
	return NewInt8(*i, true)
}
