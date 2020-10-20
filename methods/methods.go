package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) PrintFullName() {
	fmt.Printf("%s %s\n", p.FirstName, p.LastName)
}

func main() {
	p := &Person{
		FirstName: "Ujjwal",
		LastName:  "Kumar",
	}
	p.PrintFullName()
}
