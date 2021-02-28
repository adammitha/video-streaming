package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

	r.HandleFunc("/{video}", getVideo)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func getVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	res, err := http.Get(fmt.Sprintf("%s:%s/%s.mp4", os.Getenv("STORAGE_HOST"), os.Getenv("STORAGE_PORT"), vars["video"]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.Copy(w, res.Body)
	res.Body.Close()
}
