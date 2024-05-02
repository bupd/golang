package main

import "fmt"

type Person struct {
	Name string
}

type Printer interface {
	Print()
}

func (p Person) Print() {
	fmt.Println(p.Name)
}

func main() {
	p := Person{"Kumar"}
	PrintPerson(p)
}

func PrintPerson(p Printer) {
	p.Print()
}
