package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", ping)
	http.ListenAndServe(":3000", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
