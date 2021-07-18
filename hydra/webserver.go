package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleRoot)
	http.ListenAndServe(":7654", nil)
}

func HandleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to Hydra Project")
}
