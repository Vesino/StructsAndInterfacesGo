package main

import "fmt"

type parent struct {
	name string
}

func (p *parent) Info() string {
	return fmt.Sprintf("Name %s\n", p.name)
}

type child struct {
	parent
}

type Infoer interface {
	Info() string
}

func getInfo(i Infoer) string {
	return i.Info()
}

func main() {
	c := &child{
		parent: parent{
			name: "The Child",
		},
		// possible other fields
	}
	fmt.Println(c.Info())
	fmt.Println(c.name)

	// fmt.Println(getInfo(&c.parent)) 	// this is possible
	// fmt.Println(getInfo(c)) 			// this is not possible we need to create an interface

	fmt.Println(getInfo(c))

}
