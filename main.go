package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", ok)
	http.ListenAndServe(":9000", nil)
}

func ok(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
