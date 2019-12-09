package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"./API"
)

func main() {
	http.ListenAndServe(":4040", API.Mux())
}

type data struct {
}

func jindex(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["Queenie"] = "QG"
	m["Roxy"] = "Cutes Cutes"
	m["Daisy"] = "Rosy Days"
	json.NewEncoder(w).Encode(m)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	ct := r.Header.Get("Content-Type")
	if strings.Contains(ct, "application/json") {
		fmt.Println("It's JSON!")
		data := make(map[string]string)
		json.NewDecoder(r.Body).Decode(&data)
		fmt.Println(data)
	} else if strings.Contains(ct, "text/html") {
		fmt.Println("It's HTML!")
	}
}
