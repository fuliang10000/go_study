package main

import "fmt"

type Integer int

func main()  {
	var x Integer
	var y Integer
	x, y = 12, 10
	fmt.Println(x.Equal(y))
	fmt.Println(x.Dayu(y))
}

func (a Integer) Equal (b Integer) bool {
	return a == b
}

func (a Integer) Dayu (b Integer) bool {
	return  a > b
}
