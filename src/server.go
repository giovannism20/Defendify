package main

import (
	"Defendify/src/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", handlers.HelloWorld)

	http.ListenAndServe(":3000", nil)
}
