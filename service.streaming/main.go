package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName(".env")
	r := mux.NewRouter()

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./videos"))))

	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
