package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html, err := ioutil.ReadFile("staticFiles/index.html")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	fmt.Fprintf(w, "%s", html)
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("Failed to start Server: ", err)
	}
}
