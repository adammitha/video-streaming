package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/adammitha/video-streaming/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var videosCollection *mongo.Collection

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/viewed", viewedHandler).Methods("POST")
	r.HandleFunc("/", historyHandler).Methods("GET")

	r.Use(utils.LoggingMiddleware)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_HOST")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	videosCollection = client.Database(os.Getenv("DB_NAME")).Collection("history")

	port := os.Getenv("PORT")

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the History service")
}
