package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":8000", nil)
}
