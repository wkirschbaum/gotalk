package main

import (
	"fmt"
	"time"
)

func say(s string, ch chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		ch <- s
	}
}

func main() {
	ch := make(chan string)

	go say("Hello, ", ch)
	go say("世界 \n", ch)

	//TODO: Try to explain it was a joke
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch)
	}
}
