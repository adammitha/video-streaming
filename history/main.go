package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adammitha/video-streaming/utils"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", historyHandler)
	r.Use(utils.LoggingMiddleware)

	port := os.Getenv("PORT")

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the History service")
}
