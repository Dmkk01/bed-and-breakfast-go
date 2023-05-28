package main

import (
	"fmt"
	"net/http"

	"github.com/Dmkk01/bed-and-breakfast/pkg/handlers"
)

const portNumber = ":11111"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s \n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
