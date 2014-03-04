package main

import "fmt"

type Person struct {
	id     int
	name   string
	number string
	email  string
}

func main() {
	fmt.Println(Person{12, "Joe Dirt", "0800callme", "sexyjoe69@gmail.com"})
	fmt.Println(Person{id: 12, email: "sexyjoe69@gmail.com"})
}
