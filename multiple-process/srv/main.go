package main

import (
	"fmt"
	"net/http"
)

func srvHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	//w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/", srvHome)

	http.ListenAndServe(":80", nil)
	//http.ListenAndServeTLS(":8443", "example.crt", "example.key", nil)
}
