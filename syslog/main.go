package main

/*
 * To let the binary run on alpine, build the main.go with:
 * CGO_ENABLED=0 go build main.go
 */

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello World\n")
	fmt.Fprintf(w, "1: %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
