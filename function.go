package main

import "fmt"

type Person struct {
	id     int
	name   string
	number string
	email  string
}

func (person Person) say() {
	fmt.Printf("My name is %s \n", person.name)
}

func say() {
	fmt.Println("I am not a method")
}

func main() {
	person := Person{12, "Joe Dirt", "0800callme", "sexyjoe69@gmail.com"}
	person.say()
	say()
}
