package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = "80"

func welcome(w http.ResponseWriter, r *http.Request) {
	dataHTML := `<h1>Welcome to the world!</h1>`
	if r.Method == "GET" {
		fmt.Fprintf(w, dataHTML)
	}
}

func main() {
	fmt.Println("Berjalan di Port :", port)
	http.HandleFunc("/", welcome)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
