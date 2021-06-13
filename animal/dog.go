package animal

type Dog struct {
	Animal *Animal
	Pet Pet
}
func (d Dog) FavorFood() string {
	return "骨头"
}
func (d Dog) Call() string {
	return "汪汪汪"
}