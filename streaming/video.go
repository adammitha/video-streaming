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
)

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

	sendViewedMessage(fmt.Sprintf("%s", videoRecord["videoPath"]))

	res, err := http.Get(fmt.Sprintf("http://%s:%s/%s", os.Getenv("STORAGE_HOST"), os.Getenv("STORAGE_PORT"), videoRecord["videoPath"]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.Copy(w, res.Body)
	res.Body.Close()
}
