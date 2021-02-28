package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/adammitha/video-streaming/utils"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{path}", videoHandler)
	r.Use(utils.LoggingMiddleware)

	log.Printf("Listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	rc, err := client.Bucket("video-storage-306005").Object(vars["path"]).NewReader(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rc.Close()
	io.Copy(w, rc)

}
