package main

import (
	"log"
	"net/http"
)

const PORT = ":4001"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting Server on Port:", PORT)
	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		log.Fatal(err)
	}
}
