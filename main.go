package main

import (
	"fmt"
	"net/http"
)

func helloHandlerFunc (w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Mrthod Not Supported", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello")
}

func formHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form error: %v", err)
		return 
	}
}

func main() {
	fmt.Println("++++ Starting Server ++++")
	fileHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileHandler)
	http.HandleFunc("/hello",helloHandlerFunc)
	http.HandleFunc("/form", formHandlerFunc)

	fmt.Println("Server running at PORT: 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}