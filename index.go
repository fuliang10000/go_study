package main

import (
	"fmt"
	"strconv"
)

func main() {
	//v1 := []rune{'\u4ed8', '\u4eae'}
	v1 := "100"
	v2, _ := strconv.Atoi(v1)
	fmt.Println(v2)
}