package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Helloworld, %s", request.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
