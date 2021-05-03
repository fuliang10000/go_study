package main

import "fmt"

type Student struct {
	id uint
	name string
	male bool
	score float64
}

func main()  {
	Student := NewStudent(1, 'dfdf', true, 23.22)

	fmt.Println(Student)
}

func NewStudent(id uint, name string, male bool, score float64) *Student {
	return &Student(id, name, male, score)
}

