package main

import "fmt"

type Student struct {
	id uint
	name string
	male bool
	score float64
}

func main()  {
	Student := NewStudent(1, "Fu Liang", 23.22)
	Student.SetName("JAY")
	fmt.Println(Student)
}

//func NewStudent(id uint, name string, male bool, score float64) *Student {
//	return &Student{id, name, male, score}
//}
func NewStudent(id uint, name string, score float64) *Student {
	return &Student{id: id, name:name, score:score}
}

func (s Student) GetName() string  {
	return s.name
}

func (s *Student) SetName(name string)  {
	s.name = name
}

func (s Student) String() string {
	return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}", s.id, s.name, s.male, s.score)
}