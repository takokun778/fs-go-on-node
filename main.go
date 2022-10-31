package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go!")
}

func main() {
	log.Printf("go server running...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
