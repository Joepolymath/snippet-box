package main

import (
	"log"
	"net/http"
)

const PORT = ":4001"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcom to Snippet Box"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting Server on Port:", PORT)
	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		log.Fatal(err)
	}
}
