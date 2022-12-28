package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/response", response)
	http.ListenAndServe(":8080", nil)
}

func response(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}
