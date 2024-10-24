package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "fuck you!!!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/hello", helloWorldHandler)

	port := ":8088"
	fmt.Printf("Starting server on %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
