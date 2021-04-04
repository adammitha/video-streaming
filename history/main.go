package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var videosCollection *mongo.Collection

var messageChannel *amqp.Channel

func main() {

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

	rabbitConnection, err := amqp.Dial(os.Getenv("RABBIT"))
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitConnection.Close()

	messageChannel, err = rabbitConnection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer messageChannel.Close()

	consumeViewedMessages()

}
