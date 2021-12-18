package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!\n")
	})

	log.Println("listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
