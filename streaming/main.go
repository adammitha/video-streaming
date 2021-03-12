package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var videosCollection *mongo.Collection

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT env variable not set")
	}

	r := mux.NewRouter()

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
	videosCollection = client.Database("video-streaming").Collection("videos")

	r.HandleFunc("/video/{id}", getVideo)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func getVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Print(vars)

	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var videoRecord bson.M
	err = videosCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&videoRecord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	log.Print(videoRecord)

	res, err := http.Get(fmt.Sprintf("http://%s:%s/%s", os.Getenv("STORAGE_HOST"), os.Getenv("STORAGE_PORT"), videoRecord["videoPath"]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.Copy(w, res.Body)
	res.Body.Close()
}
