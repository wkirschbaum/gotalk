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

type Talker interface {
	say()
}

func chat(talker1 Talker, talker2 Talker) {
	talker1.say()
	talker2.say()
}

func main() {
	person := Person{12, "Joe Dirt", "0800callme", "sexyjoe69@gmail.com"}
	gopher := Gopher{}

	chat(person, gopher)
}
