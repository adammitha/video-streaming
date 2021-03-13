package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type viewedRequest struct {
	VideoPath string `json:"videoPath"`
}

func viewedHandler(w http.ResponseWriter, r *http.Request) {
	var req viewedRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Unable to parse request", http.StatusBadRequest)
		return
	}

	_, err = videosCollection.InsertOne(context.Background(), bson.D{
		{Key: "videoPath", Value: req.VideoPath},
	})
	if err != nil {
		http.Error(w, "Error inserting record into database", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Added video %s to history", req.VideoPath)
}
