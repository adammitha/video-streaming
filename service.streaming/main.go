package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./videos"))))

	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
