package main

import "fmt"

type Animal struct{}

func (a *Animal) Talk() {
	fmt.Println("I am an animal!\n")
}

type Cat struct {
	Animal
	name string
}

type Dog struct {
	Animal
	name string
}

func (d *Dog) Talk() {
	fmt.Printf("Hello, my name is %s.\n\n", d.name)
}

func main() {
	var (
		animal = new(Animal)
		cat    = new(Cat)
		dog    = new(Dog)
	)

	dog.name = "Alfred"

	fmt.Println("An animal talking:")
	animal.Talk()

	fmt.Println("A cat talking:")
	cat.Talk()

	fmt.Println("A dog talking:")
	dog.Talk()

	fmt.Println("An animal talking:")
	dog.Animal.Talk()
}
