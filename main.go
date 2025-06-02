package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("++++ Starting Server ++++")
	http.FileServer(http.Dir("./staic"))
}