package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hello my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {
	a := Android{
		Person: Person{Name: "Optimus Prime"},
		Model:  "2398Kjd"}

	// calling the method in two different ways for the same android
	a.Person.Talk()
	a.Talk()
}
