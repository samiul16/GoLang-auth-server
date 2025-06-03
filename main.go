package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("++++ Starting Server ++++")
	fileHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileHandler)

	fmt.Println("Server running at PORT: 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}