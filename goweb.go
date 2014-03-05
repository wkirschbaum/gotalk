package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./web/")))

	http.HandleFunc("/special", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You are special :)")
	})

	fmt.Println("Listening on localhost:4567")
	http.ListenAndServe(":4567", nil)
}
