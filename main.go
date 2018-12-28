package main

import (
	"fmt"
	"log"
	"net/http"
)

// Run server: go run -race main.go
// Try requests: curl http://127.0.0.1:8000/test
func main() {
	http.HandleFunc("/", home)
	// http.ListenAndServe(":8000", nil)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello to DevHouse from %s...", r.URL.Path)
}
