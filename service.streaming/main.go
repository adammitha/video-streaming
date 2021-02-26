package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	readEnv()
	port := viper.GetString("port")

	if port == "" {
		log.Fatal("PORT env variable not set")
	}

	r := mux.NewRouter()

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./videos"))))

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
