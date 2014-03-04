package main

import (
	"fmt"
	"time"
)

func say(s string, quit chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(s)
	}
	quit <- "yes"
}

func main() {
	quit := make(chan string)

	go say("Hello, ", quit)
	go say("世界 \n", quit)

	for {
		select {
		case <-quit:
			fmt.Println("done")
			return
		}
	}
}
