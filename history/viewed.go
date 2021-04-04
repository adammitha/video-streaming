package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
)

type viewedMessage struct {
	VideoPath string `json:"videoPath"`
}

func consumeViewedMessages() {
	messageChannel.ExchangeDeclare("viewed", "fanout", false, false, false, false, nil)

	q, err := messageChannel.QueueDeclare("", false, false, false, false, amqp.Table{"exclusive": true})
	if err != nil {
		log.Fatal(err)
	}

	err = messageChannel.QueueBind(q.Name, "", "viewed", false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := messageChannel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var decodedMessage viewedMessage
			json.NewDecoder(bytes.NewReader(d.Body)).Decode(&decodedMessage)
			_, err = videosCollection.InsertOne(context.Background(), bson.D{
				{Key: "videoPath", Value: decodedMessage.VideoPath},
			})
			if err != nil {
				log.Printf("Error inserting record into database: %v", err)
				continue
			}
			log.Printf("Successfully stored viewed message: %s", decodedMessage)
		}
	}()

	log.Printf("Waiting for messages...")
	<-forever
}
