package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) GetName() string {
	return a.Name
}

func (a Animal) Call() string {
	return "动物的叫声..."
}

type Dog struct {
	Animal
}

func main() {
	animal := Animal{"中华田园犬"}
	dog := Dog{animal}
	fmt.Println(dog.GetName())
	fmt.Println(dog.Call())
}