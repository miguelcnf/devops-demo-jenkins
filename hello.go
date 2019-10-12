package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloWorld() string {
	return "Hello World"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, HelloWorld())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
