package main

import "fmt"

type Gopher struct {
}

func (gopher Gopher) say() {
	fmt.Printf("My name is not important, I am a gopher \n")
}

type Person struct {
	id     int
	name   string
	number string
	email  string
}

func (person Person) say() {
	fmt.Printf("My name is %s \n", person.name)
}

func chat(person Person, gopher Gopher) {
	person.say()
	gopher.say()
}

func main() {
	person := Person{12, "Joe Dirt", "0800callme", "sexyjoe69@gmail.com"}
	gopher := Gopher{}

	chat(person, gopher)
}
