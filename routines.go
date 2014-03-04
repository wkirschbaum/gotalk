package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(s)
	}
}

func main() {
	go say("Hello, ")
	go say("世界 \n")

	//TODO: Fix hack before demo
	say("")
}
